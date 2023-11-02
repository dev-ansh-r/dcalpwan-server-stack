package templates

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "client_requested",
		email.FSTemplate{
			SubjectTemplate:      "A new OAuth client has been requested",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "client_requested.html.tmpl",
			TextTemplateFile:     "client_requested.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("client_requested", &email.NotificationBuilder{
		EmailTemplateName: "client_requested",
		DataBuilder:       newClientRequestedData,
	})
}

func newClientRequestedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var emailMsg ttnpb.CreateClientEmailMessage
	if err := data.Notification().GetData().UnmarshalTo(&emailMsg); err != nil {
		return nil, err
	}
	return &ClientRequestedData{
		NotificationTemplateData: data,
		CreateClientRequest:      emailMsg.GetCreateClientRequest(),
		APIKey:                   emailMsg.GetApiKey(),
	}, nil
}

// ClientRequestedData is the data for the client_requested email.
type ClientRequestedData struct {
	email.NotificationTemplateData
	*ttnpb.CreateClientRequest
	*ttnpb.APIKey
}

func (crd *ClientRequestedData) getAPIKeyName() string {
	if crd.APIKey.GetName() == "" {
		return crd.APIKey.GetId()
	}
	return crd.APIKey.GetName()
}

// SenderType returns the type of the entity that triggered the email.
func (crd *ClientRequestedData) SenderType() string {
	if crd.Notification().GetSenderIds().IDString() == "" {
		return "API key"
	}
	return "User"
}

// SenderTypeMidSentence returns the type of the entity that triggered the email, altered to fit midsentence.
func (crd *ClientRequestedData) SenderTypeMidSentence() string {
	if crd.Notification().GetSenderIds().IDString() == "" {
		return "API key"
	}
	return "user"
}

// Sender returns the name of the User or APIKey used for triggering the email.
func (crd *ClientRequestedData) Sender() string {
	if crd.Notification().GetSenderIds().IDString() == "" {
		return crd.getAPIKeyName()
	}
	return crd.Notification().GetSenderIds().IDString()
}
