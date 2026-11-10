package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/models"
)

// CreateOrder handles incoming customer delivery orders and saves them to SQLite
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	var req models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if req.CustomerName == "" || req.CustomerPhone == "" || req.DeliveryEstate == "" || len(req.Items) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing required order fields or items"})
		return
	}

	// Calculate total amount
	var totalAmount float64
	for _, item := range req.Items {
		totalAmount += item.UnitPrice * float64(item.Quantity)
	}

	// Insert Order into SQLite database
	result, err := db.DB.Exec(
		`INSERT INTO orders (customer_name, customer_phone, delivery_estate, delivery_address, total_amount, order_status, payment_status) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		req.CustomerName, req.CustomerPhone, req.DeliveryEstate, req.DeliveryAddress, totalAmount, "pending", "unpaid",
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order to database"})
		return
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to retrieve generated order ID"})
		return
	}

	// Insert individual items into SQLite order_items table
	var processedItems []models.OrderItem
	for i, item := range req.Items {
		subtotal := item.UnitPrice * float64(item.Quantity)
		_, err := db.DB.Exec(
			`INSERT INTO order_items (order_id, product_id, bottle_option, quantity, unit_price, subtotal) VALUES (?, ?, ?, ?, ?, ?)`,
			orderID, item.ProductID, item.BottleOption, item.Quantity, item.UnitPrice, subtotal,
		)
		if err != nil {
			// Log or handle error if needed
		}

		processedItems = append(processedItems, models.OrderItem{
			ID:           i + 1,
			OrderID:      int(orderID),
			ProductID:    item.ProductID,
			BottleOption: item.BottleOption,
			Quantity:     item.Quantity,
			UnitPrice:    item.UnitPrice,
			Subtotal:     subtotal,
		})
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
		CreatedAt:       time.Now(),
		Items:           processedItems,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order placed and saved to SQLite database successfully!",
		"order":   newOrder,
	})
}

// GetOrders fetches all customer orders from SQLite
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	rows, err := db.DB.Query(`SELECT id, customer_name, customer_phone, delivery_estate, delivery_address, total_amount, order_status, payment_status, created_at FROM orders ORDER BY created_at DESC`)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch orders"})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.CustomerPhone, &o.DeliveryEstate, &o.DeliveryAddress, &o.TotalAmount, &o.OrderStatus, &o.PaymentStatus, &o.CreatedAt); err != nil {
			continue
		}
		orders = append(orders, o)
	}

	if orders == nil {
		orders = []models.Order{}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
