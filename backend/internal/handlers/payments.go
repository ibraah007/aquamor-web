package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/models"
	"aquamor-backend/backend/internal/payments"
)

const paymentProvider = "flutterwave"

// PaymentConfig tells the site whether the Pay online button should be shown.
func PaymentConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	cfg := payments.FromEnv()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"enabled":  cfg.Enabled(),
		"provider": paymentProvider,
	})
}

// CreateCheckout saves the order and returns a hosted payment page URL.
func CreateCheckout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	cfg := payments.FromEnv()
	if !cfg.Enabled() {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"code":  "payments_not_configured",
			"error": "Online payment is being set up. Please send your order on WhatsApp and we will confirm the total.",
		})
		return
	}

	req, ok := decodeOrderForPayment(w, r)
	if !ok {
		return
	}
	if !validEmail(req.CustomerEmail) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "A valid email address is needed for the payment receipt"})
		return
	}

	lines, totalAmount, err := priceOrder(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	reference, err := newPaymentReference()
	if err != nil {
		log.Printf("create checkout: reference: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not start the payment. Please try again."})
		return
	}

	order, err := insertOrder(req, lines, totalAmount, paymentMeta{
		CustomerEmail:    req.CustomerEmail,
		PaymentProvider:  paymentProvider,
		PaymentReference: reference,
		PaymentStatus:    "pending",
	})
	if err != nil {
		log.Printf("create checkout: save order: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save order"})
		return
	}

	link, err := cfg.CreateCheckout(r.Context(), payments.CheckoutRequest{
		TxRef:         reference,
		Amount:        totalAmount,
		CustomerName:  req.CustomerName,
		CustomerPhone: req.CustomerPhone,
		CustomerEmail: req.CustomerEmail,
		Description:   fmt.Sprintf("Aquamor order #%d", order.ID),
	})
	if err != nil {
		log.Printf("create checkout: provider: %v", err)
		if _, updateErr := db.DB.Exec(`UPDATE orders SET payment_status = 'failed' WHERE id = ?`, order.ID); updateErr != nil {
			log.Printf("create checkout: mark failed: %v", updateErr)
		}
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(map[string]string{
			"error":    "Could not start the payment. Please try again or send your order on WhatsApp.",
			"order_id": strconv.Itoa(order.ID),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Checkout created",
		"order":       order,
		"payment_url": link,
	})
}

// VerifyPayment confirms a payment after the customer returns from the hosted
// checkout. payment.html calls this with the reference and transaction id.
func VerifyPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	reference := strings.TrimSpace(r.URL.Query().Get("reference"))
	transactionID := strings.TrimSpace(r.URL.Query().Get("transaction_id"))
	if reference == "" || transactionID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "reference and transaction_id are required"})
		return
	}

	cfg := payments.FromEnv()
	if !cfg.Enabled() {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "Online payment is not configured"})
		return
	}

	var orderID int
	var totalAmount float64
	var paymentStatus string
	err := db.DB.QueryRow(`SELECT id, total_amount, payment_status FROM orders WHERE payment_reference = ?`, reference).Scan(&orderID, &totalAmount, &paymentStatus)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Payment reference not found"})
		return
	}
	if err != nil {
		log.Printf("verify payment: lookup: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not check the payment"})
		return
	}
	if paymentStatus == "paid" {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "paid", "order_id": orderID})
		return
	}

	verified, err := cfg.Verify(r.Context(), transactionID)
	if err != nil {
		log.Printf("verify payment: provider: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not verify the payment with the provider"})
		return
	}
	if verified.TxRef != reference {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Payment reference does not match"})
		return
	}
	if !strings.EqualFold(verified.Currency, "KES") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unexpected payment currency"})
		return
	}
	if math.Abs(verified.Amount-totalAmount) > 0.01 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Payment amount does not match the order"})
		return
	}

	if !strings.EqualFold(verified.Status, "successful") {
		if _, updateErr := db.DB.Exec(`UPDATE orders SET payment_status = 'failed' WHERE id = ? AND payment_status <> 'paid'`, orderID); updateErr != nil {
			log.Printf("verify payment: mark failed: %v", updateErr)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "failed", "order_id": orderID})
		return
	}

	if err := markOrderPaid(orderID, verified.TransactionID); err != nil {
		log.Printf("verify payment: mark paid: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Payment succeeded but the order could not be updated"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "paid",
		"order_id": orderID,
	})
}

