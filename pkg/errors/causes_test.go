package errors_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestCause(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	cause := errors.New("cause")

	errInvalidFoo := errors.DefineInvalidArgument("test_cause_invalid_foo", "Invalid Foo: {foo}", "foo")

	a.So(errInvalidFoo.Cause(), should.EqualErrorOrDefinition, nil)
	a.So(errors.RootCause(errInvalidFoo), should.EqualErrorOrDefinition, errInvalidFoo)

	err1 := errInvalidFoo.WithCause(cause)

	a.So(func() { err1.WithCause(cause) }, should.Panic) //nolint:errcheck

	a.So(err1, should.HaveSameErrorDefinitionAs, errInvalidFoo)
	a.So(err1.Cause(), should.EqualErrorOrDefinition, cause)
	a.So(errors.RootCause(err1), should.EqualErrorOrDefinition, cause)
	a.So(errors.Stack(err1), should.Resemble, []error{err1, cause})

	errInvalidBar := errors.DefineInvalidArgument("test_cause_invalid_bar", "Invalid Bar")
	err2 := errInvalidBar.WithCause(err1)

	a.So(err2, should.HaveSameErrorDefinitionAs, errInvalidBar)
	a.So(err2.Cause(), should.EqualErrorOrDefinition, err1)
	a.So(errors.RootCause(err2), should.EqualErrorOrDefinition, cause)
	a.So(errors.Stack(err2), should.Resemble, []error{err2, err1, cause})
}
