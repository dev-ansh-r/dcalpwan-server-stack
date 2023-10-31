package webmiddleware

import (
	"context"
	"net/http"

	"github.com/gorilla/securecookie"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

type secureCookieCtxKeyType struct{}

var secureCookieCtxKey secureCookieCtxKeyType

// Cookies is a middleware that allows handling of secure cookies.
func Cookies(hashKey, blockKey []byte) MiddlewareFunc {
	s := securecookie.New(hashKey, blockKey)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, secureCookieCtxKey, s)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var errInvalidContext = errors.DefineInternal("secure_cookie_invalid_context", "No secure cookie value in this context")

// GetSecureCookie retrieves the secure cookie encoder from the context.
func GetSecureCookie(ctx context.Context) (*securecookie.SecureCookie, error) {
	secureCookie, _ := ctx.Value(secureCookieCtxKey).(*securecookie.SecureCookie)
	if secureCookie == nil {
		return nil, errInvalidContext.New()
	}
	return secureCookie, nil
}
