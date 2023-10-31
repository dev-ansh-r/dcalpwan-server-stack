package commands

import (
	"io"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

// asBulk enables some commands to do bulk operations.
// If there is a non-nil input decoder, asBulk keeps executing the same command
// until it returns an error (the input decoder returns io.EOF when it's done).
// If the input decoder is nil, the command is executed only once.
func asBulk(runE func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) (err error) {
		if inputDecoder == nil {
			return runE(cmd, args)
		}
		for {
			err = runE(cmd, args)
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return err
			}
		}
	}
}
