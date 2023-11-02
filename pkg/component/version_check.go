package component

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
)

const (
	versionCheckPeriod  = 12 * time.Hour
	versionCheckTimeout = 10 * time.Second
)

// versionCheckTask returns the task configuration for a periodic version check.
func versionCheckTask(ctx context.Context, clientProvider httpclient.Provider) *task.Config {
	taskConfig := task.Config{
		Context: ctx,
		ID:      "version_check",
		Func: func(ctx context.Context) error {
			checkCtx, cancel := context.WithTimeout(ctx, versionCheckTimeout)
			defer cancel()

			client, err := clientProvider.HTTPClient(checkCtx)
			if err != nil {
				return err
			}
			update, err := version.CheckUpdate(checkCtx, version.WithClient(client))
			if err != nil {
				return err
			}
			version.LogUpdate(checkCtx, update)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(versionCheckPeriod):
				return nil
			}
		},
		Restart: task.RestartAlways,
		Backoff: task.DialBackoffConfig,
	}
	return &taskConfig
}
