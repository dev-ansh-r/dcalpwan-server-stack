package rpcclient

import (
	"context"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/credentials"
)

type clientCredentials struct {
	tokenSource oauth2.TokenSource
	accessToken string
	insecure    bool
}

var (
	errFetchOAuth2Token   = errors.DefineAborted("fetch_oauth2_token", "fetch OAuth 2.0 token")
	errInvalidOAuth2Token = errors.DefineAborted("invalid_oauth2_token", "invalid OAuth 2.0 token")
)

func (c *clientCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, errFetchOAuth2Token.WithCause(err)
	}
	if !token.Valid() {
		return nil, errInvalidOAuth2Token.New()
	}
	return map[string]string{
		"authorization": fmt.Sprintf("%s %s", token.Type(), token.AccessToken),
	}, nil
}

func (c *clientCredentials) RequireTransportSecurity() bool {
	return !c.insecure
}

// OAuth2 returns per RPC client credentials using the OAuth Client Credentials flow.
func OAuth2(tokenSource oauth2.TokenSource, insecure bool) credentials.PerRPCCredentials {
	return &clientCredentials{
		tokenSource: tokenSource,
		insecure:    insecure,
	}
}
