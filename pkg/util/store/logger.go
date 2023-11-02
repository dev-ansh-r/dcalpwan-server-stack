package store

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"google.golang.org/grpc/codes"
)

// loggerHook is a bun.QueryHook that logs the queries.
type loggerHook struct {
	logger log.Interface
}

// LoggerHookOption is an option for the LoggerHook.
type LoggerHookOption func(*loggerHook)

// NewLoggerHook returns a new bun.QueryHook that logs queries to the logger.
func NewLoggerHook(logger log.Interface, opts ...LoggerHookOption) bun.QueryHook {
	h := &loggerHook{
		logger: logger,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// BeforeQuery is the hook that is executed before the query runs.
func (*loggerHook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

// AfterQuery is the hook that is executed after the query runs.
func (h *loggerHook) AfterQuery(_ context.Context, event *bun.QueryEvent) {
	operation := event.Operation()
	switch operation {
	case "BEGIN", "COMMIT", "ROLLBACK":
		return
	}
	logFields := log.Fields(
		"operation", event.Operation(),
		"duration", time.Since(event.StartTime).Round(time.Microsecond),
		"query", event.Query,
	)
	if event.Result != nil {
		if rows, err := event.Result.RowsAffected(); err == nil {
			logFields = logFields.WithField("rows", rows)
		}
	}
	if event.Err != nil {
		err := WrapDriverError(event.Err)
		logFields = logFields.WithError(err)
		switch errors.Code(err) {
		case uint32(codes.Canceled),
			uint32(codes.DeadlineExceeded),
			uint32(codes.NotFound),
			uint32(codes.AlreadyExists):
		default:
			h.logger.WithFields(logFields).Debug("Database error")
			return
		}
	}
	h.logger.WithFields(logFields).Debug("Run database query")
}
