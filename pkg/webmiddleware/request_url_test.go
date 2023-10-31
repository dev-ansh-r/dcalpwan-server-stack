package webmiddleware_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestRequestURL(t *testing.T) {
	m := RequestURL()

	t.Run("HTTP", func(t *testing.T) {
		a := assertions.New(t)

		r := httptest.NewRequest(http.MethodGet, "/path", nil)

		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.So(r.URL.String(), should.Equal, "http://example.com/path")
		})).ServeHTTP(rec, r)
	})

	t.Run("HTTPS", func(t *testing.T) {
		a := assertions.New(t)

		r := httptest.NewRequest(http.MethodGet, "/path", nil)
		r.TLS = &tls.ConnectionState{
			Version:           tls.VersionTLS12,
			HandshakeComplete: true,
			ServerName:        r.Host,
		}

		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.So(r.URL.String(), should.Equal, "https://example.com/path")
		})).ServeHTTP(rec, r)
	})
}
