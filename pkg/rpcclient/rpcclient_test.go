package rpcclient_test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/rpcclient"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func TestOptions(t *testing.T) {
	rpcclient.DefaultDialOptions(test.Context())
	// not really anything to test here
}
