package tracing

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"os"
	"strings"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func exporterFromConfig(ctx context.Context, config *Config) (sdktrace.SpanExporter, error) {
	switch strings.ToLower(config.Exporter) {
	case "otlp":
		return initOTLPExporter(ctx, &config.CollectorConfig)

	case "writer":
		return initWriterExporter(&config.WriterConfig)

	default:
		return nil, errUnknownExporter.WithAttributes("exporter", config.Exporter)
	}
}

func initOTLPExporter(ctx context.Context, config *CollectorConfig) (sdktrace.SpanExporter, error) {
	var opts []grpc.DialOption
	if config.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsConfig := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
		if err := config.TLS.ApplyTo(tlsConfig); err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}

	conn, err := grpc.DialContext(ctx, config.EndpointURL, append(opts,
		grpc.FailOnNonTempDialError(true),
		grpc.WithBlock(),
	)...)
	if err != nil {
		return nil, err
	}
	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func initWriterExporter(config *WriterConfig) (sdktrace.SpanExporter, error) {
	var opts []stdouttrace.Option
	if config.Pretty {
		opts = append(opts, stdouttrace.WithPrettyPrint())
	}
	if !config.Timestamps {
		opts = append(opts, stdouttrace.WithoutTimestamps())
	}

	var w io.Writer
	switch strings.ToLower(config.Destination) {
	case "stdout":
		w = os.Stdout

	case "stderr":
		w = os.Stderr

	default:
		return nil, errors.New("unknown writer destination")
	}
	opts = append(opts, stdouttrace.WithWriter(w))

	return stdouttrace.New(opts...)
}
