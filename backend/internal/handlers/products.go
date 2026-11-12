package handlers

import (
	"encoding/json"
	"net/http"

	"aquamor-backend/backend/internal/catalog"
	"aquamor-backend/backend/internal/models"
)

// GetProducts returns the available catalog of Aquamor products.
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	products := catalog.All()
	available := make([]models.Product, 0, len(products))
	for _, p := range products {
		if p.IsAvailable {
			available = append(available, p)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(available)
}
