package log

import "context"

// Debug calls FromContext(ctx).Debug
func Debug(ctx context.Context, msg string) {
	FromContext(ctx).Debug(msg)
}

// Info calls FromContext(ctx).Info
func Info(ctx context.Context, msg string) {
	FromContext(ctx).Info(msg)
}

// Warn calls FromContext(ctx).Warn
func Warn(ctx context.Context, msg string) {
	FromContext(ctx).Warn(msg)
}

// Error calls FromContext(ctx).Error
func Error(ctx context.Context, msg string) {
	FromContext(ctx).Error(msg)
}

// Fatal calls FromContext(ctx).Fatal
func Fatal(ctx context.Context, msg string) {
	FromContext(ctx).Fatal(msg)
}

// Debugf calls FromContext(ctx).Debugf
func Debugf(ctx context.Context, msg string, v ...any) {
	FromContext(ctx).Debugf(msg, v...)
}

// Infof calls FromContext(ctx).Infof
func Infof(ctx context.Context, msg string, v ...any) {
	FromContext(ctx).Infof(msg, v...)
}

// Warnf calls FromContext(ctx).Warnf
func Warnf(ctx context.Context, msg string, v ...any) {
	FromContext(ctx).Warnf(msg, v...)
}

// Errorf calls FromContext(ctx).Errorf
func Errorf(ctx context.Context, msg string, v ...any) {
	FromContext(ctx).Errorf(msg, v...)
}

// Fatalf calls FromContext(ctx).Fatalf
func Fatalf(ctx context.Context, msg string, v ...any) {
	FromContext(ctx).Fatalf(msg, v...)
}

// WithField calls FromContext(ctx).WithField
func WithField(ctx context.Context, k string, v any) Interface {
	return FromContext(ctx).WithField(k, v)
}

// WithFields calls FromContext(ctx).WithFields
func WithFields(ctx context.Context, f Fielder) Interface {
	return FromContext(ctx).WithFields(f)
}

// WithError calls FromContext(ctx).WithError
func WithError(ctx context.Context, err error) Interface {
	return FromContext(ctx).WithError(err)
}
