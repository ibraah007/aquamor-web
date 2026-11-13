package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/models"
	"aquamor-backend/backend/internal/payments"
)

func paymentTestEnv(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	server := httptest.NewServer(handler)
	oldBase, oldClient := payments.BaseURL, payments.HTTPClient
	payments.BaseURL = server.URL
	payments.HTTPClient = server.Client()
	t.Cleanup(func() {
		payments.BaseURL = oldBase
		payments.HTTPClient = oldClient
		server.Close()
	})
	t.Setenv("FLUTTERWAVE_SECRET_KEY", "sk_test_123")
	t.Setenv("FLUTTERWAVE_WEBHOOK_SECRET", "hook-secret")
	t.Setenv("PUBLIC_BASE_URL", "https://aquamor.test")
}

func TestPaymentConfigReadsEnvironment(t *testing.T) {
	t.Setenv("FLUTTERWAVE_SECRET_KEY", "")
	t.Setenv("FLUTTERWAVE_WEBHOOK_SECRET", "")
	t.Setenv("PUBLIC_BASE_URL", "")

	recorder := httptest.NewRecorder()
	PaymentConfig(recorder, httptest.NewRequest(http.MethodGet, "/api/v1/payments/config", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", recorder.Code)
	}
	var response struct {
		Enabled bool `json:"enabled"`
	}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	if response.Enabled {
		t.Error("enabled = true with no configuration")
	}

	t.Setenv("FLUTTERWAVE_SECRET_KEY", "sk_test_123")
	t.Setenv("FLUTTERWAVE_WEBHOOK_SECRET", "hook-secret")
	t.Setenv("PUBLIC_BASE_URL", "https://aquamor.test")
	recorder = httptest.NewRecorder()
	PaymentConfig(recorder, httptest.NewRequest(http.MethodGet, "/api/v1/payments/config", nil))
	json.Unmarshal(recorder.Body.Bytes(), &response)
	if !response.Enabled {
		t.Errorf("enabled = false with full configuration; body: %s", recorder.Body.String())
	}
}

func TestCreateCheckoutRequiresConfiguration(t *testing.T) {
	t.Setenv("FLUTTERWAVE_SECRET_KEY", "")
	t.Setenv("FLUTTERWAVE_WEBHOOK_SECRET", "")
	t.Setenv("PUBLIC_BASE_URL", "")

	req := validOrder()
	req.CustomerEmail = "amina@example.com"
	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	CreateCheckout(recorder, httptest.NewRequest(http.MethodPost, "/api/v1/payments/checkout", strings.NewReader(string(body))))

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503; body: %s", recorder.Code, recorder.Body.String())
	}
	if !strings.Contains(recorder.Body.String(), "payments_not_configured") {
		t.Errorf("body = %s, want payments_not_configured", recorder.Body.String())
	}
}

