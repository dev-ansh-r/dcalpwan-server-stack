package test_test

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var _ ttnpb.ApplicationAccessServer = &MockApplicationAccessServer{}
