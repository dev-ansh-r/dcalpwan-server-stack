package webmiddleware

import (
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/fillcontext"
)

// FillContext returns a middleware that fills global context into a call context.
func FillContext(fillers ...fillcontext.Filler) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		if len(fillers) == 0 {
			return next
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, fill := range fillers {
				ctx = fill(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
