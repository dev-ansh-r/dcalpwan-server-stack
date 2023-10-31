package account

import (
	"fmt"
	"net/http"
	"net/url"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

var errUnauthenticated = errors.DefineUnauthenticated("not_authenticated", "not authenticated")

func (s *server) requireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, _, err := s.session.Get(w, r)
		if err != nil {
			webhandlers.Error(w, r, errUnauthenticated.WithCause(err))
			return
		}
		next.ServeHTTP(w, r)
	})
}

const nextKey = "n"

func (s *server) redirectToNext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, _, err := s.session.Get(w, r)
		if err == nil {
			next := r.URL.Query().Get(nextKey)
			if next == "" {
				next = s.config.Mount
			}
			url, err := url.Parse(next)
			if err != nil {
				webhandlers.Error(w, r, err)
				return
			}
			http.Redirect(w, r, fmt.Sprintf("%s?%s", url.Path, url.RawQuery), http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
