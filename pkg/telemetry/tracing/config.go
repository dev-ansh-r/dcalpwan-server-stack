package tracing

import "go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"

// Config represents configuration for OpenTelemetry tracing.
type Config struct {
	Enable            bool            `name:"enable" description:"Enable telemetry"`
	Exporter          string          `name:"exporter" description:"Telemetry exporter (otlp, writer)"`
	CollectorConfig   CollectorConfig `name:"collector-config" description:"Trace collector exporter configuration"`
	WriterConfig      WriterConfig    `name:"writer-config" description:"Writer exporter configuration"`
	SampleProbability float64         `name:"sample-probability" description:"Sampling probability. Fractions >= 1 will always sample. Fractions < 0 are treated as zero"` //nolint:lll
}

// CollectorConfig represents configuration for the trace collector exporter.
type CollectorConfig struct {
	EndpointURL string           `name:"endpoint-url" description:"The URL of the collector endpoint"`
	Insecure    bool             `name:"insecure" description:"Use insecure connection"`
	TLS         tlsconfig.Client `name:"tls"`
}

// WriterConfig represents configuration for the stdout exporter.
type WriterConfig struct {
	Destination string `name:"destination" description:"Destination of telemetry writer (stdout, stderr)"`
	Timestamps  bool   `name:"timestamps" description:"Print timestamps"`
	Pretty      bool   `name:"pretty" description:"Human readable format"`
}
