package mac

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	ErrRequestNotFound = errors.DefineInvalidArgument(
		"request_not_found", "MAC response received, but corresponding request not found", "cid",
	)
	ErrNoPayload = errors.DefineInvalidArgument("no_payload", "no message payload specified")
)
