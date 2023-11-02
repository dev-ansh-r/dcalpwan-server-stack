package templates

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "invitation",
		email.FSTemplate{
			SubjectTemplate:      "You have been invited to join {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "invitation.html.tmpl",
			TextTemplateFile:     "invitation.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
}

// InvitationData is the data for the invitation email.
type InvitationData struct {
	email.TemplateData
	SenderIds       *ttnpb.UserIdentifiers
	InvitationToken string
	TTL             time.Duration
}
