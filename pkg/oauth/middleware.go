package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const nextKey = "n"

func (s *server) redirectToLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, _, err := s.session.Get(w, r)
		if err != nil {
			values := make(url.Values)
			values.Set(nextKey, fmt.Sprintf("%s?%s", r.URL.Path, r.URL.Query().Encode()))
			http.Redirect(w, r, fmt.Sprintf("%s?%s", path.Join(s.config.Mount, "login"), values.Encode()), http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
