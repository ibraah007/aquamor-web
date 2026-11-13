package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"aquamor-backend/backend/internal/catalog"
	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/models"
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "aquamor-handlers-")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create test directory: %v\n", err)
		os.Exit(1)
	}

	if err := db.InitDB(filepath.Join(dir, "test.db")); err != nil {
		fmt.Fprintf(os.Stderr, "initialise test database: %v\n", err)
		os.RemoveAll(dir)
		os.Exit(1)
	}

	code := m.Run()

	db.Close()
	os.RemoveAll(dir)
	os.Exit(code)
}

func validOrder() models.OrderRequest {
	return models.OrderRequest{
		CustomerName:    "Amina Wanjiku",
		CustomerPhone:   "0722000111",
		DeliveryEstate:  "Parklands",
		DeliveryAddress: "Oloitoktok Road",
		Items: []models.OrderItem{
			{ProductID: 1, BottleOption: "swap", Quantity: 2},
		},
	}
}

func postOrder(t *testing.T, req models.OrderRequest) *httptest.ResponseRecorder {
	t.Helper()
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}
	return postOrderBody(body)
}

func postOrderBody(body []byte) *httptest.ResponseRecorder {
	request := httptest.NewRequest(http.MethodPost, "/api/v1/orders", strings.NewReader(string(body)))
	recorder := httptest.NewRecorder()
	CreateOrder(recorder, request)
	return recorder
}

