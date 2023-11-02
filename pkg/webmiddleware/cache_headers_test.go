package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestNoCache(t *testing.T) {
	a := assertions.New(t)

	r := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})).ServeHTTP(rec, r)
	res := rec.Result()

	a.So(res.Header.Get("Cache-Control"), should.Equal, "no-store")
	a.So(res.Header.Get("Pragma"), should.Equal, "no-cache")
}
