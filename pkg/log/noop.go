package log

// Noop just does nothing.
var Noop = &noop{}

// noop is a log.Interface that does nothing.
type noop struct{}

// Debug implements log.Interface.
func (noop) Debug(...any) {}

// Info implements log.Interface.
func (noop) Info(...any) {}

// Warn implements log.Interface.
func (noop) Warn(...any) {}

// Error implements log.Interface.
func (noop) Error(...any) {}

// Fatal implements log.Interface.
func (noop) Fatal(...any) {}

// Debugf implements log.Interface.
func (noop) Debugf(string, ...any) {}

// Infof implements log.Interface.
func (noop) Infof(string, ...any) {}

// Warnf implements log.Interface.
func (noop) Warnf(string, ...any) {}

// Errorf implements log.Interface.
func (noop) Errorf(string, ...any) {}

// Fatalf implements log.Interface.
func (noop) Fatalf(string, ...any) {}

// WithField implements log.Interface.
func (n noop) WithField(string, any) Interface { return n }

// WithFields implements log.Interface.
func (n noop) WithFields(Fielder) Interface { return n }

// WithError implements log.Interface.
func (n noop) WithError(error) Interface { return n }

// Use implements log.Stack
func (noop) Use(Middleware) {}
