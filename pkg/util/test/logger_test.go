package test_test

import (
	"errors"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/log"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func TestGetLogger(t *testing.T) {
	var logger log.Interface = GetLogger(t)

	logger = logger.WithField("foo", "bar")

	logger = logger.WithError(errors.New("example error"))

	logger = logger.WithFields(log.Fields("k1", "v1", "k2", "v2"))

	logger.Debug("This is a debug log")
	logger.Info("This is a info log")
	logger.Warn("This is a warn log")
	logger.Error("This is an error log")

	logger.Debugf("This is a %s log", "debug")
	logger.Infof("This is a %s log", "info")
	logger.Warnf("This is a %s log", "warn")
	logger.Errorf("This is an %s log", "error")
}
