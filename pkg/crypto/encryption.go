package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

var errMalformedCipherText = errors.DefineInvalidArgument("malformed_cipher_text", "malformed cipher text")

// Encrypt encrypts a plain text message.
// Uses AES128 keys in GCM (Galois/Counter Mode).
// Since GCM uses a nonce, the encrypted message will be different each time the operation is run for the same set of inputs.
// The returned cipher is in the format |nonce(12)|tag(16)|encrypted(plaintextLen)|.
func Encrypt(key types.AES128Key, plaintext []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts an encrypted message.
// Uses AES128 keys in GCM (Galois/Counter Mode).
func Decrypt(key types.AES128Key, encrypted []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}
	if len(encrypted) < gcm.NonceSize() {
		return nil, errMalformedCipherText.New()
	}
	return gcm.Open(nil,
		encrypted[:gcm.NonceSize()],
		encrypted[gcm.NonceSize():],
		nil,
	)
}
