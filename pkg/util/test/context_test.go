package test_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestContext(t *testing.T) {
	assertions.New(t).So(Context(), should.Equal, DefaultContext)
}
