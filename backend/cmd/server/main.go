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

	// API routes. Go's ServeMux picks the longest matching pattern, so these
	// win over the static file handler at "/".
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok", "message":"Aquamor Water API is running"}`)
	})
	http.HandleFunc("/api/v1/products", handlers.GetProducts)
	http.HandleFunc("/api/v1/orders", handlers.CreateOrder)
	http.HandleFunc("/api/v1/orders/all", handlers.GetOrders)

	// Static site: web/index.html plus its assets. Run the server from the
	// repository root so this relative path resolves.
	http.Handle("/", http.FileServer(http.Dir("web")))

	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
