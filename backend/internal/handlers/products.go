package handlers

import (
	"encoding/json"
	"net/http"

	"aquamor-backend/backend/internal/models"
)

// GetProducts returns the seed catalog of Aquamor products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Seed product catalog covering refills, bottled packs, and accessories
	products := []models.Product{
		{
			ID:             1,
			Name:           "20L Purified Water Container",
			Category:       "refill",
			Volume:         "20L",
			PriceSwap:      200.00,
			PriceNewBottle: 1000.00,
			IsAvailable:    true,
		},
		{
			ID:             2,
			Name:           "10L Purified Water Container",
			Category:       "refill",
			Volume:         "10L",
			PriceSwap:      100.00,
			PriceNewBottle: 600.00,
			IsAvailable:    true,
		},
		{
			ID:             3,
			Name:           "5L Purified Water Bottle",
			Category:       "refill",
			Volume:         "5L",
			PriceSwap:      50.00,
			PriceNewBottle: 300.00,
			IsAvailable:    true,
		},
		{
			ID:             4,
			Name:           "500ml Bottled Water (Carton of 24)",
			Category:       "bottled_pack",
			Volume:         "500ml",
			PriceSwap:      0.00,
			PriceNewBottle: 550.00,
			IsAvailable:    true,
		},
		{
			ID:             5,
			Name:           "1.5L Bottled Water (Carton of 12)",
			Category:       "bottled_pack",
			Volume:         "1.5L",
			PriceSwap:      0.00,
			PriceNewBottle: 750.00,
			IsAvailable:    true,
		},
		{
			ID:             6,
			Name:           "Portable Automatic Water Pump",
			Category:       "accessory",
			Volume:         "",
			PriceSwap:      0.00,
			PriceNewBottle: 850.00,
			IsAvailable:    true,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}