package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	errStorageIntegrationNotAvailable = errors.DefineUnimplemented("storage_integration_not_available", "Storage Integration not available")

	storageDBCommand = &cobra.Command{
		Use:   "storage-db",
		Short: "Manage the Storage Integration database",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("The storage integration is not available in the open-source version of The Things Stack.")
			fmt.Println("For more information, see https://www.thethingsindustries.com/docs/integrations/storage/")
			return errStorageIntegrationNotAvailable.New()
		},
	}
)

func init() {
	Root.AddCommand(storageDBCommand)
}
