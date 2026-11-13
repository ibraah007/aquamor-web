package models

import "time"

// Product represents refills, bottled packs, and accessories
type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"` // 'refill', 'bottled_pack', 'accessory'
	Volume         string  `json:"volume"`   // '5L', '10L', '20L', '300ml', '500ml', '1L', '1.5L'
	PriceSwap      float64 `json:"price_swap"`
	PriceNewBottle float64 `json:"price_new_bottle"`
	IsAvailable    bool    `json:"is_available"`
}

// OrderItem details an individual item inside a customer's order
type OrderItem struct {
	ID           int     `json:"id"`
	OrderID      int     `json:"order_id"`
	ProductID    int     `json:"product_id"`
	BottleOption string  `json:"bottle_option"` // 'swap', 'new_purchase', 'not_applicable'
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	Subtotal     float64 `json:"subtotal"`
}

// OrderRequest handles incoming customer orders for delivery
type OrderRequest struct {
	CustomerName    string `json:"customer_name" binding:"required"`
	CustomerPhone   string `json:"customer_phone" binding:"required"`
	CustomerEmail   string `json:"customer_email,omitempty"`
	DeliveryEstate  string `json:"delivery_estate" binding:"required"` // e.g., Ngara, Parklands, CBD
	DeliveryAddress string `json:"delivery_address" binding:"required"`
	// unit_price and subtotal inside incoming items are ignored; the server
	// prices every line from the catalog.
	Items []OrderItem `json:"items" binding:"required"`
}

// Order represents a saved order in the system
type Order struct {
	ID                 int         `json:"id"`
	CustomerName       string      `json:"customer_name"`
	CustomerPhone      string      `json:"customer_phone"`
	CustomerEmail      string      `json:"customer_email,omitempty"`
	DeliveryEstate     string      `json:"delivery_estate"`
	DeliveryAddress    string      `json:"delivery_address"`
	TotalAmount        float64     `json:"total_amount"`
	OrderStatus        string      `json:"order_status"`   // 'pending', 'processing', 'out_for_delivery', 'completed'
	PaymentStatus      string      `json:"payment_status"` // 'unpaid', 'pending', 'paid', 'failed'
	PaymentProvider    string      `json:"payment_provider,omitempty"`
	PaymentReference   string      `json:"payment_reference,omitempty"`
	MpesaReceiptNumber string      `json:"mpesa_receipt_number,omitempty"`
	PaidAt             *time.Time  `json:"paid_at,omitempty"`
	CreatedAt          time.Time   `json:"created_at"`
	Items              []OrderItem `json:"items,omitempty"`
}

// EventQuoteRequest handles custom-branded bottle inquiries for special events
type EventQuoteRequest struct {
	ContactName string `json:"contact_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email"`
	EventType   string `json:"event_type" binding:"required"` // 'wedding', 'church', 'graduation', 'corporate'
	EventDate   string `json:"event_date" binding:"required"`
	BottleSize  string `json:"bottle_size" binding:"required"` // '300ml', '500ml', '1L'
	Quantity    int    `json:"quantity" binding:"required"`
	LogoFileURL string `json:"logo_file_url"`
	Notes       string `json:"notes"`
}
