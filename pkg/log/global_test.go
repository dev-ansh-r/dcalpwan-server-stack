package log

import (
	"context"
	"errors"
	"testing"
)

func TestLogInterface(t *testing.T) {
	ctx := NewContext(context.Background(), Noop)

	Debug(ctx, "test debug msg")
	Info(ctx, "test info msg")
	Warn(ctx, "test warn msg")
	Error(ctx, "test error msg")
	Debugf(ctx, "test debugf msg %d", 42)
	Infof(ctx, "test infof msg %d", 42)
	Warnf(ctx, "test warnf msg %d", 42)
	Errorf(ctx, "test errorf msg %d", 42)

	for _, li := range []Interface{
		&Logger{},
		&noop{},
		WithField(ctx, "key", "value"),
		WithFields(ctx, Fields("key", "value", "number", 42)),
		WithError(ctx, errors.New("unknown error")),
	} {
		for _, dli := range []Interface{
			li,
			li.WithField("key", "value"),
			li.WithFields(Fields("key", "value", "number", 42)),
			li.WithError(errors.New("unknown error")),
		} {
			dli.Debug("test debug msg")
			dli.Info("test info msg")
			dli.Warn("test warn msg")
			dli.Error("test error msg")
			dli.Debugf("test debugf msg %d", 42)
			dli.Infof("test infof msg %d", 42)
			dli.Warnf("test warnf msg %d", 42)
			dli.Errorf("test errorf msg %d", 42)
		}
	}
}
