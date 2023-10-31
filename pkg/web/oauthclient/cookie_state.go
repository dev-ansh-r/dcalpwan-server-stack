package oauthclient

import (
	"encoding/gob"
	"net/http"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/random"
	"go.thethings.network/lorawan-stack/v3/pkg/web/cookie"
)

// state is the shape of the state for the OAuth flow.
type state struct {
	Secret string
	Next   string
}

func init() {
	gob.Register(state{})
}

// StateCookie returns the cookie storing the state of the console.
func (oc *OAuthClient) StateCookie() *cookie.Cookie {
	sameSite := http.SameSiteLaxMode
	if oc.config.CrossSiteCookie {
		sameSite = http.SameSiteNoneMode
	}
	return &cookie.Cookie{
		Name:     oc.config.StateCookieName,
		HTTPOnly: true,
		Path:     oc.getMountPath(),
		MaxAge:   10 * time.Minute,
		SameSite: sameSite,
	}
}

func newState(next string) state {
	return state{
		Secret: random.String(16),
		Next:   next,
	}
}

func (oc *OAuthClient) getStateCookie(w http.ResponseWriter, r *http.Request) (state, error) {
	s := state{}
	ok, err := oc.StateCookie().Get(w, r, &s)
	if err != nil {
		return s, err
	}

	if !ok {
		return s, err
	}

	return s, nil
}

func (oc *OAuthClient) setStateCookie(w http.ResponseWriter, r *http.Request, value state) error {
	return oc.StateCookie().Set(w, r, value)
}

func (oc *OAuthClient) removeStateCookie(w http.ResponseWriter, r *http.Request) {
	oc.StateCookie().Remove(w, r)
}
