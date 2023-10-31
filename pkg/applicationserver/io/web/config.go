package web

import (
	"context"
	"net/url"
	"os"

	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// TemplatesConfig defines the configuration for the webhook templates registry.
type TemplatesConfig struct {
	Static      map[string][]byte `name:"-"`
	Directory   string            `name:"directory" description:"Retrieve the webhook templates from the filesystem"`
	URL         string            `name:"url" description:"Retrieve the webhook templates from a web server"`
	LogoBaseURL string            `name:"logo-base-url" description:"The base URL for the logo storage"`
}

// TemplateStore contains the webhook templates.
type TemplateStore interface {
	// GetTemplate returns the template with the given identifiers.
	GetTemplate(ctx context.Context, req *ttnpb.GetApplicationWebhookTemplateRequest) (*ttnpb.ApplicationWebhookTemplate, error)
	// ListTemplates returns the available templates.
	ListTemplates(ctx context.Context, req *ttnpb.ListApplicationWebhookTemplatesRequest) (*ttnpb.ApplicationWebhookTemplates, error)
}

// NewTemplateStore returns a TemplateStore based on the configuration.
func (c TemplatesConfig) NewTemplateStore(ctx context.Context, httpClientProvider httpclient.Provider) (TemplateStore, error) {
	var fetcher fetch.Interface
	switch {
	case c.Static != nil:
		fetcher = fetch.NewMemFetcher(c.Static)
	case c.Directory != "":
		if stat, err := os.Stat(c.Directory); err == nil && stat.IsDir() {
			fetcher = fetch.FromFilesystem(c.Directory)
			break
		}
		fallthrough
	case c.URL != "":
		var err error
		httpClient, err := httpClientProvider.HTTPClient(ctx, httpclient.WithCache(true))
		if err != nil {
			return nil, err
		}
		fetcher, err = fetch.FromHTTP(httpClient, c.URL)
		if err != nil {
			return nil, err
		}
	default:
		return &noopTemplateStore{}, nil
	}
	baseURL, err := url.Parse(c.LogoBaseURL)
	if err != nil {
		return nil, err
	}
	return &templateStore{
		fetcher:   fetcher,
		baseURL:   baseURL,
		templates: make(map[string]queryResult),
	}, nil
}

// DownlinksConfig defines the configuration for the webhook downlink queue operations.
// For public addresses, the TLS version is preferred when present.
type DownlinksConfig struct {
	PublicAddress    string `name:"public-address" description:"Public address of the HTTP webhooks frontend"`
	PublicTLSAddress string `name:"public-tls-address" description:"Public address of the HTTPS webhooks frontend"`
}
