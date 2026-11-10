package main

import (
	"fmt"
	"net/http"

	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/handlers"
)

func main() {
	// Initialize SQLite Database and migration tables
	db.InitDB()

	// Serve Frontend Static HTML at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "backend/internal/web/index.html")
	})

	// Health check endpoint
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok", "message":"Aquamor Water API is running"}`)
	})

	// Products catalog endpoint
	http.HandleFunc("/api/v1/products", handlers.GetProducts)

	// Order placement endpoint
	http.HandleFunc("/api/v1/orders", handlers.CreateOrder)

	// Order history endpoint
	http.HandleFunc("/api/v1/orders/all", handlers.GetOrders)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
