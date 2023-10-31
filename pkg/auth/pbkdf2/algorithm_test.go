package pbkdf2_test

import (
	"testing"

	. "github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/auth/pbkdf2"
)

func TestStringAlgorithm(t *testing.T) {
	a := New(t)

	a.So(Sha256.String(), ShouldEqual, "sha256")
	a.So(Sha512.String(), ShouldEqual, "sha512")
}

func TestParseAlgorithm(t *testing.T) {
	a := New(t)

	{
		alg, err := ParseAlgorithm("sha256")
		a.So(err, ShouldBeNil)
		a.So(alg, ShouldResemble, Sha256)
	}
}

func TestParseBadAlgorithm(t *testing.T) {
	a := New(t)

	_, err := ParseAlgorithm("bad")
	a.So(err, ShouldNotBeNil)
}

func TestNil(t *testing.T) {
	a := New(t)

	{
		var alg Algorithm
		h := alg.Hash()
		a.So(h, ShouldBeNil)
	}

	{
		var alg *Algorithm
		h := alg.Hash()
		a.So(h, ShouldBeNil)
	}
}
