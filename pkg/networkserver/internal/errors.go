package internal

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	ErrCorruptedMACState = errors.DefineCorruption("corrupted_mac_state", "MAC state is corrupted")
	ErrInvalidDataRate   = errors.DefineInvalidArgument("data_rate", "invalid data rate")
	ErrInvalidPayload    = errors.DefineInvalidArgument("payload", "invalid payload")
	ErrUnknownChannel    = errors.Define("unknown_channel", "channel is unknown")

	ErrNetworkDownlinkSlot  = errors.DefineCorruption("network_downlink_slot", "generate network initiated downlink slot")
	ErrUplinkChannel        = errors.DefineCorruption("uplink_channel", "channel does not allow downlinks")
	ErrDownlinkChannel      = errors.DefineCorruption("downlink_channel", "channel does not allow uplinks")
	ErrSession              = errors.DefineCorruption("session", "no device session")
	ErrMACHandler           = errors.DefineCorruption("mac_handler", "missing MAC handler")
	ErrChannelDataRateRange = errors.DefineCorruption("channel_data_rate_range", "generate channel datarate range")
	ErrChannelMask          = errors.DefineCorruption("channel_mask", "generate channel mask")
)
