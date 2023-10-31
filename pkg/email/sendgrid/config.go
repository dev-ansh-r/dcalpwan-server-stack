package sendgrid

// Config for the SendGrid email provider.
type Config struct {
	APIKey      string `name:"api-key" description:"The SendGrid API key to use"`
	SandboxMode bool   `name:"sandbox" description:"Use SendGrid sandbox mode for testing"`
}
