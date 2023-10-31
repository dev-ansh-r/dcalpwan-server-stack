package pubsub

import "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/formatters"

func init() {
	formats["json"] = Format{
		Formatter: formatters.JSON,
		Name:      "JSON",
	}
}
