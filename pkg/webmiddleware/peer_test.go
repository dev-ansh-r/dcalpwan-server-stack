package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
	"google.golang.org/grpc/peer"
)

func TestPeer(t *testing.T) {
	a := assertions.New(t)

	m := Peer()

	r := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		peer, ok := peer.FromContext(r.Context())
		a.So(ok, should.BeTrue)
		a.So(peer.Addr.String(), should.Equal, "192.0.2.1:1234")
	})).ServeHTTP(rec, r)
}
