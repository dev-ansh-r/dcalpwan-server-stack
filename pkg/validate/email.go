package validate

import (
	"regexp"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	errEmail = errors.DefineInvalidArgument("email", "`{email}` is not a valid email.")
)

// Email checks whether the input value is a valid email or not.
func Email(str string) error {
	if !emailRegex.MatchString(str) {
		return errEmail.WithAttributes("email", str)
	}
	return nil
}
