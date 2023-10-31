package component

import (
	"context"
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
)

// HTTPClient returns a new *http.Client with a default timeout and a configured transport.
func (c *Component) HTTPClient(ctx context.Context, opts ...httpclient.Option) (*http.Client, error) {
	return httpclient.NewProvider(c).HTTPClient(ctx, opts...)
}
