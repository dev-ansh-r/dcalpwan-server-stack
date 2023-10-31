package webmiddleware

import (
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

// Namespace is middleware that sets the namespace.
func Namespace(value string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(log.NewContextWithField(r.Context(), "namespace", value)))
		})
	}
}
