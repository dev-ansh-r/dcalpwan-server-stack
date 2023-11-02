package cryptoutil

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/pem"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/crypto"
)

type memKeyVault struct {
	m map[string][]byte
}

// Key implements crypto.KeyVault.
func (kv *memKeyVault) Key(_ context.Context, label string) ([]byte, error) {
	key, ok := kv.m[label]
	if !ok {
		return nil, errKeyNotFound.WithAttributes("label", label)
	}
	return key, nil
}

func (kv *memKeyVault) certificate(label string) (tls.Certificate, error) {
	raw, ok := kv.m[label]
	if !ok {
		return tls.Certificate{}, errCertificateNotFound.WithAttributes("label", label)
	}
	certPEMBlock, keyPEMBlock := &bytes.Buffer{}, &bytes.Buffer{}
	for {
		block, rest := pem.Decode(raw)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			if err := pem.Encode(certPEMBlock, block); err != nil {
				return tls.Certificate{}, err
			}
		} else if block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, " PRIVATE KEY") {
			if err := pem.Encode(keyPEMBlock, block); err != nil {
				return tls.Certificate{}, err
			}
		}
		raw = rest
	}
	return tls.X509KeyPair(certPEMBlock.Bytes(), keyPEMBlock.Bytes())
}

// ServerCertificate implements crypto.KeyVault.
func (kv *memKeyVault) ServerCertificate(_ context.Context, label string) (tls.Certificate, error) {
	return kv.certificate(label)
}

// ClientCertificate implements crypto.KeyVault.
func (kv *memKeyVault) ClientCertificate(_ context.Context, label string) (tls.Certificate, error) {
	return kv.certificate(label)
}

// NewMemKeyVault returns a crypto.KeyVault that stores keys in memory.
// Certificates must be PEM encoded.
// The given map must not be modified after calling this function.
func NewMemKeyVault(m map[string][]byte) crypto.KeyVault {
	return &memKeyVault{
		m: m,
	}
}
