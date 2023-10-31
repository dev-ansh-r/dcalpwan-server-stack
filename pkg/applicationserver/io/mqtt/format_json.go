package mqtt

import (
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/formatters"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/mqtt/topics"
)

type json struct {
	topics.Layout
	formatters.Formatter
}

// JSON is a format that uses the default topic layout and JSON formatter.
var JSON Format = &json{
	Layout:    topics.Default,
	Formatter: formatters.JSON,
}
