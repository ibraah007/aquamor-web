package requestlog

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogsMethodPathStatusSizeAndIP(t *testing.T) {
	var buf bytes.Buffer
	handler := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("hello"))
	}), log.New(&buf, "", 0))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/orders?token=secret", nil)
	req.RemoteAddr = "198.51.100.7:1234"
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status %d, want 201", rec.Code)
	}
	if rec.Body.String() != "hello" {
		t.Fatalf("body %q, want hello", rec.Body.String())
	}

	out := buf.String()
	for _, want := range []string{
		"method=POST",
		`path="/api/v1/orders"`,
		"status=201",
		"bytes=5",
		"duration=",
		"ip=198.51.100.7",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("log %q does not contain %q", out, want)
		}
	}
	if strings.Contains(out, "secret") {
		t.Fatalf("log contains the query string: %q", out)
	}
}

func TestDefaultsTo200WhenWriteHeaderIsOmitted(t *testing.T) {
	var buf bytes.Buffer
	handler := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}), log.New(&buf, "", 0))

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)

	if out := buf.String(); !strings.Contains(out, "status=200") || !strings.Contains(out, "bytes=2") {
		t.Fatalf("unexpected log %q", out)
	}
}

func TestUsesForwardedClientIP(t *testing.T) {
	var buf bytes.Buffer
	handler := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), log.New(&buf, "", 0))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:8080"
	req.Header.Set("X-Forwarded-For", "10.0.0.1, 203.0.113.9")
	handler.ServeHTTP(httptest.NewRecorder(), req)

	if out := buf.String(); !strings.Contains(out, "ip=203.0.113.9") {
		t.Fatalf("unexpected log %q", out)
	}
}
