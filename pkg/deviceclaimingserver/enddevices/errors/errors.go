// Package errors defines common error types for all upstreams.
package errors

import (
	"fmt"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// DeviceErrors contains errors during claiming/unclaiming a batch of devices.
type DeviceErrors struct {
	Errors map[types.EUI64]errors.ErrorDetails
}

// Error implements error.
func (e DeviceErrors) Error() string {
	var errs strings.Builder
	for devEUI, err := range e.Errors {
		_, err := errs.WriteString(fmt.Sprintf("%s: %s, ", devEUI, err.Error()))
		if err != nil {
			return err.Error()
		}
	}
	return fmt.Sprintf("Errors per Device EUI: %s", strings.TrimRight(errs.String(), ", "))
}
