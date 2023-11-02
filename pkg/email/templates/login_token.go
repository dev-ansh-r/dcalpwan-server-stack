package templates

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "login_token",
		email.FSTemplate{
			SubjectTemplate:      "Your login token for {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "login_token.html.tmpl",
			TextTemplateFile:     "login_token.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
}

// LoginTokenData is the data for the login_token email.
type LoginTokenData struct {
	email.TemplateData
	LoginToken string
	TTL        time.Duration
}
