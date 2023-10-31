package log

// Option for the logger
type Option func(*Logger)

// WithLevel sets the level on the logger.
func WithLevel(level Level) Option {
	return func(logger *Logger) {
		if level != invalid {
			logger.Level = level
		}
	}
}
