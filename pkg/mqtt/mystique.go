package mqtt

import (
	mqttlog "github.com/TheThingsIndustries/mystique/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

type logWrapper struct {
	log.Interface
}

func (m logWrapper) Debug(msg string) {
	m.Interface.Debug(msg)
}

func (m logWrapper) Info(msg string) {
	m.Interface.Info(msg)
}

func (m logWrapper) Warn(msg string) {
	m.Interface.Warn(msg)
}

func (m logWrapper) Error(msg string) {
	m.Interface.Error(msg)
}

func (m logWrapper) Fatal(msg string) {
	m.Interface.Fatal(msg)
}

func (m logWrapper) WithField(str string, v any) mqttlog.Interface {
	return logWrapper{
		Interface: m.Interface.WithField(str, v),
	}
}

func (m logWrapper) WithFields(f mqttlog.Fielder) mqttlog.Interface {
	return logWrapper{
		Interface: m.Interface.WithFields(f),
	}
}

func (m logWrapper) WithError(err error) mqttlog.Interface {
	return logWrapper{
		Interface: m.Interface.WithError(err),
	}
}

// Logger wraps a log.Interface to the Mystique logger interface.
func Logger(i log.Interface) mqttlog.Interface {
	return logWrapper{Interface: i}
}
