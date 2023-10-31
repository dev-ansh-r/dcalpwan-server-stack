package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestChain(t *testing.T) {
	a := assertions.New(t)

	var trace []string
	middleware := []MiddlewareFunc{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				trace = append(trace, "outer begin")
				next.ServeHTTP(w, r)
				trace = append(trace, "outer end")
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				trace = append(trace, "inner begin")
				next.ServeHTTP(w, r)
				trace = append(trace, "inner end")
			})
		},
	}

	chain := Chain(middleware, http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		trace = append(trace, "handler")
	}))
	chain.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	a.So(trace, should.Resemble, []string{"outer begin", "inner begin", "handler", "inner end", "outer end"})
}

func TestConditional(t *testing.T) {
	a := assertions.New(t)

	flag := false
	middleware := Conditional(
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				flag = true
				next.ServeHTTP(w, r)
			})
		},
		func(r *http.Request) bool {
			return r.Header.Get("X-Condition") == "foo"
		},
	)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Condition", "bar")
	rec := httptest.NewRecorder()
	middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.So(flag, should.Equal, false)
	})).ServeHTTP(rec, r)

	r = httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Condition", "foo")
	rec = httptest.NewRecorder()
	middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.So(flag, should.Equal, true)
	})).ServeHTTP(rec, r)
}
