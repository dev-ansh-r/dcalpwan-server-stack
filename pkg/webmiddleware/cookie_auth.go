package webmiddleware

import (
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/auth"
)

// CookieAuth extracts the auth cookie and forwards it
// to the Authorization header.
func CookieAuth(cookieName string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") != "" {
				next.ServeHTTP(w, r)
				return
			}
			cookieValue, err := r.Cookie(cookieName)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			sc, err := GetSecureCookie(r.Context())
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			authCookie := &auth.CookieShape{}
			err = sc.Decode(cookieName, cookieValue.Value, authCookie)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			if authCookie.SessionSecret != "" {
				key := auth.JoinToken(auth.SessionToken, authCookie.SessionID, authCookie.SessionSecret)
				r.Header.Set("Authorization", "Bearer "+key)
			}
			next.ServeHTTP(w, r)
		})
	}
}
