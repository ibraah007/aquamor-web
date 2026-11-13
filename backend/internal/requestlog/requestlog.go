// Package requestlog logs one line per HTTP request without recording customer
// data. It records the method, path, status, response size, duration and client
// address; query strings and headers are deliberately left out.
package requestlog

import (
	"log"
	"net/http"
	"time"

	"aquamor-backend/backend/internal/ratelimit"
)

// Middleware wraps next and logs each completed request with logger. A nil
// logger uses the standard logger.
func Middleware(next http.Handler, logger *log.Logger) http.Handler {
	if logger == nil {
		logger = log.Default()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		status := rec.status
		if status == 0 {
			status = http.StatusOK
		}
		logger.Printf("request method=%s path=%q status=%d bytes=%d duration=%s ip=%s",
			r.Method,
			r.URL.Path,
			status,
			rec.bytes,
			time.Since(start).Round(time.Microsecond),
			ratelimit.ClientIP(r),
		)
	})
}

// responseWriter records the status code and body size while passing writes
// through. Unwrap keeps http.ResponseController features such as Flush working.
type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *responseWriter) WriteHeader(code int) {
	if w.status != 0 {
		return
	}
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}
