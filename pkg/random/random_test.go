package random_test

import (
	"testing"
	"time"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/random"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestPseudoRandom(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	a.So(Bytes(10), assertions.ShouldHaveLength, 10)

	a.So(Int63n(10), assertions.ShouldBeGreaterThanOrEqualTo, 0)
	a.So(Int63n(10), assertions.ShouldBeLessThan, 10)

	a.So(String(10), assertions.ShouldHaveLength, 10)
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(100)
	}
}

func BenchmarkIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int63n(100)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(100)
	}
}

func TestJitter(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	d := time.Duration(424242)
	p := 0.1
	for i := 0; i < 100; i++ {
		// Jitter of 10%
		t := Jitter(d, p)
		df := float64(d)
		a.So(t, should.BeBetweenOrEqual, df-df*p, df+df*p)
	}
}

func TestCanJitter(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)
	p := 0.15
	for d := time.Duration(0); d < 10; d++ {
		// Values smaller or equal to 3 get clamped to zero, which cannot be
		// used as a upper bound for the random number generator.
		a.So(CanJitter(d, p), should.Equal, d > 3)
	}
}
