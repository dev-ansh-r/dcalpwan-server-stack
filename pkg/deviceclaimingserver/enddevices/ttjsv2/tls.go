package ttjsv2

import (
	"crypto/tls"

	"go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"
	"go.thethings.network/lorawan-stack/v3/pkg/crypto"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
)

// TLSConfig contains the TLS configuration.
type TLSConfig struct {
	RootCA      string `yaml:"root-ca"`
	Source      string `yaml:"source"`
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"key"`
}

// IsZero returns true if the TLSConfig is empty.
func (conf TLSConfig) IsZero() bool {
	return conf == (TLSConfig{})
}

// TLSConfig returns the *tls.Config.
func (conf TLSConfig) TLSConfig(fetcher fetch.Interface, ks crypto.KeyService) (*tls.Config, error) {
	res := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	clientConfig := &tlsconfig.Client{
		FileReader: tlsconfig.FromFetcher(fetcher),
		RootCA:     conf.RootCA,
	}
	if err := clientConfig.ApplyTo(res); err != nil {
		return nil, err
	}
	clientAuthConfig := &tlsconfig.ClientAuth{
		Source:      conf.Source,
		FileReader:  tlsconfig.FromFetcher(fetcher),
		Certificate: conf.Certificate,
		Key:         conf.Key,
		KeyVault: tlsconfig.ClientKeyVault{
			CertificateProvider: ks,
		},
	}
	if err := clientAuthConfig.ApplyTo(res); err != nil {
		return nil, err
	}
	return res, nil
}
