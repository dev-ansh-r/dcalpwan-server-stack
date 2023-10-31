package errors_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestStackTrace(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	err := errors.New("err")
	a.So(err.StackTrace(), should.HaveLength, 3)
}
