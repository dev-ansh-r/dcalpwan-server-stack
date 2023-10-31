package webmiddleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestRecover(t *testing.T) {
	m := Recover()

	a := assertions.New(t)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("this panic is expected")
	})).ServeHTTP(rec, r)

	res := rec.Result()

	a.So(res.StatusCode, should.Equal, http.StatusInternalServerError)

	body, _ := io.ReadAll(res.Body)
	a.So(string(body), should.ContainSubstring, "http_recovered")
}
