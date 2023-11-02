package crypto

import (
	"context"
	"crypto/tls"
)

// KeyVault provides access to private keys.
type KeyVault interface {
	Key(ctx context.Context, label string) ([]byte, error)
	ServerCertificate(ctx context.Context, label string) (tls.Certificate, error)
	ClientCertificate(ctx context.Context, label string) (tls.Certificate, error)
}
