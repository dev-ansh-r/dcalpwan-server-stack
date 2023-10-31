package assertions_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestShouldHaveSameErrorDefinition(t *testing.T) {
	a := assertions.New(t)

	errDef := errors.Define("test_error_assertions", "Error Assertions Test")
	errOtherDef := errors.Define("test_error_assertions_other", "Other Error Assertions Test")

	// Happy flow.
	a.So(ShouldHaveSameErrorDefinitionAs(errDef.WithAttributes("k", "v"), errDef.WithAttributes("foo", "bar")), should.BeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef.WithAttributes("k", "v"), errDef.WithAttributes("k", "v")), should.BeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef, errDef), should.BeEmpty)

	// Not same.
	a.So(ShouldHaveSameErrorDefinitionAs(errDef.WithAttributes("k", "v"), errOtherDef.WithAttributes("k", "v")), should.NotBeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef.WithAttributes("k", "v"), errOtherDef.WithAttributes("k", "v")), should.NotBeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef, errDef.WithAttributes("k", "v")), should.NotBeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef.WithAttributes("k", "v"), errDef), should.NotBeEmpty)
	a.So(ShouldEqualErrorOrDefinition(errDef, errOtherDef), should.NotBeEmpty)
}
