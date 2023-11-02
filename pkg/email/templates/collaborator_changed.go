package templates

import (
	"context"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "collaborator_changed",
		email.FSTemplate{
			SubjectTemplate:      "A collaborator of your {{ .Notification.EntityIds.EntityType }} has been changed",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "collaborator_changed.html.tmpl",
			TextTemplateFile:     "collaborator_changed.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("collaborator_changed", &email.NotificationBuilder{
		EmailTemplateName: "collaborator_changed",
		DataBuilder:       newCollaboratorChangedData,
	})
}

func newCollaboratorChangedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var nData ttnpb.Collaborator
	if err := data.Notification().GetData().UnmarshalTo(&nData); err != nil {
		return nil, err
	}
	return &CollaboratorChangedData{
		NotificationTemplateData: data,
		Collaborator:             &nData,
	}, nil
}

// CollaboratorChangedData is the data for the collaborator_changed email.
type CollaboratorChangedData struct {
	email.NotificationTemplateData
	*ttnpb.Collaborator
}

// ConsoleURL returns the URL to the API key in the Console.
func (d *CollaboratorChangedData) ConsoleURL() string {
	return fmt.Sprintf("%s/collaborators/%s/%s", d.NotificationTemplateData.ConsoleURL(), d.Collaborator.EntityType(), d.Collaborator.IDString())
}
