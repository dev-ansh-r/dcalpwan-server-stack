package templates

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "user_requested",
		email.FSTemplate{
			SubjectTemplate:      "A new user has requested to join {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "user_requested.html.tmpl",
			TextTemplateFile:     "user_requested.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("user_requested", &email.NotificationBuilder{
		EmailTemplateName: "user_requested",
		DataBuilder:       newUserRequestedData,
	})
}

func newUserRequestedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var nData ttnpb.CreateUserRequest
	if err := data.Notification().GetData().UnmarshalTo(&nData); err != nil {
		return nil, err
	}
	return &UserRequestedData{
		NotificationTemplateData: data,
		CreateUserRequest:        &nData,
	}, nil
}

// UserRequestedData is the data for the user_requested email.
type UserRequestedData struct {
	email.NotificationTemplateData
	*ttnpb.CreateUserRequest
}
