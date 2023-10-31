package console

import (
	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	"go.thethings.network/lorawan-stack/v3/pkg/console"
	"go.thethings.network/lorawan-stack/v3/pkg/web/oauthclient"
	"go.thethings.network/lorawan-stack/v3/pkg/webui"
)

// DefaultConsoleConfig is the default configuration for the Console.
var DefaultConsoleConfig = console.Config{
	OAuth: oauthclient.Config{
		AuthorizeURL:    shared.DefaultOAuthPublicURL + "/authorize",
		LogoutURL:       shared.DefaultOAuthPublicURL + "/logout",
		TokenURL:        shared.DefaultOAuthPublicURL + "/token",
		ClientID:        "console",
		ClientSecret:    "console",
		CrossSiteCookie: false,
	},
	UI: console.UIConfig{
		TemplateData: webui.TemplateData{
			SiteName:      "The Things Stack for LoRaWAN",
			Title:         "Console",
			SubTitle:      "Management platform for The Things Stack for LoRaWAN",
			Language:      "en",
			CanonicalURL:  shared.DefaultConsolePublicURL,
			AssetsBaseURL: shared.DefaultAssetsBaseURL,
			IconPrefix:    "console-",
			CSSFiles:      []string{"console.css"},
			JSFiles:       []string{"libs.bundle.js", "console.js"},
		},
		FrontendConfig: console.FrontendConfig{
			DocumentationBaseURL: "https://thethingsindustries.com/docs",
			AccountURL:           "/oauth",
			StackConfig: console.StackConfig{
				IS:   webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				GS:   webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				NS:   webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				AS:   webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				JS:   webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				EDTC: webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				QRG:  webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				GCS:  webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
				DCS:  webui.APIConfig{Enabled: true, BaseURL: shared.DefaultPublicURL + "/api/v3"},
			},
		},
	},
}
