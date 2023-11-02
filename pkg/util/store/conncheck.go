package store

import (
	"database/sql/driver"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

type errConnCheckFailed struct {
	inner error
}

func (e *errConnCheckFailed) Error() string {
	if e.inner != nil {
		return fmt.Sprintf("connection check failed: %v", e.inner)
	}
	return "connection check failed"
}

func (e *errConnCheckFailed) Unwrap() error {
	return e.inner
}

func (*errConnCheckFailed) Is(other error) bool {
	switch {
	case other == driver.ErrBadConn: //nolint:errorlint
		return true
	default:
		return false
	}
}

var errUnexpectedRead = &errConnCheckFailed{errors.New("unexpected read")}
