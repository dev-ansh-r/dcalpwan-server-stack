package webmiddleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestFillContext(t *testing.T) {
	a := assertions.New(t)

	type ctxKeyType struct{}
	var ctxKey ctxKeyType

	m := FillContext(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxKey, "value")
	})

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.So(r.Context().Value(ctxKey), should.Equal, "value")
	})).ServeHTTP(rec, r)
}
