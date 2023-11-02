package band

import "go.thethings.network/lorawan-stack/v3/pkg/errors"

var (
	errBandNotFound          = errors.DefineNotFound("band_not_found", "band `{id}@{version}` not found")
	errDataRateIndexTooHigh  = errors.DefineInvalidArgument("data_rate_index_too_high", "data rate index must be lower or equal to {max}")
	errDataRateOffsetTooHigh = errors.DefineInvalidArgument("data_rate_offset_too_high", "data rate offset must be lower or equal to {max}")
	errInvalidChannelCount   = errors.DefineInvalidArgument("invalid_channel_count", "invalid number of channels defined")
	errUnsupportedChMaskCntl = errors.DefineInvalidArgument("chmaskcntl_unsupported", "ChMaskCntl `{chmaskcntl}` unsupported")
)
