package templates

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "password_changed",
		email.FSTemplate{
			SubjectTemplate:      "Your password on {{ .Network.Name }} has been changed",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "password_changed.html.tmpl",
			TextTemplateFile:     "password_changed.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("password_changed", &email.NotificationBuilder{
		EmailTemplateName: "password_changed",
		DataBuilder:       newPasswordChangedData,
	})
}

func newPasswordChangedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	return &PasswordChangedData{
		NotificationTemplateData: data,
	}, nil
}

// PasswordChangedData is the data for the password_changed email.
type PasswordChangedData struct {
	email.NotificationTemplateData
}
