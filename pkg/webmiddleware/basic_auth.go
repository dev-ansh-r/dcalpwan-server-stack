package webmiddleware

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

// Authenticator is the interface the Basic Auth middleware uses to authenticate users.
type Authenticator interface {
	// Authenticate the user with the given password and return true if successful.
	// An error may be returned if an internal error occurred.
	Authenticate(username, password string) (ok bool, err error)
}

type mapAuthenticator struct {
	usernamesPasswords map[string]string
}

func (a *mapAuthenticator) Authenticate(username, password string) (bool, error) {
	validPassword, ok := a.usernamesPasswords[username]
	if !ok {
		return false, nil
	}
	return subtle.ConstantTimeCompare([]byte(password), []byte(validPassword)) == 1, nil
}

// AuthUsers authenticates users with the given map[username]password.
func AuthUsers(usernamesPasswords map[string]string) Authenticator {
	return &mapAuthenticator{usernamesPasswords}
}

// AuthUser is the same as AuthUsers, but for a single user.
func AuthUser(username, password string) Authenticator {
	return AuthUsers(map[string]string{username: password})
}

var noopAuthenticator = &mapAuthenticator{}

// BasicAuth returns a middleware that authenticates users with Basic Auth.
func BasicAuth(realm string, authenticator Authenticator) MiddlewareFunc {
	if authenticator == nil {
		authenticator = noopAuthenticator
	}
	writeUnauthorized := func(w http.ResponseWriter) {
		w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
		w.WriteHeader(http.StatusUnauthorized)
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok {
				writeUnauthorized(w)
				return
			}
			authenticated, err := authenticator.Authenticate(username, password)
			if err != nil {
				webhandlers.Error(w, r, err)
				return
			}
			if !authenticated {
				writeUnauthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
