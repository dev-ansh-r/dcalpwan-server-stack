// Package crypto implements LoRaWAN crypto.
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

var pingOffsetCipher cipher.Block

func init() {
	c, err := aes.NewCipher([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		panic(fmt.Sprintf("failed to create ping offset cipher: %s", err))
	}
	pingOffsetCipher = c
}

const (
	minPingPeriod = 1 << 5
	maxPingPeriod = 1 << 12
)

var errInvalidPingPeriod = errors.DefineInvalidArgument("ping_period", fmt.Sprintf("ping period must be a power of 2 between %d and %d, got '{value}'", minPingPeriod, maxPingPeriod))

func ComputePingOffset(beaconTime uint32, devAddr types.DevAddr, pingPeriod uint16) (uint16, error) {
	if pingPeriod < minPingPeriod || pingPeriod > maxPingPeriod {
		return 0, errInvalidPingPeriod.WithAttributes("value", pingPeriod)
	}
	var buf [16]byte
	binary.LittleEndian.PutUint32(buf[0:4], beaconTime)
	copy(buf[4:8], reverse(devAddr[:]))

	var rand [16]byte
	pingOffsetCipher.Encrypt(rand[:], buf[:])

	return (uint16(rand[0]) + uint16(rand[1])*256) % pingPeriod, nil
}
