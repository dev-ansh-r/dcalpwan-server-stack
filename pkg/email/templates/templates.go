//go:generate ./generate.sh

// Package templates is responsible for handling the email templates sent by The Things Stack.
package templates

import "embed"

//go:embed "*.tmpl"
var fsys embed.FS
