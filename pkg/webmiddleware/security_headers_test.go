package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestSecurityHeaders(t *testing.T) {
	a := assertions.New(t)
	m := SecurityHeaders()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})).ServeHTTP(rec, r)
	res := rec.Result()
	for k, v := range map[string]string{
		"Referrer-Policy":        "strict-origin-when-cross-origin",
		"X-Content-Type-Options": "nosniff",
		"X-Frame-Options":        "SAMEORIGIN",
		"X-XSS-Protection":       "1; mode=block",
	} {
		a.So(res.Header.Get(k), should.Equal, v)
	}
}
