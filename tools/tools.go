//go:build tools
// +build tools

package tools

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/magefile/mage"
	_ "github.com/mattn/goveralls"
	_ "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/commands"
	_ "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-stack/commands"
)
