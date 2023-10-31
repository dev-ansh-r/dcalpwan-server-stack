package ttnpb_test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/apitest"
)

func TestFieldMaskSpecified(t *testing.T) {
	t.Parallel()
	apitest.RunTestFieldMaskSpecified(t, "ttn.lorawan.v3", ttnpb.RPCFieldMaskPaths)
}
