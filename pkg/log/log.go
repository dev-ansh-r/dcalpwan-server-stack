// Package log contains all the structs to log TTN.
package log

// Interface is the interface for logging TTN.
type Interface interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Debugf(msg string, v ...any)
	Infof(msg string, v ...any)
	Warnf(msg string, v ...any)
	Errorf(msg string, v ...any)
	Fatalf(msg string, v ...any)
	WithField(string, any) Interface
	WithFields(Fielder) Interface
	WithError(error) Interface
}

// Stack is the interface of loggers that have a handler stack with middleware.
type Stack interface {
	// Log is an Interface.
	Interface

	// Use installs the specified middleware in the middleware stack.
	Use(Middleware)
}
