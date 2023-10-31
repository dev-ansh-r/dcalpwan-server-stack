package oauthclient

// Config is the configuration for the OAuth client.
type Config struct {
	AuthorizeURL string `name:"authorize-url" description:"The OAuth Authorize URL"`
	LogoutURL    string `name:"logout-url" description:"The OAuth Logout URL"`
	TokenURL     string `name:"token-url" description:"The OAuth Token Exchange URL"`
	RootURL      string `name:"-"`

	ClientID     string `name:"client-id" description:"The OAuth client ID"`
	ClientSecret string `name:"client-secret" description:"The OAuth client secret" json:"-"`

	CrossSiteCookie bool `name:"cross-site-cookie" description:"Whether to make OAuth cookies accessible cross-site"`

	StateCookieName string `name:"-"`
	AuthCookieName  string `name:"-"`

	customProvider bool
}
