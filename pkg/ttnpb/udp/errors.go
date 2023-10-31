package udp

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	errNoEUI        = errors.DefineInvalidArgument("no_eui", "packet is not long enough to contain the EUI")
	errPayload      = errors.DefineInvalidArgument("payload", "parse binary payload")
	errEUI          = errors.DefineInvalidArgument("eui", "parse EUI")
	errTimestamp    = errors.DefineInvalidArgument("timestamp", "parse timestamp")
	errDataRate     = errors.DefineInvalidArgument("data_rate", "invalid data rate")
	errNotScheduled = errors.DefineInvalidArgument("not_scheduled", "downlink message not scheduled")
)
