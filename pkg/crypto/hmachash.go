package crypto

import (
	"crypto/hmac"
	"crypto/sha256"

	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// HMACHash calculates the  Keyed-Hash Message Authentication Code (HMAC, RFC 2104) hash of the data.
func HMACHash(key types.AES128Key, payload []byte) ([]byte, error) {
	h := hmac.New(sha256.New, key[:])
	_, err := h.Write(payload)
	if err != nil {
		return nil, err
	}
	return h.Sum([]byte{}), nil
}
