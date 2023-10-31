// ttn-lw-cli is the binary for the Command-line interface of The Things Stack for LoRaWAN.
package main

import (
	"os"

	cli_errors "go.thethings.network/lorawan-stack/v3/cmd/internal/errors"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/commands"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

func main() {
	cmd, err := commands.Root.ExecuteC()
	if err != nil {
		if errors.IsCanceled(err) {
			os.Exit(130)
		}
		cli_errors.PrintStack(os.Stderr, err)
		os.Exit(-1)
	}
	if cmd.Run == nil && cmd.RunE == nil {
		os.Exit(2)
	}
}
