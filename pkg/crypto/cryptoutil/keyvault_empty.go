package cryptoutil

import (
	"context"
	"crypto/tls"

	"go.thethings.network/lorawan-stack/v3/pkg/crypto"
)

type emptyKeyVault struct{}

// Key implements crypto.KeyVault.
func (emptyKeyVault) Key(_ context.Context, label string) ([]byte, error) {
	return nil, errKeyNotFound.WithAttributes("label", label)
}

// ServerCertificate implements crypto.KeyVault.
func (emptyKeyVault) ServerCertificate(_ context.Context, label string) (tls.Certificate, error) {
	return tls.Certificate{}, errCertificateNotFound.WithAttributes("label", label)
}

// ClientCertificate implements crypto.KeyVault.
func (emptyKeyVault) ClientCertificate(_ context.Context, label string) (tls.Certificate, error) {
	return tls.Certificate{}, errCertificateNotFound.WithAttributes("label", label)
}

// EmptyKeyVault is an empty key vault.
var EmptyKeyVault crypto.KeyVault = emptyKeyVault{}
