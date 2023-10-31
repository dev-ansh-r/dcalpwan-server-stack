package commands

import (
	"os"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var authInfo = &cobra.Command{
	Use:    "auth-info",
	Hidden: true,
	Short:  "Get auth info",
	RunE: func(cmd *cobra.Command, args []string) error {
		is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
		if err != nil {
			return err
		}
		res, err := ttnpb.NewEntityAccessClient(is).AuthInfo(ctx, ttnpb.Empty)
		if err != nil {
			return err
		}

		return io.Write(os.Stdout, config.OutputFormat, res)
	},
}

func init() {
	Root.AddCommand(authInfo)
}
