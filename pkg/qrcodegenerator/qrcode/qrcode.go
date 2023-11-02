// Package qrcode provides generic interfaces to work with QR codes.
package qrcode

import (
	"encoding"
)

// Data represents QR code data.
type Data interface {
	Validate() error
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}
