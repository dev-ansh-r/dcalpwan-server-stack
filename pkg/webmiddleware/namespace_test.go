package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/log/handler/memory"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestNamespace(t *testing.T) {
	a := assertions.New(t)

	mem := memory.New()
	logger := log.NewLogger(mem)

	ctx := log.NewContext(test.Context(), logger)

	m := Namespace("test")
	req := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.FromContext(r.Context()).Info("foobar")
		entry := mem.Entries[0]
		a.So(entry.Message(), should.Equal, "foobar")
		fields := entry.Fields().Fields()
		a.So(fields, should.HaveLength, 1)
		a.So(fields["namespace"], should.Equal, "test")
	})).ServeHTTP(rec, req)
}
