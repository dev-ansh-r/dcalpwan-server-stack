package email

// NetworkConfig is the configuration of the network that sends the emails.
// This configuration is used by email templates.
type NetworkConfig struct {
	Name              string `name:"name" description:"The name of the network"`
	IdentityServerURL string `name:"identity-server-url" description:"The URL of the Identity Server"`
	ConsoleURL        string `name:"console-url" description:"The URL of the Console"`
	AssetsBaseURL     string `name:"assets-base-url" description:"The base URL to the email assets"`
	BrandingBaseURL   string `name:"branding-base-url" description:"The base URL to the email branding assets"`
}

// Config is the configuration for sending emails.
type Config struct {
	SenderName    string        `name:"sender-name" description:"The name of the sender"`
	SenderAddress string        `name:"sender-address" description:"The address of the sender"`
	Provider      string        `name:"provider" description:"Email provider to use"`
	Network       NetworkConfig `name:"network" description:"The network of the sender"`
}
