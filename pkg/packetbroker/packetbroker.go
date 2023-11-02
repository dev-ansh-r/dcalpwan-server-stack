package packetbroker

import (
	"fmt"
	"time"
)

// Default values for Packet Broker IAM.
const (
	DefaultTokenIssuer       = "https://iam.packetbroker.net"
	DefaultTokenURL          = DefaultTokenIssuer + "/token"
	DefaultPublicKeyCacheTTL = 10 * time.Minute
)

// TokenPublicKeysURL returns the URL with public keys with which a token are signed.
func TokenPublicKeysURL(issuer string) string {
	return fmt.Sprintf("%s/.well-known/jwks.json", issuer)
}
