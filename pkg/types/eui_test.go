package types_test

import (
	"encoding/json"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	. "go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

var (
	_ config.Configurable = &EUI64Prefix{}
	_ config.Stringer     = EUI64Prefix{}
)

func TestEUI64(t *testing.T) {
	a := assertions.New(t)

	eui := EUI64{0x26, 0x12, 0x34, 0x56, 0x42, 0x42, 0x42, 0x42}

	prefix := EUI64Prefix{EUI64{0x26}, 7}
	a.So(prefix.Matches(eui), should.BeTrue)

	// Prefix list
	{
		addr := EUI64{1, 2, 3, 4}
		a.So(addr.HasPrefix(EUI64Prefix{EUI64{0, 0, 0, 0}, 0}), should.BeTrue)
		a.So(addr.HasPrefix(EUI64Prefix{EUI64{1, 2, 3, 0}, 24}), should.BeTrue)
		a.So(addr.HasPrefix(EUI64Prefix{EUI64{2, 2, 3, 4}, 31}), should.BeFalse)
		a.So(addr.HasPrefix(EUI64Prefix{EUI64{1, 1, 3, 4}, 31}), should.BeFalse)
		a.So(addr.HasPrefix(EUI64Prefix{EUI64{1, 1, 1, 1}, 15}), should.BeFalse)
	}

	t.Run("JSON", func(t *testing.T) {
		a := assertions.New(t)

		const encodedEUI = `"2612345642424242"`

		b, err := json.Marshal(eui)
		if a.So(err, should.BeNil) {
			a.So(string(b), should.Equal, encodedEUI)
		}

		var decodedEUI EUI64
		err = json.Unmarshal([]byte(encodedEUI), &decodedEUI)
		if a.So(err, should.BeNil) {
			a.So(decodedEUI, should.Equal, eui)
		}

		const encodedPrefix = `"2600000000000000/7"`

		b, err = json.Marshal(prefix)
		if a.So(err, should.BeNil) {
			a.So(string(b), should.Equal, encodedPrefix)
		}

		var decodedPrefix EUI64Prefix
		err = json.Unmarshal([]byte(encodedPrefix), &decodedPrefix)
		if a.So(err, should.BeNil) {
			a.So(decodedPrefix, should.Equal, prefix)
		}
	})

	t.Run("Text", func(t *testing.T) {
		a := assertions.New(t)

		const encodedEUI = `2612345642424242`

		b, err := eui.MarshalText()
		if a.So(err, should.BeNil) {
			a.So(string(b), should.Equal, encodedEUI)
		}

		var decodedEUI EUI64
		err = decodedEUI.UnmarshalText([]byte(encodedEUI))
		if a.So(err, should.BeNil) {
			a.So(decodedEUI, should.Equal, eui)
		}

		const encodedPrefix = `2600000000000000/7`

		b, err = prefix.MarshalText()
		if a.So(err, should.BeNil) {
			a.So(string(b), should.Equal, encodedPrefix)
		}

		var decodedPrefix EUI64Prefix
		err = decodedPrefix.UnmarshalText([]byte(encodedPrefix))
		if a.So(err, should.BeNil) {
			a.So(decodedPrefix, should.Equal, prefix)
		}
	})

	t.Run("Number", func(t *testing.T) {
		a := assertions.New(t)

		const encodedEUI uint64 = 2743312668105523778

		n := eui.MarshalNumber()
		a.So(n, should.Resemble, encodedEUI)

		var decodedEUI EUI64
		decodedEUI.UnmarshalNumber(encodedEUI)
		a.So(decodedEUI, should.Equal, eui)
	})
}
