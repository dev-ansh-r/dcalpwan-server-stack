package test_test

import (
	"testing"
	"time"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestMockTime(t *testing.T) {
	a := assertions.New(t)

	now := time.Unix(0, 42)
	clock := NewMockClock(now)
	a.So(clock.Now(), should.Resemble, now)

	d := 6*time.Hour + 5*time.Minute + 4*time.Second + 3*time.Millisecond + 2*time.Microsecond + 1*time.Nanosecond
	now = now.Add(d)
	a.So(clock.Add(d), should.Resemble, now)
	a.So(clock.Now(), should.Resemble, now)

	oldNow := now
	now = now.Add(d)
	a.So(clock.Set(now), should.Resemble, oldNow)
	a.So(clock.Now(), should.Resemble, now)

	n := 5
	afterCh := clock.After(time.Duration(n) * time.Nanosecond)
	for i := 0; i < n; i++ {
		select {
		case <-afterCh:
			t.Error("After channel read succeeded too soon")
		default:
		}
		now = now.Add(time.Nanosecond)
		a.So(clock.Add(time.Nanosecond), should.Resemble, now)
	}
	// Let the goroutine in After send the time on the channel.
	time.Sleep(Delay)
	select {
	case afterNow := <-afterCh:
		a.So(afterNow, should.Resemble, now)
	default:
		t.Error("After channel read should have succeeded")
	}

	select {
	case afterNow := <-afterCh:
		a.So(afterNow, should.BeZeroValue)
	default:
		t.Error("After channel should have been closed")
	}
}