func decodeCreatedOrder(t *testing.T, recorder *httptest.ResponseRecorder) models.Order {
	t.Helper()
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d; body: %s", recorder.Code, http.StatusCreated, recorder.Body.String())
	}
	var response struct {
		Order models.Order `json:"order"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v; body: %s", err, recorder.Body.String())
	}
	return response.Order
}

func TestCreateOrderUsesServerPrices(t *testing.T) {
	req := validOrder()
	// The client sends fake prices; the handler must ignore them.
	req.Items = []models.OrderItem{
		{ProductID: 1, BottleOption: "swap", Quantity: 2, UnitPrice: 1, Subtotal: 1},
		{ProductID: 4, BottleOption: "new_purchase", Quantity: 1, UnitPrice: 1, Subtotal: 1},
	}

	order := decodeCreatedOrder(t, postOrder(t, req))

	if order.TotalAmount != 1800 {
		t.Errorf("total = %.2f, want 1800", order.TotalAmount)
	}
	if len(order.Items) != 2 {
		t.Fatalf("items = %d, want 2", len(order.Items))
	}
	want := []struct {
		product  int
		option   string
		quantity int
		unit     float64
		subtotal float64
	}{
		{1, "swap", 2, 200, 400},
		{4, "new_purchase", 1, 1400, 1400},
	}
	for i, w := range want {
		item := order.Items[i]
		if item.ProductID != w.product || item.BottleOption != w.option || item.Quantity != w.quantity ||
			item.UnitPrice != w.unit || item.Subtotal != w.subtotal {
			t.Errorf("item %d = %+v, want product=%d option=%s quantity=%d unit=%.2f subtotal=%.2f",
				i, item, w.product, w.option, w.quantity, w.unit, w.subtotal)
		}
	}

	var storedTotal float64
	if err := db.DB.QueryRow("SELECT total_amount FROM orders WHERE id = ?", order.ID).Scan(&storedTotal); err != nil {
		t.Fatalf("read stored total: %v", err)
	}
	if storedTotal != 1800 {
		t.Errorf("stored total = %.2f, want 1800", storedTotal)
	}
}

func TestCreateOrderAllowsMultipleOrdersWithoutPaymentReference(t *testing.T) {
	first := decodeCreatedOrder(t, postOrder(t, validOrder()))
	second := decodeCreatedOrder(t, postOrder(t, validOrder()))

	if first.ID == second.ID {
		t.Fatalf("both orders got ID %d", first.ID)
	}
	var withNull int
	if err := db.DB.QueryRow(
		`SELECT COUNT(*) FROM orders WHERE payment_reference IS NULL AND id IN (?, ?)`,
		first.ID, second.ID,
	).Scan(&withNull); err != nil {
		t.Fatalf("count orders: %v", err)
	}
	if withNull != 2 {
		t.Errorf("orders with NULL payment_reference = %d, want 2", withNull)
	}
}

func TestCreateOrderRejectsUnknownOrUnavailableProduct(t *testing.T) {
	for _, id := range []int{6, 999} {
		t.Run(fmt.Sprintf("product_%d", id), func(t *testing.T) {
			req := validOrder()
			req.Items[0].ProductID = id

			var before int
			if err := db.DB.QueryRow("SELECT COUNT(*) FROM orders").Scan(&before); err != nil {
				t.Fatalf("count orders: %v", err)
			}

			recorder := postOrder(t, req)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400; body: %s", recorder.Code, recorder.Body.String())
			}
			if !strings.Contains(recorder.Body.String(), "unknown or unavailable") {
				t.Errorf("body = %s, want an unavailable-product message", recorder.Body.String())
			}

			var after int
			if err := db.DB.QueryRow("SELECT COUNT(*) FROM orders").Scan(&after); err != nil {
				t.Fatalf("count orders: %v", err)
			}
			if after != before {
				t.Errorf("orders changed from %d to %d", before, after)
			}
		})
	}
}

func TestCreateOrderRejectsMissingBottlePrice(t *testing.T) {
	tests := []struct {
		name    string
		req     models.OrderRequest
		errPart string
	}{
		{
			name: "refill has no new-bottle price",
			req: func() models.OrderRequest {
				req := validOrder()
				req.Items[0].BottleOption = "new_purchase"
				return req
			}(),
			errPart: "No price configured",
		},
		{
			name: "hard bottle has no swap price",
			req: func() models.OrderRequest {
				req := validOrder()
				req.Items[0] = models.OrderItem{ProductID: 4, BottleOption: "swap", Quantity: 1}
				return req
			}(),
			errPart: "No price configured",
		},
		{
			name: "unknown bottle option",
			req: func() models.OrderRequest {
				req := validOrder()
				req.Items[0].BottleOption = "rental"
				return req
			}(),
			errPart: "No price configured",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := postOrder(t, test.req)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400; body: %s", recorder.Code, recorder.Body.String())
			}
			if !strings.Contains(recorder.Body.String(), test.errPart) {
				t.Errorf("body = %s, want %q", recorder.Body.String(), test.errPart)
			}
		})
	}
}

func TestCreateOrderValidatesFields(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*models.OrderRequest)
		errPart string
	}{
		{"short name", func(r *models.OrderRequest) { r.CustomerName = "A" }, "Customer name"},
		{"long name", func(r *models.OrderRequest) { r.CustomerName = strings.Repeat("a", 81) }, "Customer name"},
		{"short phone", func(r *models.OrderRequest) { r.CustomerPhone = "123456" }, "Customer phone"},
		{"long phone", func(r *models.OrderRequest) { r.CustomerPhone = strings.Repeat("1", 21) }, "Customer phone"},
		{"short estate", func(r *models.OrderRequest) { r.DeliveryEstate = "N" }, "Delivery estate"},
		{"long estate", func(r *models.OrderRequest) { r.DeliveryEstate = strings.Repeat("a", 81) }, "Delivery estate"},
		{"short address", func(r *models.OrderRequest) { r.DeliveryAddress = "X" }, "Delivery address"},
		{"long address", func(r *models.OrderRequest) { r.DeliveryAddress = strings.Repeat("a", 161) }, "Delivery address"},
		{"no items", func(r *models.OrderRequest) { r.Items = nil }, "between 1 and 50 items"},
		{"too many items", func(r *models.OrderRequest) {
			r.Items = make([]models.OrderItem, 51)
			for i := range r.Items {
				r.Items[i] = models.OrderItem{ProductID: 1, BottleOption: "swap", Quantity: 1}
			}
		}, "between 1 and 50 items"},
		{"zero quantity", func(r *models.OrderRequest) { r.Items[0].Quantity = 0 }, "quantity"},
		{"too many units", func(r *models.OrderRequest) { r.Items[0].Quantity = 1001 }, "quantity"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := validOrder()
			test.mutate(&req)

			recorder := postOrder(t, req)
			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400; body: %s", recorder.Code, recorder.Body.String())
			}
			if !strings.Contains(recorder.Body.String(), test.errPart) {
				t.Errorf("body = %s, want %q", recorder.Body.String(), test.errPart)
			}
		})
	}
}

func TestCreateOrderRejectsTrailingJSON(t *testing.T) {
	req := validOrder()
	body, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal order: %v", err)
	}
	body = append(body, []byte("\n{}")...)

	recorder := postOrderBody(body)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body: %s", recorder.Code, recorder.Body.String())
	}
	if !strings.Contains(recorder.Body.String(), "Invalid request payload") {
		t.Errorf("body = %s, want an invalid-payload message", recorder.Body.String())
	}
}

func TestCreateOrderRejectsOversizedBody(t *testing.T) {
	body := []byte(`{"customer_name":"` + strings.Repeat("a", maxOrderBody) + `"}`)

	recorder := postOrderBody(body)
	if recorder.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d, want 413; body: %s", recorder.Code, recorder.Body.String())
	}
}

func TestCreateOrderRejectsWrongMethod(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
	recorder := httptest.NewRecorder()
	CreateOrder(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", recorder.Code)
	}
}

func TestGetOrdersRequiresToken(t *testing.T) {
	t.Run("token not configured", func(t *testing.T) {
		t.Setenv("ADMIN_TOKEN", "")
		recorder := httptest.NewRecorder()
		GetOrders(recorder, httptest.NewRequest(http.MethodGet, "/api/v1/orders/all", nil))

		if recorder.Code != http.StatusServiceUnavailable {
			t.Fatalf("status = %d, want 503; body: %s", recorder.Code, recorder.Body.String())
		}
	})

	t.Run("missing and wrong token", func(t *testing.T) {
		t.Setenv("ADMIN_TOKEN", "correct-token")

		for _, header := range []string{"", "Bearer wrong-token", "Basic correct-token"} {
			request := httptest.NewRequest(http.MethodGet, "/api/v1/orders/all", nil)
			if header != "" {
				request.Header.Set("Authorization", header)
			}
			recorder := httptest.NewRecorder()
			GetOrders(recorder, request)

			if recorder.Code != http.StatusUnauthorized {
				t.Errorf("Authorization %q: status = %d, want 401", header, recorder.Code)
			}
		}
	})

	t.Run("correct token", func(t *testing.T) {
		t.Setenv("ADMIN_TOKEN", "correct-token")
		request := httptest.NewRequest(http.MethodGet, "/api/v1/orders/all", nil)
		request.Header.Set("Authorization", "Bearer correct-token")
		recorder := httptest.NewRecorder()
		GetOrders(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, want 200; body: %s", recorder.Code, recorder.Body.String())
		}
		if got := recorder.Header().Get("Cache-Control"); got != "no-store" {
			t.Errorf("Cache-Control = %q, want no-store", got)
		}
		var orders []models.Order
		if err := json.Unmarshal(recorder.Body.Bytes(), &orders); err != nil {
			t.Fatalf("decode orders: %v; body: %s", err, recorder.Body.String())
		}
	})
}

func TestUnitPrice(t *testing.T) {
	refill, ok := catalog.ByID(1)
	if !ok {
		t.Fatal("refill product 1 is missing")
	}
	hardBottle, ok := catalog.ByID(4)
	if !ok {
		t.Fatal("hard bottle product 4 is missing")
	}

	tests := []struct {
		name    string
		product models.Product
		option  string
		want    float64
		ok      bool
	}{
		{"refill swap", refill, "swap", 200, true},
		{"refill new bottle", refill, "new_purchase", 0, false},
		{"refill not applicable", refill, "not_applicable", 200, true},
		{"hard bottle new", hardBottle, "new_purchase", 1400, true},
		{"hard bottle swap", hardBottle, "swap", 0, false},
		{"unknown option", refill, "rental", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := unitPrice(test.product, test.option)
			if got != test.want || ok != test.ok {
				t.Errorf("unitPrice(%s) = (%.2f, %v), want (%.2f, %v)", test.option, got, ok, test.want, test.ok)
			}
		})
	}
}
