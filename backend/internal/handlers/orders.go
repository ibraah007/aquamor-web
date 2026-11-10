package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"aquamor-backend/backend/internal/models"
)

// CreateOrder handles incoming customer delivery orders
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Ensure method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	var req models.OrderRequest

	// Decode incoming JSON body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Basic Validation
	if req.CustomerName == "" || req.CustomerPhone == "" || req.DeliveryEstate == "" || len(req.Items) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing required order fields or items"})
		return
	}

	// Calculate total amount and assign subtotals (mock pricing logic for now)
	var totalAmount float64
	var processedItems []models.OrderItem

	for i, item := range req.Items {
		// Calculate item subtotal (UnitPrice * Quantity)
		subtotal := item.UnitPrice * float64(item.Quantity)
		totalAmount += subtotal

		processedItems = append(processedItems, models.OrderItem{
			ID:           i + 1,
			ProductID:    item.ProductID,
			BottleOption: item.BottleOption,
			Quantity:     item.Quantity,
			UnitPrice:    item.UnitPrice,
			Subtotal:     subtotal,
		})
	}

	// Construct saved order response object
	newOrder := models.Order{
		ID:              1001, // Mock generated ID (will be replaced by DB serial ID later)
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

	// Return success response with the created order details
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order placed successfully! Awaiting M-Pesa payment verification.",
		"order":   newOrder,
	})
}