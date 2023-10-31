package commands

import (
	"os"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	selectApplicationWebhookFlags = util.NormalizedFlagSet()

	selectAllApplicationWebhookFlags = util.SelectAllFlagSet("application webhook")
)

func applicationWebhookIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.String("webhook-id", "", "")
	return flagSet
}

var errNoWebhookID = errors.DefineInvalidArgument("no_webhook_id", "no webhook ID set")

func getApplicationWebhookID(flagSet *pflag.FlagSet, args []string) (*ttnpb.ApplicationWebhookIdentifiers, error) {
	applicationID, _ := flagSet.GetString("application-id")
	webhookID, _ := flagSet.GetString("webhook-id")
	switch len(args) {
	case 0:
	case 1:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 2:
		applicationID = args[0]
		webhookID = args[1]
	default:
		logger.Warn("Multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		webhookID = args[1]
	}
	if applicationID == "" {
		return nil, errNoApplicationID.New()
	}
	if webhookID == "" {
		return nil, errNoWebhookID.New()
	}
	return &ttnpb.ApplicationWebhookIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: applicationID},
		WebhookId:      webhookID,
	}, nil
}

var (
	applicationsWebhooksCommand = &cobra.Command{
		Use:     "webhooks",
		Aliases: []string{"webhook"},
		Short:   "Application webhooks commands",
	}
	applicationsWebhooksGetFormatsCommand = &cobra.Command{
		Use:     "get-formats",
		Aliases: []string{"formats", "list-formats"},
		Short:   "Get the available formats for application webhooks",
		RunE: func(cmd *cobra.Command, args []string) error {
			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationWebhookRegistryClient(as).GetFormats(ctx, ttnpb.Empty)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsWebhooksGetCommand = &cobra.Command{
		Use:     "get [application-id] [webhook-id]",
		Aliases: []string{"info"},
		Short:   "Get the properties of an application webhook",
		RunE: func(cmd *cobra.Command, args []string) error {
			webhookID, err := getApplicationWebhookID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationWebhookFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationWebhookFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationWebhookRegistry/Get"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationWebhookRegistryClient(as).Get(ctx, &ttnpb.GetApplicationWebhookRequest{
				Ids:       webhookID,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsWebhooksListCommand = &cobra.Command{
		Use:     "list [application-id]",
		Aliases: []string{"ls"},
		Short:   "List application webhooks",
		RunE: func(cmd *cobra.Command, args []string) error {
			appID := getApplicationID(cmd.Flags(), args)
			if appID == nil {
				return errNoApplicationID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationWebhookFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationWebhookFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationWebhookRegistry/List"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationWebhookRegistryClient(as).List(ctx, &ttnpb.ListApplicationWebhooksRequest{
				ApplicationIds: appID,
				FieldMask:      ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsWebhooksSetCommand = &cobra.Command{
		Use:     "set [application-id] [webhook-id]",
		Aliases: []string{"update"},
		Short:   "Set the properties of an application webhook",
		RunE: func(cmd *cobra.Command, args []string) error {
			webhook := &ttnpb.ApplicationWebhook{}
			paths, err := webhook.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			webhookID, err := getApplicationWebhookID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			webhook.Ids = webhookID

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationWebhookRegistryClient(as).Set(ctx, &ttnpb.SetApplicationWebhookRequest{
				Webhook:   webhook,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsWebhooksDeleteCommand = &cobra.Command{
		Use:     "delete [application-id] [webhook-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an application webhook",
		RunE: func(cmd *cobra.Command, args []string) error {
			webhookID, err := getApplicationWebhookID(cmd.Flags(), args)
			if err != nil {
				return err
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewApplicationWebhookRegistryClient(as).Delete(ctx, webhookID)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	ttnpb.AddSelectFlagsForApplicationWebhook(selectApplicationWebhookFlags, "", false)
	applicationsWebhooksCommand.AddCommand(applicationsWebhooksGetFormatsCommand)
	applicationsWebhooksGetCommand.Flags().AddFlagSet(applicationWebhookIDFlags())
	applicationsWebhooksGetCommand.Flags().AddFlagSet(selectApplicationWebhookFlags)
	applicationsWebhooksGetCommand.Flags().AddFlagSet(selectAllApplicationWebhookFlags)
	applicationsWebhooksCommand.AddCommand(applicationsWebhooksGetCommand)
	applicationsWebhooksListCommand.Flags().AddFlagSet(applicationIDFlags())
	applicationsWebhooksListCommand.Flags().AddFlagSet(selectApplicationWebhookFlags)
	applicationsWebhooksListCommand.Flags().AddFlagSet(selectAllApplicationWebhookFlags)
	applicationsWebhooksCommand.AddCommand(applicationsWebhooksListCommand)
	ttnpb.AddSetFlagsForApplicationWebhook(applicationsWebhooksSetCommand.Flags(), "", false)
	flagsplugin.AddAlias(applicationsWebhooksSetCommand.Flags(), "ids.application-ids.application-id", "application-id", flagsplugin.WithHidden(false))
	flagsplugin.AddAlias(applicationsWebhooksSetCommand.Flags(), "ids.webhook-id", "webhook-id", flagsplugin.WithHidden(false))
	applicationsWebhooksCommand.AddCommand(applicationsWebhooksSetCommand)
	applicationsWebhooksDeleteCommand.Flags().AddFlagSet(applicationWebhookIDFlags())
	applicationsWebhooksCommand.AddCommand(applicationsWebhooksDeleteCommand)
	applicationsCommand.AddCommand(applicationsWebhooksCommand)
}
