package templates

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "validate",
		email.FSTemplate{
			SubjectTemplate:      "Please confirm your email address for {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "validate.html.tmpl",
			TextTemplateFile:     "validate.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
}

// ValidateData is the data for the validate email.
type ValidateData struct {
	email.TemplateData
	*ttnpb.EntityIdentifiers
	ID    string
	Token string
	TTL   time.Duration
}
