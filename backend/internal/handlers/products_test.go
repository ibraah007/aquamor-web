package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"aquamor-backend/backend/internal/models"
)

func TestGetProductsReturnsOnlyAvailableItems(t *testing.T) {
	recorder := httptest.NewRecorder()
	GetProducts(recorder, httptest.NewRequest(http.MethodGet, "/api/v1/products", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", recorder.Code)
	}

	var products []models.Product
	if err := json.Unmarshal(recorder.Body.Bytes(), &products); err != nil {
		t.Fatalf("decode products: %v; body: %s", err, recorder.Body.String())
	}
	if len(products) != 5 {
		t.Fatalf("products = %d, want 5", len(products))
	}
	for _, product := range products {
		if !product.IsAvailable {
			t.Errorf("product %d (%s) is unavailable but was returned", product.ID, product.Name)
		}
		if product.PriceSwap == 0 && product.PriceNewBottle == 0 {
			t.Errorf("product %d (%s) has no price but was returned", product.ID, product.Name)
		}
	}
}

func TestGetProductsRejectsWrongMethod(t *testing.T) {
	recorder := httptest.NewRecorder()
	GetProducts(recorder, httptest.NewRequest(http.MethodPost, "/api/v1/products", nil))

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", recorder.Code)
	}
}
