package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"aquamor-backend/backend/internal/db"
	"aquamor-backend/backend/internal/handlers"
)

// The site uses inline JSON-LD, a few inline style attributes and one inline
// font-loading handler, so script/style need 'unsafe-inline'. Everything else
// is limited to the site itself and Google Fonts.
const contentSecurityPolicy = "default-src 'self'; " +
	"base-uri 'self'; " +
	"object-src 'none'; " +
	"frame-ancestors 'none'; " +
	"frame-src https://www.google.com https://maps.google.com; " +
	"form-action 'self'; " +
	"img-src 'self' data:; " +
	"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; " +
	"font-src 'self' https://fonts.gstatic.com; " +
	"script-src 'self' 'unsafe-inline'; " +
	"connect-src 'self'"

func main() {
	addr := envOr("ADDR", "127.0.0.1:8080")
	dbPath := envOr("DB_PATH", "./aquamor.db")
	webDir := envOr("WEB_DIR", "web")

	if _, err := os.Stat(filepath.Join(webDir, "index.html")); err != nil {
		log.Fatalf("web directory %q does not contain index.html: %v", webDir, err)
	}

	if err := db.InitDB(dbPath); err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()
	log.Printf("database ready at %s", dbPath)

	mux := http.NewServeMux()

	// API routes. Go's ServeMux picks the longest matching pattern, so these
	// win over the static file handler at "/".
	mux.HandleFunc("/api/v1/health", health)
	mux.HandleFunc("/api/v1/products", handlers.GetProducts)
	mux.HandleFunc("/api/v1/orders", handlers.CreateOrder)
	mux.HandleFunc("/api/v1/orders/all", handlers.GetOrders)
	mux.HandleFunc("/api/v1/payments/config", handlers.PaymentConfig)
	mux.HandleFunc("/api/v1/payments/checkout", handlers.CreateCheckout)
	mux.HandleFunc("/api/v1/payments/verify", handlers.VerifyPayment)
	mux.HandleFunc("/api/v1/payments/webhook", handlers.PaymentWebhook)

	// Static site: web/index.html plus its assets. Run the server from the
	// repository root, or set WEB_DIR to the absolute path.
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	srv := &http.Server{
		Addr:              addr,
		Handler:           securityHeaders(cacheControl(mux)),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      5 * time.Minute, // allows a slow video download
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Aquamor listening on http://%s (web: %s)", addr, webDir)
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %v", err)
		}
	case <-ctx.Done():
		log.Println("shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown: %v", err)
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		log.Printf("health: database: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Database is not reachable"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if r.Method == http.MethodHead {
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "Aquamor Water API is running"})
}

// securityHeaders adds browser hardening that is safe for this static site and
// its JSON API. HSTS is only sent when the request arrived over HTTPS (directly
// or through the reverse proxy), so local plain-HTTP development is unaffected.
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		h.Set("Permissions-Policy", "camera=(), microphone=(), geolocation=(self), payment=()")
		h.Set("Content-Security-Policy", contentSecurityPolicy)
		if r.URL.Path == "/admin.html" {
			h.Set("X-Robots-Tag", "noindex, nofollow")
		}
		if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
			h.Set("Strict-Transport-Security", "max-age=31536000")
		}
		next.ServeHTTP(w, r)
	})
}

// cacheControl keeps HTML fresh while allowing browsers to reuse versioned
// static assets for a short period. The filenames are not content-hashed, so
// the lifetime stays short enough that a replaced photo is picked up quickly.
func cacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.Header().Set("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/admin.html" {
			w.Header().Set("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
			return
		}

		switch strings.ToLower(filepath.Ext(r.URL.Path)) {
		case ".css", ".js":
			w.Header().Set("Cache-Control", "public, max-age=3600")
		case ".jpg", ".jpeg", ".png", ".webp", ".svg", ".mp4", ".woff", ".woff2":
			w.Header().Set("Cache-Control", "public, max-age=86400")
		default:
			w.Header().Set("Cache-Control", "no-cache")
		}
		next.ServeHTTP(w, r)
	})
}

func envOr(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
