package webmiddleware

import (
	"net"
	"net/http"

	"google.golang.org/grpc/peer"
)

// Peer sets the remote address as a peer in the request context.
func Peer() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if addr, err := net.ResolveTCPAddr("tcp", r.RemoteAddr); err == nil {
				ctx = peer.NewContext(ctx, &peer.Peer{
					Addr: addr,
				})
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
