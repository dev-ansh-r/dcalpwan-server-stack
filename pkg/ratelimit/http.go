package ratelimit

import (
	"net"
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

const httpXRealIPHeader = "X-Real-IP"

func httpRemoteIP(r *http.Request) string {
	if xRealIP := r.Header.Get(httpXRealIPHeader); xRealIP != "" {
		return xRealIP
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// HTTPMiddleware is an HTTP middleware that rate limits by remote IP and the request URL.
// The remote IP is retrieved by the X-Real-IP header. Use this middleware after webmiddleware.ProxyHeaders()
func HTTPMiddleware(limiter Interface, class string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resource := httpRequestResource(r, class)
			limit, result := limiter.RateLimit(resource)
			result.SetHTTPHeaders(w.Header())
			if limit {
				webhandlers.Error(w, r, errRateLimitExceeded.WithAttributes("key", resource.Key(), "rate", result.Limit))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
