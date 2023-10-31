package commands

import (
	"github.com/spf13/cobra"
)

var (
	drDBCommand = &cobra.Command{
		Use:   "dr-db",
		Short: "Device Repository commands",
	}
	drInitCommand = &cobra.Command{
		Use:   "init",
		Short: "Fetch Device Repository files and generate index",
		RunE: func(cmd *cobra.Command, args []string) error {
			overwrite, _ := cmd.Flags().GetBool("overwrite")

			return config.DR.Initialize(ctx, config.Blob, overwrite)
		},
	}
)

func init() {
	Root.AddCommand(drDBCommand)

	drInitCommand.Flags().Bool("overwrite", true, "Overwrite existing index files")
	drDBCommand.AddCommand(drInitCommand)
}
