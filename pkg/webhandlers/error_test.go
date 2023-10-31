package webhandlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

func TestErrorHandler(t *testing.T) {
	ctx, getError := NewContextWithErrorValue(test.Context())

	err := errors.New("some_error")

	a := assertions.New(t)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r = r.WithContext(ctx)
	rec := httptest.NewRecorder()

	Error(rec, r, err)

	res := rec.Result()
	a.So(res.StatusCode, should.Equal, http.StatusInternalServerError)

	body, _ := io.ReadAll(res.Body)
	a.So(string(body), should.ContainSubstring, "some_error")

	a.So(getError(), should.EqualErrorOrDefinition, err)
}
