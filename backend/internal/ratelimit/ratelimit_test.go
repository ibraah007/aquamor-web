package ratelimit

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAllowWithinAndOverLimit(t *testing.T) {
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	l := New(2, time.Minute)
	l.now = func() time.Time { return now }

	if ok, _ := l.Allow("1.2.3.4"); !ok {
		t.Fatal("first request should be allowed")
	}
	if ok, _ := l.Allow("1.2.3.4"); !ok {
		t.Fatal("second request should be allowed")
	}
	ok, retry := l.Allow("1.2.3.4")
	if ok {
		t.Fatal("third request should be blocked")
	}
	if retry <= 0 || retry > time.Minute {
		t.Fatalf("unexpected retry duration %v", retry)
	}
}

func TestResetsAfterWindow(t *testing.T) {
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	l := New(1, time.Minute)
	l.now = func() time.Time { return now }

	if ok, _ := l.Allow("key"); !ok {
		t.Fatal("first request should be allowed")
	}
	if ok, _ := l.Allow("key"); ok {
		t.Fatal("second request should be blocked")
	}

	now = now.Add(time.Minute)
	if ok, _ := l.Allow("key"); !ok {
		t.Fatal("request after the window should be allowed")
	}
}

func TestSeparateKeysHaveSeparateWindows(t *testing.T) {
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	l := New(1, time.Minute)
	l.now = func() time.Time { return now }

	if ok, _ := l.Allow("a"); !ok {
		t.Fatal("first key should be allowed")
	}
	if ok, _ := l.Allow("b"); !ok {
		t.Fatal("second key should not inherit the first key's count")
	}
}

func TestConcurrentAllowsRespectLimit(t *testing.T) {
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	l := New(50, time.Minute)
	l.now = func() time.Time { return now }

	var allowed atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if ok, _ := l.Allow("same"); ok {
				allowed.Add(1)
			}
		}()
	}
	wg.Wait()

	if got := allowed.Load(); got != 50 {
		t.Fatalf("allowed %d requests, want 50", got)
	}
}

func TestMiddlewareReturns429(t *testing.T) {
	now := time.Date(2026, 11, 13, 12, 0, 0, 0, time.UTC)
	l := New(2, time.Minute)
	l.now = func() time.Time { return now }

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})
	handler := l.Middleware(next)

	for i := 0; i < 2; i++ {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/orders", nil)
		req.RemoteAddr = "203.0.113.5:4321"
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusCreated {
			t.Fatalf("request %d: status %d, want 201", i+1, rec.Code)
		}
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/orders", nil)
	req.RemoteAddr = "203.0.113.5:4321"
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("status %d, want 429", rec.Code)
	}
	if got := rec.Header().Get("Retry-After"); got == "" {
		t.Fatal("Retry-After header is missing")
	}
	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "application/json") {
		t.Fatalf("Content-Type %q, want application/json", got)
	}
	if !strings.Contains(rec.Body.String(), "Too many requests") {
		t.Fatalf("unexpected body %q", rec.Body.String())
	}
}

func TestClientIP(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		forwarded  string
		want       string
	}{
		{name: "remote address", remoteAddr: "198.51.100.7:9999", want: "198.51.100.7"},
		{name: "forwarded last value", remoteAddr: "127.0.0.1:8080", forwarded: "10.0.0.1, 203.0.113.9", want: "203.0.113.9"},
		{name: "forwarded single value", remoteAddr: "127.0.0.1:8080", forwarded: "203.0.113.9", want: "203.0.113.9"},
		{name: "malformed remote address", remoteAddr: "not-a-host", want: "not-a-host"},
		{name: "empty remote address", remoteAddr: "", want: "unknown"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.RemoteAddr = tc.remoteAddr
			if tc.forwarded != "" {
				req.Header.Set("X-Forwarded-For", tc.forwarded)
			}
			if got := ClientIP(req); got != tc.want {
				t.Fatalf("ClientIP() = %q, want %q", got, tc.want)
			}
		})
	}
}
