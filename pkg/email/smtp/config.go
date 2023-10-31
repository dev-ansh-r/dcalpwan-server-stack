// Package smtp implements SMTP as email provider.
package smtp

import (
	"crypto/tls"
)

// Config for the SMTP email provider.
type Config struct {
	Address     string `name:"address" description:"SMTP server address"`
	Username    string `name:"username" description:"Username to authenticate with"`
	Password    string `name:"password" description:"Password to authenticate with"`
	Connections int    `name:"connections" description:"Maximum number of connections to the SMTP server"`
	TLSConfig   *tls.Config
}
