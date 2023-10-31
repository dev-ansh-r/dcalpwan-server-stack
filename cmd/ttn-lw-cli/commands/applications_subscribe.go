package commands

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var applicationsSubscribeCommand = &cobra.Command{
	Use:     "subscribe [application-id]",
	Aliases: []string{"sub"},
	Short:   "Subscribe to application uplink",
	RunE: func(cmd *cobra.Command, args []string) error {
		appID := getApplicationID(cmd.Flags(), args)
		if appID == nil {
			return errNoApplicationID.New()
		}

		as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		stream, err := ttnpb.NewAppAsClient(as).Subscribe(ctx, appID)
		if err != nil {
			return err
		}

		var streamErr error
		go func() {
			defer cancel()
			for {
				up, err := stream.Recv()
				if err != nil {
					streamErr = err
					return
				}
				if err = io.Write(os.Stdout, config.OutputFormat, up); err != nil {
					streamErr = err
					return
				}
			}
		}()

		<-ctx.Done()

		if streamErr != nil {
			return streamErr
		}
		return ctx.Err()
	},
}

func init() {
	applicationsSubscribeCommand.Flags().AddFlagSet(applicationIDFlags())
	applicationsCommand.AddCommand(applicationsSubscribeCommand)
}
