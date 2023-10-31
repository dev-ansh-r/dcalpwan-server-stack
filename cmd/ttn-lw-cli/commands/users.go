package commands

import (
	"os"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

var (
	selectUserFlags     = util.NormalizedFlagSet()
	profilePictureFlags = util.NormalizedFlagSet()

	selectAllUserFlags = util.SelectAllFlagSet("user")
)

func userIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("user-id", "", "")
	return flagSet
}

var errNoUserID = errors.DefineInvalidArgument("no_user_id", "no user ID set")

func getUserID(flagSet *pflag.FlagSet, args []string) *ttnpb.UserIdentifiers {
	var userID string
	if len(args) > 0 {
		if len(args) > 1 {
			logger.Warn("Multiple IDs found in arguments, considering only the first")
		}
		userID = args[0]
	} else {
		userID, _ = flagSet.GetString("user-id")
	}
	if userID == "" {
		return nil
	}
	return &ttnpb.UserIdentifiers{UserId: userID}
}

func printPasswordRequirements(msg *ttnpb.IsConfiguration_UserRegistration_PasswordRequirements) {
	if msg == nil {
		return
	}
	logger.Info("Password Requirements:")
	if v := msg.GetMinLength().GetValue(); v > 0 {
		logger.Infof("Minimum length: %d", v)
	}
	if v := msg.GetMaxLength().GetValue(); v > 0 {
		logger.Infof("Maximum length: %d", v)
	}
	if v := msg.GetMinUppercase().GetValue(); v > 0 {
		logger.Infof("Minimum uppercase: %d", v)
	}
	if v := msg.GetMinDigits().GetValue(); v > 0 {
		logger.Infof("Minimum digits: %d", v)
	}
	if v := msg.GetMinSpecial().GetValue(); v > 0 {
		logger.Infof("Minimum special characters: %d", v)
	}
}

var errPasswordMismatch = errors.DefineInvalidArgument("password_mismatch", "password did not match")

