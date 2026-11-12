package handlers

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"aquamor-backend/backend/internal/catalog"
	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/models"
)

const maxOrderBody = 1 << 20 // 1 MB

// CreateOrder handles incoming customer delivery orders and saves them to SQLite.
// Prices are taken from the server catalog; any unit_price sent by the client is ignored.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxOrderBody)
	decoder := json.NewDecoder(r.Body)
	var req models.OrderRequest
	if err := decoder.Decode(&req); err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			json.NewEncoder(w).Encode(map[string]string{"error": "Request body is too large"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	// Reject a second JSON value or any trailing content after the order.
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			json.NewEncoder(w).Encode(map[string]string{"error": "Request body is too large"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	req.CustomerName = strings.TrimSpace(req.CustomerName)
	req.CustomerPhone = strings.TrimSpace(req.CustomerPhone)
	req.DeliveryEstate = strings.TrimSpace(req.DeliveryEstate)
	req.DeliveryAddress = strings.TrimSpace(req.DeliveryAddress)

	if err := validateOrder(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	type orderLine struct {
		product  models.Product
		option   string
		quantity int
		unit     float64
		subtotal float64
	}

	lines := make([]orderLine, 0, len(req.Items))
	var totalAmount float64
	for _, item := range req.Items {
		product, ok := catalog.ByID(item.ProductID)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Product %d is unknown or unavailable", item.ProductID)})
			return
		}

		option := strings.ToLower(strings.TrimSpace(item.BottleOption))
		if option == "" {
			option = "not_applicable"
		}
		unit, ok := unitPrice(product, option)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("No price configured for product %d (%s)", product.ID, option)})
			return
		}

		subtotal := unit * float64(item.Quantity)
		totalAmount += subtotal
		lines = append(lines, orderLine{product: product, option: option, quantity: item.Quantity, unit: unit, subtotal: subtotal})
	}

	tx, err := db.DB.Begin()
	if err != nil {
		log.Printf("create order: begin transaction: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order"})
		return
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		`INSERT INTO orders (customer_name, customer_phone, delivery_estate, delivery_address, total_amount, order_status, payment_status) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		req.CustomerName, req.CustomerPhone, req.DeliveryEstate, req.DeliveryAddress, totalAmount, "pending", "unpaid",
	)
	if err != nil {
		log.Printf("create order: insert order: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order"})
		return
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		log.Printf("create order: last insert id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order"})
		return
	}

	processedItems := make([]models.OrderItem, 0, len(lines))
	for _, line := range lines {
		itemResult, err := tx.Exec(
			`INSERT INTO order_items (order_id, product_id, bottle_option, quantity, unit_price, subtotal) VALUES (?, ?, ?, ?, ?, ?)`,
			orderID, line.product.ID, line.option, line.quantity, line.unit, line.subtotal,
		)
		if err != nil {
			log.Printf("create order: insert item: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order items"})
			return
		}
		itemID, err := itemResult.LastInsertId()
		if err != nil {
			log.Printf("create order: item last insert id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order items"})
			return
		}
		processedItems = append(processedItems, models.OrderItem{
			ID:           int(itemID),
			OrderID:      int(orderID),
			ProductID:    line.product.ID,
			BottleOption: line.option,
			Quantity:     line.quantity,
			UnitPrice:    line.unit,
			Subtotal:     line.subtotal,
		})
	}

	if err := tx.Commit(); err != nil {
		log.Printf("create order: commit: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order"})
		return
	}

	newOrder := models.Order{
		ID:              int(orderID),
		CustomerName:    req.CustomerName,
		CustomerPhone:   req.CustomerPhone,
		DeliveryEstate:  req.DeliveryEstate,
		DeliveryAddress: req.DeliveryAddress,
		TotalAmount:     totalAmount,
		OrderStatus:     "pending",
		PaymentStatus:   "unpaid",
		CreatedAt:       time.Now().UTC(),
		Items:           processedItems,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order placed and saved successfully",
		"order":   newOrder,
	})
}

// GetOrders returns customer orders. Customer contact details are PII, so the
// endpoint requires ADMIN_TOKEN and an "Authorization: Bearer <token>" header.
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	token := os.Getenv("ADMIN_TOKEN")
	if token == "" {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "ADMIN_TOKEN is not configured; orders are not exposed"})
		return
	}
	provided, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !ok || subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	rows, err := db.DB.Query(`SELECT id, customer_name, customer_phone, delivery_estate, delivery_address, total_amount, order_status, payment_status, created_at FROM orders ORDER BY created_at DESC LIMIT 500`)
	if err != nil {
		log.Printf("get orders: query: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch orders"})
		return
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.CustomerPhone, &o.DeliveryEstate, &o.DeliveryAddress, &o.TotalAmount, &o.OrderStatus, &o.PaymentStatus, &o.CreatedAt); err != nil {
			log.Printf("get orders: scan: %v", err)
			continue
		}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		log.Printf("get orders: rows: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func validateOrder(req models.OrderRequest) error {
	if l := len(req.CustomerName); l < 2 || l > 80 {
		return errors.New("Customer name must be between 2 and 80 characters")
	}
	if l := len(req.CustomerPhone); l < 7 || l > 20 {
		return errors.New("Customer phone must be between 7 and 20 characters")
	}
	if l := len(req.DeliveryEstate); l < 2 || l > 80 {
		return errors.New("Delivery estate must be between 2 and 80 characters")
	}
	if l := len(req.DeliveryAddress); l < 2 || l > 160 {
		return errors.New("Delivery address must be between 2 and 160 characters")
	}
	if len(req.Items) == 0 || len(req.Items) > 50 {
		return errors.New("An order must contain between 1 and 50 items")
	}
	for _, item := range req.Items {
		if item.Quantity < 1 || item.Quantity > 1000 {
			return errors.New("Item quantity must be between 1 and 1000")
		}
	}
	return nil
}

// unitPrice chooses the catalog price for the requested bottle option.
func unitPrice(product models.Product, option string) (float64, bool) {
	switch option {
	case "new_purchase":
		if product.PriceNewBottle > 0 {
			return product.PriceNewBottle, true
		}
		return 0, false
	case "swap":
		if product.PriceSwap > 0 {
			return product.PriceSwap, true
		}
		return 0, false
	case "not_applicable":
		if product.PriceSwap > 0 {
			return product.PriceSwap, true
		}
		if product.PriceNewBottle > 0 {
			return product.PriceNewBottle, true
		}
		return 0, false
	default:
		return 0, false
	}
}
