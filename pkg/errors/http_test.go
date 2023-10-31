package errors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	_ "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestHTTP(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errDef := errors.DefineInvalidArgument("test_http_conversion_err_def", "HTTP Conversion Error", "foo")
	a.So(errors.FromGRPCStatus(errDef.GRPCStatus()).Definition, should.EqualErrorOrDefinition, errDef)

	errHello := errDef.WithAttributes("foo", "bar", "baz", "qux")
	errHelloExpected := errDef.WithAttributes("foo", "bar")

	handler := func(w http.ResponseWriter, _ *http.Request) {
		err := errors.ToHTTP(errHello, w)
		a.So(err, should.BeNil)
	}

	req := httptest.NewRequest("GET", "http://example.com/err", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	a.So(w.Result().StatusCode, should.Equal, http.StatusBadRequest)
	a.So(errors.FromHTTP(resp), should.EqualErrorOrDefinition, errHelloExpected)
}
