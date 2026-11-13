// Package payments talks to a hosted checkout provider. Flutterwave is the
// first provider; the small surface (create a checkout, verify a transaction,
// check a webhook signature) keeps another provider easy to add.
package payments

import (
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// BaseURL and HTTPClient are variables so tests can point them at a local
// server. Production uses the Flutterwave API.
var (
	BaseURL    = "https://api.flutterwave.com"
	HTTPClient = &http.Client{Timeout: 20 * time.Second}
)

const maxResponseBody = 1 << 20

// Config holds the settings needed to create and verify hosted payments.
type Config struct {
	SecretKey     string
	WebhookSecret string
	PublicBaseURL string
}

// FromEnv reads the payment settings from the environment.
func FromEnv() Config {
	return Config{
		SecretKey:     strings.TrimSpace(os.Getenv("FLUTTERWAVE_SECRET_KEY")),
		WebhookSecret: strings.TrimSpace(os.Getenv("FLUTTERWAVE_WEBHOOK_SECRET")),
		PublicBaseURL: strings.TrimRight(strings.TrimSpace(os.Getenv("PUBLIC_BASE_URL")), "/"),
	}
}

// Enabled reports whether enough settings exist to create a checkout and
// receive its verification callback.
func (c Config) Enabled() bool {
	return c.SecretKey != "" && c.WebhookSecret != "" && c.PublicBaseURL != ""
}

// CheckoutRequest describes one hosted checkout.
type CheckoutRequest struct {
	TxRef         string
	Amount        float64
	CustomerName  string
	CustomerPhone string
	CustomerEmail string
	Description   string
}

// Verification is the provider's answer for a completed or failed transaction.
type Verification struct {
	Status        string
	TxRef         string
	TransactionID string
	Amount        float64
	Currency      string
}

type envelope struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// CreateCheckout asks Flutterwave for a hosted payment page and returns its URL.
func (c Config) CreateCheckout(ctx context.Context, req CheckoutRequest) (string, error) {
	if c.SecretKey == "" {
		return "", errors.New("payment provider is not configured")
	}
	if c.PublicBaseURL == "" {
		return "", errors.New("PUBLIC_BASE_URL is not configured")
	}

	payload := map[string]interface{}{
		"tx_ref":          req.TxRef,
		"amount":          req.Amount,
		"currency":        "KES",
		"redirect_url":    c.PublicBaseURL + "/payment.html",
		"payment_options": "card,mpesa,mobilemoney",
		"customer": map[string]string{
			"email":       req.CustomerEmail,
			"phonenumber": req.CustomerPhone,
			"name":        req.CustomerName,
		},
		"customizations": map[string]string{
			"title":       "Aquamor order",
			"description": req.Description,
		},
	}

	body, err := c.do(ctx, http.MethodPost, "/v3/payments", payload)
	if err != nil {
		return "", err
	}

	var env envelope
	if err := json.Unmarshal(body, &env); err != nil {
		return "", fmt.Errorf("decode checkout response: %w", err)
	}
	if !strings.EqualFold(env.Status, "success") {
		return "", fmt.Errorf("checkout rejected: %s", safeMessage(env.Message))
	}
	var data struct {
		Link string `json:"link"`
	}
	if err := json.Unmarshal(env.Data, &data); err != nil || data.Link == "" {
		return "", errors.New("checkout response did not include a payment link")
	}
	return data.Link, nil
}

// Verify fetches the transaction from Flutterwave by its transaction id.
func (c Config) Verify(ctx context.Context, transactionID string) (Verification, error) {
	if c.SecretKey == "" {
		return Verification{}, errors.New("payment provider is not configured")
	}
	if strings.TrimSpace(transactionID) == "" {
		return Verification{}, errors.New("transaction id is required")
	}

	body, err := c.do(ctx, http.MethodGet, "/v3/transactions/"+url.PathEscape(transactionID)+"/verify", nil)
	if err != nil {
		return Verification{}, err
	}

	var env envelope
	if err := json.Unmarshal(body, &env); err != nil {
		return Verification{}, fmt.Errorf("decode verification response: %w", err)
	}
	if !strings.EqualFold(env.Status, "success") {
		return Verification{}, fmt.Errorf("verification rejected: %s", safeMessage(env.Message))
	}
	var data struct {
		ID       int64   `json:"id"`
		TxRef    string  `json:"tx_ref"`
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
		Status   string  `json:"status"`
	}
	if err := json.Unmarshal(env.Data, &data); err != nil {
		return Verification{}, fmt.Errorf("decode transaction: %w", err)
	}
	return Verification{
		Status:        data.Status,
		TxRef:         data.TxRef,
		TransactionID: strconv.FormatInt(data.ID, 10),
		Amount:        data.Amount,
		Currency:      data.Currency,
	}, nil
}

// WebhookMatches reports whether the verif-hash header matches the configured
// webhook secret.
func (c Config) WebhookMatches(header string) bool {
	if c.WebhookSecret == "" || header == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(header), []byte(c.WebhookSecret)) == 1
}

func (c Config) do(ctx context.Context, method, path string, payload interface{}) ([]byte, error) {
	var reader io.Reader
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("encode request: %w", err)
		}
		reader = bytes.NewReader(encoded)
	}

	req, err := http.NewRequestWithContext(ctx, method, BaseURL+path, reader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("payment request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBody))
	if err != nil {
		return nil, fmt.Errorf("read payment response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("payment provider returned %d: %s", resp.StatusCode, safeMessage(string(body)))
	}
	return body, nil
}

// safeMessage keeps provider errors useful in logs without dumping a whole
// HTML or JSON body into the response.
func safeMessage(message string) string {
	message = strings.TrimSpace(message)
	if len(message) > 200 {
		message = message[:200] + "..."
	}
	return message
}
