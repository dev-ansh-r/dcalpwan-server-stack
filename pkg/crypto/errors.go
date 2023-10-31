package crypto

import (
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

func errInvalidSize(typeName, typeDescription, expectedSize string) *errors.Definition {
	return errors.DefineInvalidArgument(
		fmt.Sprintf("%s_size", typeName),
		fmt.Sprintf("invalid %s size of {size} bytes, expected %s bytes", typeDescription, expectedSize),
	)
}
