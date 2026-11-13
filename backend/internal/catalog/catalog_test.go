package catalog

import "testing"

func TestConfirmedPrices(t *testing.T) {
	want := []struct {
		id        int
		name      string
		swap      float64
		newBottle float64
	}{
		{1, "20L Purified Water Refill", 200, 0},
		{2, "10L Purified Water Refill", 100, 0},
		{3, "5L Purified Water Refill", 50, 0},
		{4, "20L Hard Bottle + Water", 0, 1400},
		{5, "20L Regular Bottle", 0, 450},
	}

	for _, w := range want {
		product, ok := ByID(w.id)
		if !ok {
			t.Fatalf("product %d (%s) should be available", w.id, w.name)
		}
		if product.Name != w.name {
			t.Errorf("product %d name = %q, want %q", w.id, product.Name, w.name)
		}
		if product.PriceSwap != w.swap {
			t.Errorf("product %d swap price = %.2f, want %.2f", w.id, product.PriceSwap, w.swap)
		}
		if product.PriceNewBottle != w.newBottle {
			t.Errorf("product %d new-bottle price = %.2f, want %.2f", w.id, product.PriceNewBottle, w.newBottle)
		}
	}
}

func TestUnconfirmedProductsAreNotSold(t *testing.T) {
	// These items stay in the catalog without a price, so ByID must not expose
	// them and the API must never quote them.
	for _, id := range []int{6, 7, 8, 999} {
		if product, ok := ByID(id); ok {
			t.Errorf("product %d is unavailable but ByID returned %+v", id, product)
		}
	}
}

func TestAllReturnsACopy(t *testing.T) {
	first := All()
	if len(first) == 0 {
		t.Fatal("catalog is empty")
	}
	first[0].Name = "changed"

	second := All()
	if second[0].Name == "changed" {
		t.Fatal("All returned the backing slice instead of a copy")
	}
}
