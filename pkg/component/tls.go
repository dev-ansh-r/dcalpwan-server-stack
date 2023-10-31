package component

import (
	"context"
	"crypto/tls"

	"go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"
)

func (c *Component) getTLSConfig(ctx context.Context) tlsconfig.Config {
	conf := c.GetBaseConfig(ctx).TLS
	// TODO: Remove detection mechanism (https://github.com/TheThingsNetwork/lorawan-stack/issues/1450)
	if conf.Source == "" {
		switch {
		case conf.ACME.Enable:
			conf.Source = "acme"
		case conf.Certificate != "" && conf.Key != "":
			conf.Source = "file"
		case conf.KeyVault.ID != "":
			conf.Source = "key-vault"
		}
	}
	if conf.Source == "key-vault" {
		conf.KeyVault.CertificateProvider = c.keyService
	}
	return conf
}

// GetTLSServerConfig gets the component's server TLS config and applies the given options.
func (c *Component) GetTLSServerConfig(ctx context.Context, opts ...tlsconfig.Option) (*tls.Config, error) {
	return tlsconfig.ConfigurationProvider(c.getTLSConfig).GetTLSServerConfig(ctx, opts...)
}

// GetTLSClientConfig gets the component's client TLS config and applies the given options.
func (c *Component) GetTLSClientConfig(ctx context.Context, opts ...tlsconfig.Option) (*tls.Config, error) {
	return tlsconfig.ConfigurationProvider(c.getTLSConfig).GetTLSClientConfig(ctx, opts...)
}
