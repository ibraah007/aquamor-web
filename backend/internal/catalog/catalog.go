// Package catalog holds the single source of truth for the product list and
// prices used by the API. Only prices confirmed by the client belong here.
package catalog

import "aquamor-backend/backend/internal/models"

var products = []models.Product{
	{
		ID:             1,
		Name:           "20L Purified Water Refill",
		Category:       "refill",
		Volume:         "20L",
		PriceSwap:      200.00,
		PriceNewBottle: 0,
		IsAvailable:    true,
	},
	{
		ID:             2,
		Name:           "10L Purified Water Refill",
		Category:       "refill",
		Volume:         "10L",
		PriceSwap:      100.00,
		PriceNewBottle: 0,
		IsAvailable:    true,
	},
	{
		ID:             3,
		Name:           "5L Purified Water Refill",
		Category:       "refill",
		Volume:         "5L",
		PriceSwap:      50.00,
		PriceNewBottle: 0,
		IsAvailable:    true,
	},
	{
		ID:             4,
		Name:           "20L Hard Bottle + Water",
		Category:       "container",
		Volume:         "20L",
		PriceSwap:      0,
		PriceNewBottle: 1400.00,
		IsAvailable:    true,
	},
	{
		ID:             5,
		Name:           "20L Regular Bottle",
		Category:       "container",
		Volume:         "20L",
		PriceSwap:      0,
		PriceNewBottle: 450.00,
		IsAvailable:    true,
	},
	// Prices for the items below have not been confirmed by the client yet.
	// Keep them in the catalog but unavailable so the API cannot quote them.
	{
		ID:             6,
		Name:           "500ml Bottled Water (Carton of 24)",
		Category:       "bottled_pack",
		Volume:         "500ml",
		PriceSwap:      0,
		PriceNewBottle: 0,
		IsAvailable:    false,
	},
	{
		ID:             7,
		Name:           "1.5L Bottled Water (Carton of 12)",
		Category:       "bottled_pack",
		Volume:         "1.5L",
		PriceSwap:      0,
		PriceNewBottle: 0,
		IsAvailable:    false,
	},
	{
		ID:             8,
		Name:           "Portable Automatic Water Pump",
		Category:       "accessory",
		Volume:         "",
		PriceSwap:      0,
		PriceNewBottle: 0,
		IsAvailable:    false,
	},
}

// All returns a copy of the catalog.
func All() []models.Product {
	out := make([]models.Product, len(products))
	copy(out, products)
	return out
}

// ByID finds an available product by its ID.
func ByID(id int) (models.Product, bool) {
	for _, p := range products {
		if p.ID == id && p.IsAvailable {
			return p, true
		}
	}
	return models.Product{}, false
}
