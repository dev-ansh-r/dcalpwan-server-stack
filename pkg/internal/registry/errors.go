// Package registry contains commonly used device registry functionality.
package registry

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

var errEndDeviceEUIsTaken = errors.DefineAlreadyExists(
	"end_device_euis_taken",
	"an end device with JoinEUI `{join_eui}` and DevEUI `{dev_eui}` is already registered as `{device_id}` in application `{application_id}`", // nolint:lll
)

// UniqueEUIViolationErr creates a unique EUI violation error with the given UID.
func UniqueEUIViolationErr(_ context.Context, joinEUI, devEUI types.EUI64, uid string) error {
	deviceIDs, err := unique.ToDeviceID(uid)
	if err != nil {
		return err
	}
	attributes := []any{
		"join_eui", joinEUI,
		"dev_eui", devEUI,
		"device_id", deviceIDs.GetDeviceId(),
		"application_id", deviceIDs.GetApplicationIds().GetApplicationId(),
	}
	return errEndDeviceEUIsTaken.WithAttributes(attributes...)
}
