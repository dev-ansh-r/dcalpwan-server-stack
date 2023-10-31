package oauth

import (
	"go.thethings.network/lorawan-stack/v3/pkg/webui"
)

// UIConfig is the combined configuration for the OAuth and Account UI.
type UIConfig struct {
	webui.TemplateData `name:",squash"`
	FrontendConfig     `name:",squash"`
}

// StackConfig is the configuration of the stack components.
type StackConfig struct {
	IS webui.APIConfig `json:"is" name:"is"`
}

// FrontendConfig is the configuration for the OAuth frontend.
type FrontendConfig struct {
	DocumentationBaseURL   string `json:"documentation_base_url" name:"documentation-base-url" description:"The base URL for generating documentation links"`
	StatusPage             string `json:"status_page_base_url" name:"status-page-base-url" description:"The base URL for generating status page links"`
	Language               string `json:"language" name:"-"`
	StackConfig            `json:"stack_config" name:",squash"`
	EnableUserRegistration bool   `json:"enable_user_registration" name:"-"`
	ConsoleURL             string `json:"console_url" name:"console-url" description:"The URL that points to the root of the Console"`
}

// Config is the configuration for the OAuth server.
type Config struct {
	Mount       string   `name:"mount" description:"Path on the server where the Account application and OAuth services will be served"`
	UI          UIConfig `name:"ui"`
	CSRFAuthKey []byte   `name:"-"`
}
