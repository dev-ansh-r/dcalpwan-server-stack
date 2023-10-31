package alcsyncv1

import "go.thethings.network/lorawan-stack/v3/pkg/errors"

var (
	errNoAssociation  = errors.DefineInternal("no_association", "no association available")
	errUnknownCommand = errors.DefineNotFound(
		"unknown_command", "unknown command", "command_id", "command_payload",
	)
	errUnsuportedCommand = errors.DefineUnimplemented(
		"unsupported_command", "unsupported command", "command_id", "command_payload",
	)
	// ErrIgnoreDownlink is a sentinel error returned when the command result should be ignored.
	errIgnoreDownlink = errors.DefineUnavailable("downlink_unavailable", "downlink unavailable")

	errCommandCreationFailed = errors.Define(
		"command_creation_failed", "create command", "command_id", "command_payload", "remaining_payload",
	)
	errDownlinkCreationFailed = errors.Define("downlink_creation_failed", "create downlink")

	errInvalidFieldType = errors.DefineCorruption("invalid_field_type", "field `{field}` has the wrong type `{type}`")
	errPkgDataMerge     = errors.DefineCorruption("pkg_data_merge", "merge package data")

	errInsufficientLength = errors.DefineInvalidArgument(
		"insufficient_length", "command payload has insufficient length", "expected_length", "actual_length",
	)
)
