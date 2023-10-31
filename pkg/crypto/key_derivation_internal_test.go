package crypto_test

import (
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/crypto"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestReverse(t *testing.T) {
	assertions.New(t).So(Reverse([]byte{1, 2, 3, 4}), should.Resemble, []byte{4, 3, 2, 1})
}
