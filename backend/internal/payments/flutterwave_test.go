package payments

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func withTestServer(t *testing.T, handler http.HandlerFunc) Config {
	t.Helper()
	server := httptest.NewServer(handler)
	oldBase, oldClient := BaseURL, HTTPClient
	BaseURL = server.URL
	HTTPClient = server.Client()
	t.Cleanup(func() {
		BaseURL = oldBase
		HTTPClient = oldClient
		server.Close()
	})
	return Config{SecretKey: "sk_test_123", WebhookSecret: "webhook_hash", PublicBaseURL: "https://aquamor.test"}
}

func TestCreateCheckout(t *testing.T) {
	var got map[string]interface{}
	cfg := withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v3/payments" {
			t.Errorf("path = %s, want /v3/payments", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer sk_test_123" {
			t.Errorf("authorization = %q", r.Header.Get("Authorization"))
		}
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Errorf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"data":   map[string]string{"link": "https://checkout.flutterwave.test/pay/abc"},
		})
	})

	link, err := cfg.CreateCheckout(context.Background(), CheckoutRequest{
		TxRef:         "aquamor-1",
		Amount:        400,
		CustomerName:  "Amina Wanjiku",
		CustomerPhone: "0722000111",
		CustomerEmail: "amina@example.com",
		Description:   "Aquamor order #1",
	})
	if err != nil {
		t.Fatalf("CreateCheckout: %v", err)
	}
	if link != "https://checkout.flutterwave.test/pay/abc" {
		t.Fatalf("link = %q", link)
	}
	if got["tx_ref"] != "aquamor-1" || got["currency"] != "KES" || got["amount"] != float64(400) {
		t.Errorf("payload = %#v", got)
	}
	if got["redirect_url"] != "https://aquamor.test/payment.html" {
		t.Errorf("redirect_url = %v", got["redirect_url"])
	}
	if options, _ := got["payment_options"].(string); !strings.Contains(options, "mpesa") {
		t.Errorf("payment_options = %v, want mpesa enabled", got["payment_options"])
	}
	customer, _ := got["customer"].(map[string]interface{})
	if customer["email"] != "amina@example.com" || customer["name"] != "Amina Wanjiku" {
		t.Errorf("customer = %#v", customer)
	}
}

func TestCreateCheckoutRejectsProviderError(t *testing.T) {
	cfg := withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid authorization key"})
	})

	_, err := cfg.CreateCheckout(context.Background(), CheckoutRequest{TxRef: "aquamor-2", Amount: 50, CustomerEmail: "a@example.com"})
	if err == nil {
		t.Fatal("CreateCheckout succeeded, want an error")
	}
	if !strings.Contains(err.Error(), "401") {
		t.Errorf("error = %v, want status code", err)
	}
}

func TestVerify(t *testing.T) {
	cfg := withTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v3/transactions/987/verify" {
			t.Errorf("path = %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"id":       987,
				"tx_ref":   "aquamor-3",
				"amount":   400,
				"currency": "KES",
				"status":   "successful",
			},
		})
	})

	verified, err := cfg.Verify(context.Background(), "987")
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if verified.Status != "successful" || verified.TxRef != "aquamor-3" || verified.TransactionID != "987" || verified.Amount != 400 || verified.Currency != "KES" {
		t.Errorf("verified = %#v", verified)
	}
}

func TestVerifyRequiresTransactionID(t *testing.T) {
	cfg := Config{SecretKey: "sk_test_123"}
	if _, err := cfg.Verify(context.Background(), "  "); err == nil {
		t.Fatal("Verify succeeded without a transaction id")
	}
}

func TestWebhookMatches(t *testing.T) {
	cfg := Config{WebhookSecret: "secret-hash"}
	if !cfg.WebhookMatches("secret-hash") {
		t.Error("WebhookMatches rejected the configured secret")
	}
	if cfg.WebhookMatches("wrong") {
		t.Error("WebhookMatches accepted a wrong secret")
	}
	if cfg.WebhookMatches("") {
		t.Error("WebhookMatches accepted an empty header")
	}
}
