package webmiddleware

import "net/http"

// RequestURL populates (*http.Request).URL with the scheme and host.
func RequestURL() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Host = r.Host
			r.URL.Scheme = schemeHTTP
			if r.TLS != nil {
				r.URL.Scheme = schemeHTTPS
			}
			next.ServeHTTP(w, r)
		})
	}
}