var (
	usersCommand = &cobra.Command{
		Use:     "users",
		Aliases: []string{"user", "usr", "u"},
		Short:   "User commands",
	}
	usersListCommand = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths := util.SelectFieldMask(cmd.Flags(), selectUserFlags)
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.UserRegistry/List"].Allowed)

			req := &ttnpb.ListUsersRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			if req.FieldMask == nil {
				req.FieldMask = ttnpb.FieldMask(paths...)
			}
			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, _, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewUserRegistryClient(is).List(ctx, req, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res.Users)
		},
	}
	usersSearchCommand = &cobra.Command{
		Use:   "search",
		Short: "Search for users",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths := util.SelectFieldMask(cmd.Flags(), selectUserFlags)
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.EntityRegistrySearch/SearchUsers"].Allowed)

			req := &ttnpb.SearchUsersRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			var (
				opt      grpc.CallOption
				getTotal func() uint64
			)
			_, _, opt, getTotal = withPagination(cmd.Flags())
			if req.FieldMask == nil {
				req.FieldMask = ttnpb.FieldMask(paths...)
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewEntityRegistrySearchClient(is).SearchUsers(ctx, req, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res.Users)
		},
	}
	usersGetCommand = &cobra.Command{
		Use:     "get [user-id]",
		Aliases: []string{"info"},
		Short:   "Get a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectUserFlags)
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.UserRegistry/Get"].Allowed)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewUserRegistryClient(is).Get(ctx, &ttnpb.GetUserRequest{
				UserIds:   usrID,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	usersCreateCommand = &cobra.Command{
		Use:               "create [user-id]",
		Aliases:           []string{"add", "register"},
		Short:             "Create a user",
		PersistentPreRunE: preRun(optionalAuth),
		RunE: asBulk(func(cmd *cobra.Command, args []string) (err error) {
			usrID := getUserID(cmd.Flags(), args)
			var user ttnpb.User
			user.State = ttnpb.State_STATE_APPROVED // This may not be honored by the server.
			if inputDecoder != nil {
				err := inputDecoder.Decode(&user)
				if err != nil {
					return err
				}
			}
			_, err = user.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			if usrID.GetUserId() != "" {
				user.Ids = &ttnpb.UserIdentifiers{UserId: usrID.GetUserId()}
			}
			if user.GetIds().GetUserId() == "" {
				return errNoUserID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}

			isConfig, err := ttnpb.NewIsClient(is).GetConfiguration(ctx, &ttnpb.GetIsConfigurationRequest{})
			if err != nil {
				if errors.IsUnimplemented(err) {
					logger.Warn("Could not get user registration configuration from the server, but we can continue without it")
				} else {
					return err
				}
			}
			registrationConfig := isConfig.GetConfiguration().GetUserRegistration()

			if user.Password == "" {
				printPasswordRequirements(registrationConfig.GetPasswordRequirements())
				pw, err := gopass.GetPasswdPrompt("Please enter password:", true, os.Stdin, os.Stderr)
				if err != nil {
					return err
				}
				user.Password = string(pw)
				pw, err = gopass.GetPasswdPrompt("Please confirm password:", true, os.Stdin, os.Stderr)
				if err != nil {
					return err
				}
				if string(pw) != user.Password {
					return errPasswordMismatch.New()
				}
			}

			if !isConfig.GetConfiguration().GetProfilePicture().GetDisableUpload().GetValue() {
				if profilePicture, err := cmd.Flags().GetString("profile_picture"); err == nil && profilePicture != "" {
					user.ProfilePicture, err = readPicture(profilePicture)
					if err != nil {
						return err
					}
				}
			}

			invitationToken, _ := cmd.Flags().GetString("invitation-token")

			res, err := ttnpb.NewUserRegistryClient(is).Create(ctx, &ttnpb.CreateUserRequest{
				User:            &user,
				InvitationToken: invitationToken,
			})
			if err != nil {
				return err
			}

			if err = io.Write(os.Stdout, config.OutputFormat, res); err != nil {
				return err
			}

			if registrationConfig.GetContactInfoValidation().GetRequired().GetValue() {
				logger.Warn("You need to confirm your email address before you can use your user account")
			}
			if registrationConfig.GetAdminApproval().GetRequired().GetValue() {
				logger.Warn("Your user account needs to be approved by an admin before you can use it")
			}

			return nil
		}),
	}
	usersSetCommand = &cobra.Command{
		Use:     "set [user-id]",
		Aliases: []string{"update"},
		Short:   "Set properties of a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			var user ttnpb.User
			paths, err := user.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			user.Ids = usrID
			if profilePicture, err := cmd.Flags().GetString("profile_picture"); err == nil && profilePicture != "" {
				user.ProfilePicture, err = readPicture(profilePicture)
				if err != nil {
					return err
				}
				paths = append(paths, "profile_picture")
			}

			if len(paths) == 0 {
				logger.Warn("No fields selected, won't update anything")
				return nil
			}
			if ttnpb.ContainsField("password", paths) {
				logger.Warn("Most servers do not allow changing the password of a user like this.")
				logger.Warnf("Use \"%s [user-id]\" to change your password.", usersUpdatePasswordCommand.CommandPath())
				logger.Warnf("Use \"%s [user-id]\" if you forgot your password and want to request a temporary password.", usersForgotPasswordCommand.CommandPath())
				logger.Warnf("Alternatively, admins may use \"%s [user-id] --temporary-password [pass]\" to set a temporary password for a user.", cmd.CommandPath())
				logger.Warn("The user can then use this temporary password when changing their password.")
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewUserRegistryClient(is).Update(ctx, &ttnpb.UpdateUserRequest{
				User:      &user,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			res.SetFields(&user, "ids")
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	usersForgotPasswordCommand = &cobra.Command{
		Use:               "forgot-password [user-id]",
		Short:             "Request a temporary user password",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}

			usrID.Email, _ = cmd.Flags().GetString("email")

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewUserRegistryClient(is).CreateTemporaryPassword(ctx, &ttnpb.CreateTemporaryPasswordRequest{
				UserIds: usrID,
			})

			return err
		},
	}
	usersUpdatePasswordCommand = &cobra.Command{
		Use:               "update-password [user-id]",
		Aliases:           []string{"change-password"},
		Short:             "Update a user password",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			refreshToken() // NOTE: ignore errors.
			optionalAuth()

			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}

			isConfig, err := ttnpb.NewIsClient(is).GetConfiguration(ctx, &ttnpb.GetIsConfigurationRequest{})
			if err != nil {
				if errors.IsUnimplemented(err) {
					logger.Warn("Could not get password requirements from the server, but we can continue without knowing those")
				} else {
					return err
				}
			}

			old, _ := cmd.Flags().GetString("old")
			if old == "" {
				pw, err := gopass.GetPasswdPrompt("Please enter old password:", true, os.Stdin, os.Stderr)
				if err != nil {
					return err
				}
				old = string(pw)
			}

			new, _ := cmd.Flags().GetString("new")
			if new == "" {
				printPasswordRequirements(isConfig.GetConfiguration().GetUserRegistration().GetPasswordRequirements())
				pw, err := gopass.GetPasswdPrompt("Please enter new password:", true, os.Stdin, os.Stderr)
				if err != nil {
					return err
				}
				new = string(pw)
			}

			revokeAllAccess, _ := cmd.Flags().GetBool("revoke-all-access")

			res, err := ttnpb.NewUserRegistryClient(is).UpdatePassword(ctx, &ttnpb.UpdateUserPasswordRequest{
				UserIds:         usrID,
				Old:             old,
				New:             new,
				RevokeAllAccess: revokeAllAccess,
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	usersDeleteCommand = &cobra.Command{
		Use:     "delete [user-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewUserRegistryClient(is).Delete(ctx, usrID)
			if err != nil {
				return err
			}

			return nil
		},
	}
	usersRestoreCommand = &cobra.Command{
		Use:   "restore [user-id]",
		Short: "Restore a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewUserRegistryClient(is).Restore(ctx, usrID)
			if err != nil {
				return err
			}

			return nil
		},
	}
	usersContactInfoCommand = contactInfoCommands("user", func(cmd *cobra.Command, args []string) (*ttnpb.EntityIdentifiers, error) {
		usrID := getUserID(cmd.Flags(), args)
		if usrID == nil {
			return nil, errNoUserID.New()
		}
		return usrID.GetEntityIdentifiers(), nil
	})

	usersPurgeCommand = &cobra.Command{
		Use:     "purge [user-id]",
		Aliases: []string{"permanent-delete", "hard-delete"},
		Short:   "Purge a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err
			}
			if !confirmChoice(userPurgeWarning, force) {
				return errNoConfirmation.New()
			}
			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewUserRegistryClient(is).Purge(ctx, usrID)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	profilePictureFlags.String("profile_picture", "", "upload the profile picture from this file")
	ttnpb.AddSelectFlagsForUser(selectUserFlags, "", false)
	ttnpb.AddSetFlagsForListUsersRequest(usersListCommand.Flags(), "", false)
	usersListCommand.Flags().AddFlagSet(selectUserFlags)
	usersListCommand.Flags().AddFlagSet(selectAllUserFlags)
	usersCommand.AddCommand(usersListCommand)
	ttnpb.AddSetFlagsForSearchUsersRequest(usersSearchCommand.Flags(), "", false)
	usersSearchCommand.Flags().AddFlagSet(selectUserFlags)
	usersSearchCommand.Flags().AddFlagSet(selectAllUserFlags)
	usersCommand.AddCommand(usersSearchCommand)
	usersGetCommand.Flags().AddFlagSet(userIDFlags())
	usersGetCommand.Flags().AddFlagSet(selectUserFlags)
	usersGetCommand.Flags().AddFlagSet(selectAllUserFlags)
	usersCommand.AddCommand(usersGetCommand)
	ttnpb.AddSetFlagsForUser(usersCreateCommand.Flags(), "", false)
	flagsplugin.AddAlias(usersCreateCommand.Flags(), "ids.user-id", "user-id", flagsplugin.WithHidden(false))
	usersCreateCommand.Flags().AddFlagSet(profilePictureFlags)
	usersCreateCommand.Flags().Lookup("state").DefValue = ttnpb.State_STATE_APPROVED.String()
	usersCreateCommand.Flags().String("invitation-token", "", "")
	usersCommand.AddCommand(usersCreateCommand)
	ttnpb.AddSetFlagsForUser(usersSetCommand.Flags(), "", false)
	flagsplugin.AddAlias(usersSetCommand.Flags(), "ids.user-id", "user-id", flagsplugin.WithHidden(false))
	usersSetCommand.Flags().AddFlagSet(profilePictureFlags)
	usersCommand.AddCommand(usersSetCommand)
	usersForgotPasswordCommand.Flags().AddFlagSet(userIDFlags())
	usersForgotPasswordCommand.Flags().String("email", "", "")
	usersCommand.AddCommand(usersForgotPasswordCommand)
	usersUpdatePasswordCommand.Flags().AddFlagSet(userIDFlags())
	usersUpdatePasswordCommand.Flags().String("old", "", "")
	usersUpdatePasswordCommand.Flags().String("new", "", "")
	usersUpdatePasswordCommand.Flags().Bool("revoke-all-access", true, "revoke all sessions and access tokens after the password is updated")
	usersCommand.AddCommand(usersUpdatePasswordCommand)
	usersDeleteCommand.Flags().AddFlagSet(userIDFlags())
	usersCommand.AddCommand(usersDeleteCommand)
	usersRestoreCommand.Flags().AddFlagSet(userIDFlags())
	usersCommand.AddCommand(usersRestoreCommand)
	usersContactInfoCommand.PersistentFlags().AddFlagSet(userIDFlags())
	usersCommand.AddCommand(usersContactInfoCommand)
	usersPurgeCommand.Flags().AddFlagSet(userIDFlags())
	usersPurgeCommand.Flags().AddFlagSet(forceFlags())
	usersCommand.AddCommand(usersPurgeCommand)
	Root.AddCommand(usersCommand)
}
