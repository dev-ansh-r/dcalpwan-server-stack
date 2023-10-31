package errors_test

import (
	gerrors "errors"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestDetails(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errInvalidFoo := errors.DefineInvalidArgument("test_details_invalid_foo", "Invalid Foo")

	a.So(errInvalidFoo.Details(), should.BeEmpty)

	a.So(errors.Details(errInvalidFoo), should.BeEmpty)
	a.So(errors.Details(errInvalidFoo.GRPCStatus().Err()), should.BeEmpty)
	a.So(errors.Details(gerrors.New("go stdlib error")), should.BeEmpty)
}
