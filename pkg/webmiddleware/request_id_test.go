package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestRequestID(t *testing.T) {
	m := RequestID()

	t.Run("New", func(t *testing.T) {
		a := assertions.New(t)

		r := httptest.NewRequest(http.MethodGet, "/", nil)

		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.So(r.Header.Get("X-Request-Id"), should.NotBeEmpty)
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.Header.Get("X-Request-Id"), should.NotBeEmpty)
	})

	t.Run("Existing", func(t *testing.T) {
		a := assertions.New(t)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.Header.Set("X-Request-Id", "XXX")

		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a.So(r.Header.Get("X-Request-Id"), should.Equal, "XXX")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.Header.Get("X-Request-Id"), should.Equal, "XXX")
	})
}
