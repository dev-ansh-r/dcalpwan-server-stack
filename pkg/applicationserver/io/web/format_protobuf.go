package web

import "go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/formatters"

func init() {
	formats["protobuf"] = Format{
		Formatter:   formatters.Protobuf,
		Name:        "Protocol Buffers",
		ContentType: "application/octet-stream",
	}
}
