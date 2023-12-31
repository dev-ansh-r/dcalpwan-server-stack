package identityserver

import (
	"path"
	"time"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver"
	"go.thethings.network/lorawan-stack/v3/pkg/oauth"
	"go.thethings.network/lorawan-stack/v3/pkg/webui"
)

// DefaultIdentityServerConfig is the default configuration for the Identity Server.
var DefaultIdentityServerConfig = identityserver.Config{
	DatabaseURI: "postgres://root:yKxvwr5POAR4@ep-green-cherry-82769755.us-east-2.aws.neon.tech/lorawan_dev?sslmode=require",
	OAuth: oauth.Config{
		Mount: "/oauth",
		UI: oauth.UIConfig{
			TemplateData: webui.TemplateData{
				SiteName:      "The dcalpwan Server for LoRaWAN Networks",
				Title:         "Account",
				Language:      "en",
				CanonicalURL:  shared.DefaultOAuthPublicURL,
				AssetsBaseURL: shared.DefaultAssetsBaseURL,
				IconPrefix:    "oauth-",
				CSSFiles:      []string{"account.css"},
				JSFiles:       []string{"libs.bundle.js", "account.js"},
			},
			FrontendConfig: oauth.FrontendConfig{
				DocumentationBaseURL: "https://github.com/dev-ansh-r/lw-stack/blob/main/DEVELOPMENT.md",
				StackConfig: oauth.StackConfig{
					IS: webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				},
				ConsoleURL: "/console",
			},
		},
	},
}

func init() {
	DefaultIdentityServerConfig.AuthCache.MembershipTTL = 10 * time.Minute
	DefaultIdentityServerConfig.UserRegistration.Enabled = true
	DefaultIdentityServerConfig.UserRegistration.Invitation.TokenTTL = 7 * 24 * time.Hour
	DefaultIdentityServerConfig.UserRegistration.ContactInfoValidation.TokenTTL = 2 * 24 * time.Hour
	DefaultIdentityServerConfig.UserRegistration.ContactInfoValidation.RetryInterval = time.Hour
	DefaultIdentityServerConfig.UserRegistration.PasswordRequirements.MinLength = 8
	DefaultIdentityServerConfig.UserRegistration.PasswordRequirements.MaxLength = 1000
	DefaultIdentityServerConfig.UserRegistration.PasswordRequirements.MinUppercase = 1
	DefaultIdentityServerConfig.UserRegistration.PasswordRequirements.MinDigits = 1
	DefaultIdentityServerConfig.UserRegistration.PasswordRequirements.RejectUserID = true
	DefaultIdentityServerConfig.Email.Network.Name = DefaultIdentityServerConfig.OAuth.UI.SiteName
	DefaultIdentityServerConfig.Email.Network.IdentityServerURL = shared.DefaultOAuthPublicURL
	DefaultIdentityServerConfig.Email.Network.ConsoleURL = shared.DefaultConsolePublicURL
	DefaultIdentityServerConfig.Email.Network.AssetsBaseURL = shared.DefaultPublicURL + shared.DefaultHTTPConfig.Static.Mount
	DefaultIdentityServerConfig.ProfilePicture.Bucket = "profile_pictures"
	DefaultIdentityServerConfig.ProfilePicture.BucketURL = path.Join(shared.DefaultAssetsBaseURL, "blob", "profile_pictures")
	DefaultIdentityServerConfig.ProfilePicture.UseGravatar = true
	DefaultIdentityServerConfig.EndDevicePicture.Bucket = "end_device_pictures"
	DefaultIdentityServerConfig.EndDevicePicture.BucketURL = path.Join(shared.DefaultAssetsBaseURL, "blob", "end_device_pictures")
	DefaultIdentityServerConfig.UserRights.CreateApplications = true
	DefaultIdentityServerConfig.UserRights.CreateClients = true
	DefaultIdentityServerConfig.UserRights.CreateGateways = true
	DefaultIdentityServerConfig.UserRights.CreateOrganizations = true
	DefaultIdentityServerConfig.CollaboratorRights.SetOthersAsContacts = true
	DefaultIdentityServerConfig.LoginTokens.TokenTTL = time.Hour
	DefaultIdentityServerConfig.Delete.Restore = 24 * time.Hour
}
