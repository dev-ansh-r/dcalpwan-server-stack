package oauthclient

import (
	"encoding/gob"
	"net/http"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/web/cookie"
)

func init() {
	gob.Register(authCookie{})
}

// AuthCookie returns a new authCookie.
func (oc *OAuthClient) AuthCookie() *cookie.Cookie {
	return &cookie.Cookie{
		Name:     oc.config.AuthCookieName,
		HTTPOnly: true,
		Path:     oc.getMountPath(),
	}
}

// authCookie is the shape of the authentication cookie.
type authCookie struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
}

var errNoAuthCookie = errors.DefineUnauthenticated("no_auth_cookie", "no auth cookie")

func (oc *OAuthClient) getAuthCookie(w http.ResponseWriter, r *http.Request) (authCookie, error) {
	value := authCookie{}
	ok, err := oc.AuthCookie().Get(w, r, &value)
	if err != nil {
		return authCookie{}, err
	}
	if !ok {
		return authCookie{}, errNoAuthCookie.New()
	}
	return value, nil
}

func (oc *OAuthClient) setAuthCookie(w http.ResponseWriter, r *http.Request, value authCookie) error {
	return oc.AuthCookie().Set(w, r, value)
}

func (oc *OAuthClient) removeAuthCookie(w http.ResponseWriter, r *http.Request) {
	oc.AuthCookie().Remove(w, r)
}
