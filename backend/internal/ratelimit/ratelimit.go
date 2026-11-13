// Package ratelimit provides a small in-memory fixed-window limiter for
// public write endpoints.
package ratelimit

import (
	"encoding/json"
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type entry struct {
	count int
	reset time.Time
}

// Limiter allows at most limit requests per key (usually a client IP) in each
// window. It is safe for concurrent use. State is per process and is lost on
// restart, which is acceptable for abuse protection.
type Limiter struct {
	mu        sync.Mutex
	limit     int
	window    time.Duration
	now       func() time.Time
	entries   map[string]*entry
	lastSweep time.Time
}

// New returns a limiter that permits limit requests per key per window.
func New(limit int, window time.Duration) *Limiter {
	if limit < 1 {
		limit = 1
	}
	return &Limiter{
		limit:   limit,
		window:  window,
		now:     time.Now,
		entries: make(map[string]*entry),
	}
}

// Allow records a request for key and reports whether it is permitted. When it
// is not, the returned duration is the time until the current window resets.
func (l *Limiter) Allow(key string) (bool, time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	l.sweepLocked(now)

	current, ok := l.entries[key]
	if !ok || !now.Before(current.reset) {
		l.entries[key] = &entry{count: 1, reset: now.Add(l.window)}
		return true, 0
	}
	if current.count >= l.limit {
		return false, current.reset.Sub(now)
	}
	current.count++
	return true, 0
}

// sweepLocked drops expired keys once per window so the map cannot grow without
// bound. It must be called with the mutex held.
func (l *Limiter) sweepLocked(now time.Time) {
	if !l.lastSweep.IsZero() && now.Sub(l.lastSweep) < l.window {
		return
	}
	for key, current := range l.entries {
		if !now.Before(current.reset) {
			delete(l.entries, key)
		}
	}
	l.lastSweep = now
}

// Middleware rejects requests over the limit with 429 and a JSON body.
func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed, retryAfter := l.Allow(ClientIP(r))
		if !allowed {
			seconds := int(math.Ceil(retryAfter.Seconds()))
			if seconds < 1 {
				seconds = 1
			}
			w.Header().Set("Retry-After", strconv.Itoa(seconds))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{"error": "Too many requests. Please try again shortly."})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ClientIP returns the requesting client address. It uses the last
// X-Forwarded-For value when a reverse proxy supplied one, and otherwise the
// connection's remote address. The production server binds to loopback behind
// Caddy, which appends the client address to this header; taking the last value
// ignores any spoofed entries a client may have sent.
func ClientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		if comma := strings.LastIndexByte(forwarded, ','); comma >= 0 {
			forwarded = forwarded[comma+1:]
		}
		if ip := strings.TrimSpace(forwarded); ip != "" {
			return ip
		}
	}
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return host
	}
	if r.RemoteAddr == "" {
		return "unknown"
	}
	return r.RemoteAddr
}
