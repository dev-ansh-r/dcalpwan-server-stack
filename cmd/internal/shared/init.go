package shared

import (
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

// InitializeFallbacks initializes configuration fallbacks.
func InitializeFallbacks(conf *config.ServiceBase) error {
	// Fallback to the default Redis configuration for the cache system
	if conf.Cache.Redis.IsZero() {
		conf.Cache.Redis = conf.Redis
	}
	// Fallback to the default Redis configuration for the events system
	if conf.Events.Redis.Config.IsZero() {
		conf.Events.Redis.Config = conf.Redis
	}
	if !conf.Redis.Equals(DefaultRedisConfig) {
		// Fallback to the default Redis configuration for the cache system
		if conf.Cache.Redis.Equals(DefaultRedisConfig) {
			conf.Cache.Redis = conf.Redis
		}
		// Fallback to the default Redis configuration for the events system
		if conf.Events.Redis.Config.Equals(DefaultRedisConfig) {
			conf.Events.Redis.Config = conf.Redis
		}
	}
	return nil
}

// InitializeLogger initializes the logger.
func InitializeLogger(conf *config.Log) (log.Stack, error) {
	var (
		logHandler log.Handler
		err        error
	)
	format := strings.ToLower(conf.Format)
	switch format {
	case "json", "console":
		logHandler, err = log.NewZap(format)
		if err != nil {
			return nil, ErrInitializeLogger.WithCause(err)
		}
	default:
		return nil, ErrInvalidLogFormat.WithAttributes("format", format)
	}
	return log.NewLogger(
		logHandler,
		log.WithLevel(conf.Level),
	), nil
}
