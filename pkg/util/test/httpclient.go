package test

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"
	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
)

// HTTPClientProvider is a valid httpclient.Provider.
var HTTPClientProvider = httpclient.NewProvider(
	tlsconfig.ConfigurationProvider(func(context.Context) tlsconfig.Config {
		return tlsconfig.Config{}
	}),
)
