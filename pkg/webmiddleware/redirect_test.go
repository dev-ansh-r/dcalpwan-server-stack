package webmiddleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestRedirect(t *testing.T) {
	m := Redirect(RedirectConfiguration{
		Scheme: func(s string) string { return SchemeHTTPS },
		HostName: func(h string) string {
			if strings.HasPrefix(h, "dev.") {
				return h
			}
			return "dev." + h
		},
		Port: func(p uint) uint {
			switch p {
			case 1885:
				return 8885
			default:
				return 0
			}
		},
		Path: strings.ToLower,
	})

	t.Run("None", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "https://dev.example.com/path?query=true", nil)
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusOK)
		body, _ := io.ReadAll(res.Body)
		a.So(string(body), should.Equal, "OK")
	})

	for _, tc := range []struct {
		Name     string
		URL      string
		Redirect string
	}{
		{
			Name:     "HostName",
			URL:      "https://example.com/",
			Redirect: "https://dev.example.com/",
		},
		{
			Name:     "Scheme",
			URL:      "http://dev.example.com/",
			Redirect: "https://dev.example.com/",
		},
		{
			Name:     "Port",
			URL:      "http://dev.example.com:1885/",
			Redirect: "https://dev.example.com:8885/",
		},
		{
			Name:     "Path",
			URL:      "https://dev.example.com/PATH",
			Redirect: "https://dev.example.com/path",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)
			r := httptest.NewRequest(http.MethodGet, tc.URL, nil)
			rec := httptest.NewRecorder()
			m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Error("Handler was called when it shouldn't have")
			})).ServeHTTP(rec, r)

			res := rec.Result()

			a.So(res.StatusCode, should.Equal, http.StatusFound)
			a.So(res.Header.Get("Location"), should.Equal, tc.Redirect)
		})
	}
}
