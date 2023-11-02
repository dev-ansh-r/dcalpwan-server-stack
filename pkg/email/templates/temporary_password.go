package templates

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "temporary_password",
		email.FSTemplate{
			SubjectTemplate:      "Your temporary password for {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "temporary_password.html.tmpl",
			TextTemplateFile:     "temporary_password.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
}

// TemporaryPasswordData is the data for the temporary_password email.
type TemporaryPasswordData struct {
	email.TemplateData
	TemporaryPassword string
	TTL               time.Duration
}