func TestCreateCheckoutSavesOrderAndReturnsLink(t *testing.T) {
	var mu sync.Mutex
	var checkoutPayload map[string]interface{}

	paymentTestEnv(t, func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decode checkout payload: %v", err)
		}
		mu.Lock()
		checkoutPayload = payload
		mu.Unlock()
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"data":   map[string]string{"link": "https://checkout.flutterwave.test/pay/abc"},
		})
	})

	req := validOrder()
	req.CustomerEmail = "amina@example.com"
	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	CreateCheckout(recorder, httptest.NewRequest(http.MethodPost, "/api/v1/payments/checkout", strings.NewReader(string(body))))

	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201; body: %s", recorder.Code, recorder.Body.String())
	}
	var response struct {
		Order      models.Order `json:"order"`
		PaymentURL string       `json:"payment_url"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if response.PaymentURL != "https://checkout.flutterwave.test/pay/abc" {
		t.Errorf("payment_url = %q", response.PaymentURL)
	}
	if response.Order.PaymentStatus != "pending" || response.Order.PaymentProvider != paymentProvider || response.Order.PaymentReference == "" {
		t.Errorf("order payment fields = %#v", response.Order)
	}
	if response.Order.TotalAmount != 400 {
		t.Errorf("total = %v, want 400", response.Order.TotalAmount)
	}

	mu.Lock()
	amount, _ := checkoutPayload["amount"].(float64)
	txRef, _ := checkoutPayload["tx_ref"].(string)
	mu.Unlock()
	if amount != 400 || txRef != response.Order.PaymentReference {
		t.Errorf("checkout payload amount=%v tx_ref=%v, want 400 and %s", amount, txRef, response.Order.PaymentReference)
	}
}

func TestVerifyPaymentMarksOrderPaid(t *testing.T) {
	var mu sync.Mutex
	var reference string

	paymentTestEnv(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v3/payments":
			var payload struct {
				TxRef string `json:"tx_ref"`
			}
			json.NewDecoder(r.Body).Decode(&payload)
			mu.Lock()
			reference = payload.TxRef
			mu.Unlock()
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "success",
				"data":   map[string]string{"link": "https://checkout.flutterwave.test/pay/abc"},
			})
		case strings.HasPrefix(r.URL.Path, "/v3/transactions/"):
			mu.Lock()
			ref := reference
			mu.Unlock()
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "success",
				"data": map[string]interface{}{
					"id":       555,
					"tx_ref":   ref,
					"amount":   400,
					"currency": "KES",
					"status":   "successful",
				},
			})
		default:
			http.NotFound(w, r)
		}
	})

	req := validOrder()
	req.CustomerEmail = "amina@example.com"
	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	CreateCheckout(recorder, httptest.NewRequest(http.MethodPost, "/api/v1/payments/checkout", strings.NewReader(string(body))))
	if recorder.Code != http.StatusCreated {
		t.Fatalf("checkout status = %d; body: %s", recorder.Code, recorder.Body.String())
	}
	var created struct {
		Order models.Order `json:"order"`
	}
	json.Unmarshal(recorder.Body.Bytes(), &created)

	verifyReq := httptest.NewRequest(http.MethodGet, "/api/v1/payments/verify?reference="+created.Order.PaymentReference+"&transaction_id=555", nil)
	verifyRecorder := httptest.NewRecorder()
	VerifyPayment(verifyRecorder, verifyReq)
	if verifyRecorder.Code != http.StatusOK {
		t.Fatalf("verify status = %d; body: %s", verifyRecorder.Code, verifyRecorder.Body.String())
	}
	if !strings.Contains(verifyRecorder.Body.String(), `"status":"paid"`) {
		t.Errorf("verify body = %s, want paid", verifyRecorder.Body.String())
	}

	var paymentStatus, orderStatus, receipt string
	if err := db.DB.QueryRow(`SELECT payment_status, order_status, mpesa_receipt_number FROM orders WHERE id = ?`, created.Order.ID).Scan(&paymentStatus, &orderStatus, &receipt); err != nil {
		t.Fatalf("read order: %v", err)
	}
	if paymentStatus != "paid" || orderStatus != "processing" || receipt != "555" {
		t.Errorf("order state = %s/%s/%s, want paid/processing/555", paymentStatus, orderStatus, receipt)
	}

	t.Setenv("ADMIN_TOKEN", "test-token")
	adminReq := httptest.NewRequest(http.MethodGet, "/api/v1/orders/all", nil)
	adminReq.Header.Set("Authorization", "Bearer test-token")
	adminRecorder := httptest.NewRecorder()
	GetOrders(adminRecorder, adminReq)
	if adminRecorder.Code != http.StatusOK {
		t.Fatalf("admin orders status = %d; body: %s", adminRecorder.Code, adminRecorder.Body.String())
	}
	var orders []models.Order
	if err := json.Unmarshal(adminRecorder.Body.Bytes(), &orders); err != nil {
		t.Fatalf("decode orders: %v", err)
	}
	var found *models.Order
	for i := range orders {
		if orders[i].ID == created.Order.ID {
			found = &orders[i]
			break
		}
	}
	if found == nil {
		t.Fatalf("paid order %d not returned by GetOrders", created.Order.ID)
	}
	if found.PaidAt == nil || found.PaymentReference != created.Order.PaymentReference || found.MpesaReceiptNumber != "555" {
		t.Errorf("paid order fields = %#v", found)
	}
}

func TestPaymentWebhookMarksOrderPaid(t *testing.T) {
	paymentTestEnv(t, func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	req := validOrder()
	lines, total, err := priceOrder(req)
	if err != nil {
		t.Fatalf("price order: %v", err)
	}
	reference := fmt.Sprintf("aquamor-webhook-test-%d", time.Now().UnixNano())
	order, err := insertOrder(req, lines, total, paymentMeta{
		PaymentProvider:  paymentProvider,
		PaymentReference: reference,
		PaymentStatus:    "pending",
	})
	if err != nil {
		t.Fatalf("insert order: %v", err)
	}

	payload := `{"event":"charge.completed","data":{"id":777,"tx_ref":"` + reference + `","status":"successful","amount":400,"currency":"KES"}}`

	bad := httptest.NewRecorder()
	badReq := httptest.NewRequest(http.MethodPost, "/api/v1/payments/webhook", strings.NewReader(payload))
	badReq.Header.Set("verif-hash", "wrong")
	PaymentWebhook(bad, badReq)
	if bad.Code != http.StatusUnauthorized {
		t.Errorf("wrong signature status = %d, want 401", bad.Code)
	}

	recorder := httptest.NewRecorder()
	webhookReq := httptest.NewRequest(http.MethodPost, "/api/v1/payments/webhook", strings.NewReader(payload))
	webhookReq.Header.Set("verif-hash", "hook-secret")
	PaymentWebhook(recorder, webhookReq)
	if recorder.Code != http.StatusOK {
		t.Fatalf("webhook status = %d; body: %s", recorder.Code, recorder.Body.String())
	}

	var paymentStatus, receipt string
	if err := db.DB.QueryRow(`SELECT payment_status, mpesa_receipt_number FROM orders WHERE id = ?`, order.ID).Scan(&paymentStatus, &receipt); err != nil {
		t.Fatalf("read order: %v", err)
	}
	if paymentStatus != "paid" || receipt != "777" {
		t.Errorf("order state = %s/%s, want paid/777", paymentStatus, receipt)
	}
}
