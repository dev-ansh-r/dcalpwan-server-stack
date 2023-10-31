package oauthclient

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// Option is an OAuth2Client configuration option.
type Option func(*OAuthClient)

// OAuth2ConfigProvider provides an OAuth2 client config based on the context.
type OAuth2ConfigProvider func(context.Context) (*oauth2.Config, error)

// WithOAuth2ConfigProvider overrides the default OAuth2 configuration provider.
func WithOAuth2ConfigProvider(provider OAuth2ConfigProvider) Option {
	return func(o *OAuthClient) {
		o.config.customProvider = true
		o.oauthConfig = provider
	}
}

// WithNextKey overrides the default query parameter used for callback return.
func WithNextKey(key string) Option {
	return func(o *OAuthClient) {
		o.nextKey = key
	}
}

// Callback occurs after the OAuth2 token exchange has been performed successfully.
type Callback func(http.ResponseWriter, *http.Request, *oauth2.Token, string) error

// WithCallback adds a callback to be executed at the end of the OAuth2
// token exchange.
func WithCallback(cb Callback) Option {
	return func(o *OAuthClient) {
		o.callback = cb
	}
}

// OAuth2AuthCodeURLOptionsProvider provides OAuth2 authorization URL options
// based on the context.
type OAuth2AuthCodeURLOptionsProvider func(context.Context) ([]oauth2.AuthCodeOption, error)

// WithAuthCodeURLOptions changes the OAuth2 authorization URL options provided to the
// oauth2 package.
func WithAuthCodeURLOptions(provider OAuth2AuthCodeURLOptionsProvider) Option {
	return func(o *OAuthClient) {
		o.authCodeURLOpts = provider
	}
}
