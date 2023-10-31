package test_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestMust(t *testing.T) {
	for i, tc := range []struct {
		Value       any
		Error       error
		ShouldPanic bool
	}{
		{
			42,
			nil,
			false,
		},
		{
			errors.New("42"),
			nil,
			false,
		},
		{
			(error)(nil),
			errors.New("test"),
			true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assertions.New(t)

			fn := func() { Must(tc.Value, tc.Error) }

			if tc.ShouldPanic {
				a.So(fn, should.Panic)
			} else if a.So(fn, should.NotPanic) {
				v := Must(tc.Value, tc.Error)
				a.So(v, should.Resemble, tc.Value)
			}
		})
	}
}

func TestWaitTimeout(t *testing.T) {
	for _, tc := range []struct {
		Timeout time.Duration
		OK      bool
	}{
		{
			Timeout: Delay,
			OK:      false,
		},
		{
			Timeout: 10 * Delay,
			OK:      true,
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.Timeout), func(t *testing.T) {
			a := assertions.New(t)

			ok := WaitTimeout(tc.Timeout, func() { time.Sleep(2 * Delay) })
			a.So(ok, should.Equal, tc.OK)
		})
	}
}
