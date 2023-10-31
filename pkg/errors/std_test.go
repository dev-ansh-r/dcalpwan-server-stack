package errors_test

import (
	"io"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestStandardLibraryErrors(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	var (
		d1 = errors.Define("test_std_1", "definition 1")
		d2 = errors.Define("test_std_2", "definition 2")
		d3 = errors.Define("test_std_3", "definition 3")
	)

	a.So(errors.Is(d1, d1), should.BeTrue)
	a.So(errors.Is(d1, d2), should.BeFalse)

	e := d1.WithCause(d2.WithCause(io.EOF))

	a.So(errors.Is(e, d3), should.BeFalse)

	a.So(errors.Is(e, d1), should.BeTrue)
	a.So(errors.Is(e, d2), should.BeTrue)
	a.So(errors.Is(e, io.EOF), should.BeTrue)
}
