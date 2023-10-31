package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
	"google.golang.org/grpc/metadata"
)

func TestMetadata(t *testing.T) {
	a := assertions.New(t)

	m := Metadata("authorization", "X-Request-Id")

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Authorization", "Bearer token")
	r.Header.Set("X-Request-Id", "XXX")

	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md, _ := metadata.FromIncomingContext(r.Context())
		a.So(md.Get("authorization"), should.Resemble, []string{"Bearer token"})
		a.So(md.Get("x-request-id"), should.Resemble, []string{"XXX"})
	})).ServeHTTP(rec, r)
}
