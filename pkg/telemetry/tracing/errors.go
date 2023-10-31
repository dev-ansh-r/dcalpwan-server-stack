package tracing

import "go.thethings.network/lorawan-stack/v3/pkg/errors"

var errUnknownExporter = errors.DefineInvalidArgument("unknown_exporter", "unknown exporter {exporter}")
