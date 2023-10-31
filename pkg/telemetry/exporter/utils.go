package telemetry

import (
	"context"
	"runtime"

	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/telemetry/exporter/models"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
)

// DispatchTask returns a task that calls Dispatch with the provided consumerID.
func DispatchTask(q TaskQueue, consumerID string) task.Func {
	return func(ctx context.Context) error {
		return q.Dispatch(ctx, consumerID)
	}
}

// PopTask returns a task that calls Pop with the provided consumerID.
func PopTask(q TaskQueue, consumerID string) task.Func {
	return func(ctx context.Context) error {
		return q.Pop(ctx, consumerID)
	}
}

// OSTelemetryData returns the OS telemetry data which is attached to telemetry messages.
func OSTelemetryData() *models.OSTelemetry {
	return &models.OSTelemetry{
		OperatingSystem: runtime.GOOS,
		Arch:            runtime.GOARCH,
		BinaryVersion:   version.String(),
		GolangVersion:   runtime.Version(),
	}
}
