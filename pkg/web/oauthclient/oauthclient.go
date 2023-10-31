package oauthclient

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"golang.org/x/oauth2"
)

// OAuthClient is the OAuth client component.
type OAuthClient struct {
	component       *component.Component
	config          Config
	oauthConfig     OAuth2ConfigProvider
	nextKey         string
	callback        Callback
	authCodeURLOpts OAuth2AuthCodeURLOptionsProvider
	schemaDecoder   *schema.Decoder
}

var errNoOAuthConfig = errors.DefineInvalidArgument("no_oauth_config", "no OAuth configuration found for the OAuth client")

func (c Config) isZero() bool {
	return (c.AuthorizeURL == "" || c.TokenURL == "" || c.ClientID == "" || c.ClientSecret == "") && !c.customProvider
}

func (oc *OAuthClient) getMountPath() string {
	var path string
	u, err := url.Parse(oc.config.RootURL)
	if err != nil || u.Path == "" {
		path = "/"
	} else {
		path = u.Path
	}

	return path
}

func (oc *OAuthClient) withHTTPClient(ctx context.Context) (context.Context, error) {
	client, err := oc.component.HTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, oauth2.HTTPClient, client), nil
}

// New returns a new OAuth client instance.
func New(c *component.Component, config Config, opts ...Option) (*OAuthClient, error) {
	oc := &OAuthClient{
		component:     c,
		config:        config,
		nextKey:       "next",
		schemaDecoder: schema.NewDecoder(),
	}
	oc.schemaDecoder.IgnoreUnknownKeys(true)

	oc.callback = oc.defaultCallback
	oc.oauthConfig = oc.defaultOAuthConfig
	oc.authCodeURLOpts = oc.defaultAuthCodeURLOptions

	for _, opt := range opts {
		opt(oc)
	}

	if oc.config.isZero() {
		return nil, errNoOAuthConfig.New()
	}
	return oc, nil
}

type ctxKeyType struct{}

var ctxKey ctxKeyType

func (oc *OAuthClient) configFromContext(ctx context.Context) *Config {
	if config, ok := ctx.Value(ctxKey).(*Config); ok {
		return config
	}
	return &oc.config
}

func (oc *OAuthClient) defaultOAuthConfig(ctx context.Context) (*oauth2.Config, error) {
	config := oc.configFromContext(ctx)

	authorizeURL := config.AuthorizeURL
	redirectURL := fmt.Sprintf("%s/oauth/callback", strings.TrimSuffix(config.RootURL, "/"))
	if oauthRootURL, err := url.Parse(config.RootURL); err == nil {
		rootURL := (&url.URL{Scheme: oauthRootURL.Scheme, Host: oauthRootURL.Host}).String()
		if strings.HasPrefix(authorizeURL, rootURL) {
			authorizeURL = strings.TrimPrefix(authorizeURL, rootURL)
			redirectURL = strings.TrimPrefix(redirectURL, rootURL)
		}
	}

	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			TokenURL:  config.TokenURL,
			AuthURL:   authorizeURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}, nil
}
