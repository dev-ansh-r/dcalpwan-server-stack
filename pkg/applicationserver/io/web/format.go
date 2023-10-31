package web

import (
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/formatters"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

// Format is a format to use for web-based frontends.
type Format struct {
	formatters.Formatter
	Name        string
	ContentType string
}

var (
	formats = map[string]Format{}

	errFormatNotFound = errors.DefineNotFound("format_not_found", "format `{format}` not found")
)
