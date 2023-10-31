package oauthclient

import (
	"net/http"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

// HandleLogin is the handler for redirecting the user to the authorization
// endpoint.
func (oc *OAuthClient) HandleLogin(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get(oc.nextKey)
	// Only allow relative paths.
	if !strings.HasPrefix(next, "/") && !strings.HasPrefix(next, "#") && !strings.HasPrefix(next, "?") {
		next = ""
	}

	// Set state cookie.
	state := newState(next)
	if err := oc.setStateCookie(w, r, state); err != nil {
		webhandlers.Error(w, r, err)
		return
	}

	conf, err := oc.oauthConfig(r.Context())
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	opts, err := oc.authCodeURLOpts(r.Context())
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	http.Redirect(w, r, conf.AuthCodeURL(state.Secret, opts...), http.StatusFound)
}
