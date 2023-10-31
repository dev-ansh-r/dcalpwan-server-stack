package webmiddleware

import (
	"net/http"
	"net/textproto"
	"strings"

	"google.golang.org/grpc/metadata"
)

// Metadata returns a middleware that sets gRPC metadata from the request headers.
func Metadata(keys ...string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			header := textproto.MIMEHeader(r.Header)
			md := make(metadata.MD)
			for _, key := range keys {
				if value := header.Values(key); value != nil {
					md[strings.ToLower(key)] = value
				}
			}
			if len(md) > 0 {
				if existingMD, ok := metadata.FromIncomingContext(ctx); ok {
					md = metadata.Join(existingMD, md)
				}
				ctx = metadata.NewIncomingContext(ctx, md)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
