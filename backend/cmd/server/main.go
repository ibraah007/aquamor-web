package main

import (
	"fmt"
	"net/http"

	"aquamor-backend/backend/internal/handlers"
)

func main() {
	// Health check endpoint
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok", "message":"Aquamor Water API is running"}`)
	})

	// Products endpoint
	http.HandleFunc("/api/v1/products", handlers.GetProducts)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}