// PaymentWebhook is the provider's server-to-server confirmation. It is a
// fallback for customers who close the browser before returning to the site.
func PaymentWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	cfg := payments.FromEnv()
	if cfg.WebhookSecret == "" {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "Webhook is not configured"})
		return
	}
	if !cfg.WebhookMatches(r.Header.Get("verif-hash")) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid webhook signature"})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxOrderBody)
	var event struct {
		Event string `json:"event"`
		Data  struct {
			ID       int64   `json:"id"`
			TxRef    string  `json:"tx_ref"`
			Status   string  `json:"status"`
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid webhook payload"})
		return
	}
	if event.Event != "charge.completed" || !strings.EqualFold(event.Data.Status, "successful") {
		json.NewEncoder(w).Encode(map[string]string{"status": "ignored"})
		return
	}
	if event.Data.TxRef == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Webhook is missing tx_ref"})
		return
	}

	var orderID int
	var totalAmount float64
	err := db.DB.QueryRow(`SELECT id, total_amount FROM orders WHERE payment_reference = ?`, event.Data.TxRef).Scan(&orderID, &totalAmount)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("payment webhook: unknown reference %q", event.Data.TxRef)
		json.NewEncoder(w).Encode(map[string]string{"status": "ignored"})
		return
	}
	if err != nil {
		log.Printf("payment webhook: lookup: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not check the payment"})
		return
	}
	if !strings.EqualFold(event.Data.Currency, "KES") || math.Abs(event.Data.Amount-totalAmount) > 0.01 {
		log.Printf("payment webhook: amount mismatch for order %d (got %.2f %s, want %.2f KES)", orderID, event.Data.Amount, event.Data.Currency, totalAmount)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Payment amount or currency does not match the order"})
		return
	}

	if err := markOrderPaid(orderID, strconv.FormatInt(event.Data.ID, 10)); err != nil {
		log.Printf("payment webhook: mark paid: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not update the order"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func markOrderPaid(orderID int, receipt string) error {
	_, err := db.DB.Exec(
		`UPDATE orders SET payment_status = 'paid', order_status = 'processing', mpesa_receipt_number = ?, paid_at = CURRENT_TIMESTAMP WHERE id = ? AND payment_status <> 'paid'`,
		receipt, orderID,
	)
	return err
}

func decodeOrderForPayment(w http.ResponseWriter, r *http.Request) (models.OrderRequest, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, maxOrderBody)
	decoder := json.NewDecoder(r.Body)
	var req models.OrderRequest
	if err := decoder.Decode(&req); err != nil {
		writeDecodeError(w, err)
		return req, false
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		writeDecodeError(w, err)
		return req, false
	}

	req.CustomerName = strings.TrimSpace(req.CustomerName)
	req.CustomerPhone = strings.TrimSpace(req.CustomerPhone)
	req.CustomerEmail = strings.TrimSpace(req.CustomerEmail)
	req.DeliveryEstate = strings.TrimSpace(req.DeliveryEstate)
	req.DeliveryAddress = strings.TrimSpace(req.DeliveryAddress)

	if err := validateOrder(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return req, false
	}
	return req, true
}

func writeDecodeError(w http.ResponseWriter, err error) {
	var tooLarge *http.MaxBytesError
	if errors.As(err, &tooLarge) {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		json.NewEncoder(w).Encode(map[string]string{"error": "Request body is too large"})
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
}

func newPaymentReference() (string, error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("aquamor-%d-%s", time.Now().UTC().Unix(), hex.EncodeToString(buf)), nil
}

func validEmail(value string) bool {
	if len(value) < 5 || len(value) > 120 || strings.ContainsAny(value, " \t\r\n") {
		return false
	}
	at := strings.LastIndex(value, "@")
	if at <= 0 || at >= len(value)-3 {
		return false
	}
	return strings.Contains(value[at+1:], ".")
}
