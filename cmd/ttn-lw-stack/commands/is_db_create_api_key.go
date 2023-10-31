package commands

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	is "go.thethings.network/lorawan-stack/v3/pkg/identityserver"
	bunstore "go.thethings.network/lorawan-stack/v3/pkg/identityserver/bunstore"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	storeutil "go.thethings.network/lorawan-stack/v3/pkg/util/store"
)

var (
	errExpiryDateInPast  = errors.DefineInvalidArgument("expiry_date_invalid", "expiry date is in the past")
	errInvalidDateFormat = errors.DefineInvalidArgument("expiry_date_format_invalid", "invalid expiry date format (RFC3339: YYYY-MM-DDTHH:MM:SSZ)")
)

var createAPIKeyCommand = &cobra.Command{
	Use:   "create-user-api-key",
	Short: "Create an API key with full rights on the user in the Identity Server database",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		logger.Info("Connecting to Identity Server database...")

		db, err := storeutil.OpenDB(ctx, config.IS.DatabaseURI)
		if err != nil {
			return err
		}
		bunDB := bun.NewDB(db, pgdialect.New())
		st, err := bunstore.NewStore(ctx, bunDB)
		if err != nil {
			return err
		}
		defer db.Close()

		userID, err := cmd.Flags().GetString("user-id")
		if err != nil {
			return err
		}
		name, _ := cmd.Flags().GetString("name")

		expiry, _ := cmd.Flags().GetString("api-key-expiry")
		var expiryDate *time.Time

		if expiry != "" {
			expiryDate, err := time.Parse(time.RFC3339, expiry)
			if err != nil {
				return errInvalidDateFormat.New()
			}
			if expiryDate.Before(time.Now()) {
				return errExpiryDateInPast.New()
			}
		}

		usr := &ttnpb.User{
			Ids: &ttnpb.UserIdentifiers{UserId: userID},
		}
		rights := []ttnpb.Right{ttnpb.Right_RIGHT_ALL}
		key, token, err := is.GenerateAPIKey(ctx, name, expiryDate, rights...)
		if err != nil {
			return err
		}
		key, err = st.CreateAPIKey(ctx, usr.GetEntityIdentifiers(), key)
		if err != nil {
			return err
		}
		key.Key = token
		logger.Infof("API key ID: %s", key.Id)
		logger.Infof("API key value: %s", key.Key)
		logger.Warn("The API key value will never be shown again")
		logger.Warn("Make sure to copy it to a safe place")

		return io.Write(os.Stdout, config.OutputFormat, key)
	},
}

func init() {
	createAPIKeyCommand.Flags().String("user-id", "admin", "User ID")
	createAPIKeyCommand.Flags().String("name", "admin-api-key", "API key name")
	createAPIKeyCommand.Flags().String("api-key-expiry", "", "(YYYY-MM-DDTHH:MM:SSZ)")
	isDBCommand.AddCommand(createAPIKeyCommand)
}
