// ttn-lw-stack is the binary that runs the entire The Things Stack for LoRaWAN.
package main

import (
	"os"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/errors"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-stack/commands"
)

func main() {
	cmd, err := commands.Root.ExecuteC()
	if err != nil {
		errors.PrintStack(os.Stderr, err)
		os.Exit(-1)
	}
	if cmd.Run == nil && cmd.RunE == nil {
		os.Exit(2)
	}
}
