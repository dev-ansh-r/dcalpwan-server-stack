package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (s *server) redirectToLocal(w http.ResponseWriter, r *http.Request) {
	port, err := strconv.Atoi(r.URL.Query().Get("port"))
	if err != nil {
		port = 11885
	}

	params := make(url.Values)
	for k, v := range r.URL.Query() {
		params[k] = v
	}
	delete(params, "port")

	url := url.URL{
		Scheme:   "http",
		Host:     fmt.Sprintf("localhost:%d", port),
		Path:     "/oauth/callback",
		RawQuery: params.Encode(),
	}

	http.Redirect(w, r, url.String(), http.StatusFound)
}
