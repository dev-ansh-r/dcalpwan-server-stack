package mqtt_test

import (
	"errors"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/log/handler/memory"
	"go.thethings.network/lorawan-stack/v3/pkg/mqtt"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestLogger(t *testing.T) {
	a := assertions.New(t)

	rec := memory.New()
	logger := &log.Logger{
		Level:   log.InfoLevel,
		Handler: rec,
	}
	mystiqueLogger := mqtt.Logger(logger)

	mystiqueLogger.Info("Yo!")
	a.So(rec.Entries, should.HaveLength, 1)

	mystiqueLogger.WithField("test", "hi").Warn("Oops")
	a.So(rec.Entries, should.HaveLength, 2)
	a.So(rec.Entries[1].Fields().Fields()["test"], should.Equal, "hi")

	mystiqueLogger.WithError(errors.New("fatal")).Warn("Failure")
	a.So(rec.Entries, should.HaveLength, 3)
	err, isError := rec.Entries[2].Fields().Fields()["error"].(error)
	a.So(isError, should.BeTrue)
	a.So(err.Error(), should.Equal, "fatal")
}
