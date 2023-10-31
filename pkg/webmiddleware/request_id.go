package webmiddleware

import (
	"crypto/rand"
	"net/http"

	ulid "github.com/oklog/ulid/v2"
)

const requestIDHeader = "X-Request-ID"

// RequestID returns a middleware that inserts a request ID into each request.
func RequestID() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = ulid.MustNew(ulid.Now(), rand.Reader).String()
				r.Header.Set(requestIDHeader, requestID)
			}
			w.Header().Set(requestIDHeader, requestID)
			next.ServeHTTP(w, r)
		})
	}
}
