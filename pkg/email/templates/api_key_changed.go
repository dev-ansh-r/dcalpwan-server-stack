package templates

import (
	"context"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "api_key_changed",
		email.FSTemplate{
			SubjectTemplate:      "An API key of your {{ .Notification.EntityIds.EntityType }} has been changed",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "api_key_changed.html.tmpl",
			TextTemplateFile:     "api_key_changed.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("api_key_changed", &email.NotificationBuilder{
		EmailTemplateName: "api_key_changed",
		DataBuilder:       newAPIKeyChangedData,
	})
}

func newAPIKeyChangedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var nData ttnpb.APIKey
	if err := data.Notification().GetData().UnmarshalTo(&nData); err != nil {
		return nil, err
	}
	return &APIKeyChangedData{
		NotificationTemplateData: data,
		APIKey:                   &nData,
	}, nil
}

// APIKeyChangedData is the data for the api_key_changed email.
type APIKeyChangedData struct {
	email.NotificationTemplateData
	*ttnpb.APIKey
}

// ConsoleURL returns the URL to the API key in the Console.
func (a *APIKeyChangedData) ConsoleURL() string {
	if a.Notification().GetEntityIds().EntityType() == "user" {
		return fmt.Sprintf("%s/user/api-keys/%s", a.Network().ConsoleURL, a.APIKey.GetId())
	}
	return fmt.Sprintf("%s/api-keys/%s", a.NotificationTemplateData.ConsoleURL(), a.APIKey.GetId())
}
