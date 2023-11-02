package templates

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func init() {
	tmpl, err := email.NewTemplateFS(
		fsys, "entity_state_changed",
		email.FSTemplate{
			SubjectTemplate:      "The state of your {{ .Notification.EntityIds.EntityType }} has been changed",
			HTMLTemplateBaseFile: "base.html.tmpl",
			HTMLTemplateFile:     "entity_state_changed.html.tmpl",
			TextTemplateFile:     "entity_state_changed.txt.tmpl",
		},
	)
	if err != nil {
		panic(err)
	}
	email.RegisterTemplate(tmpl)
	email.RegisterNotification("entity_state_changed", &email.NotificationBuilder{
		EmailTemplateName: "entity_state_changed",
		DataBuilder:       newEntityStateChangedData,
	})
}

func newEntityStateChangedData(_ context.Context, data email.NotificationTemplateData) (email.NotificationTemplateData, error) {
	var nData ttnpb.EntityStateChangedNotification
	if err := data.Notification().GetData().UnmarshalTo(&nData); err != nil {
		return nil, err
	}
	return &EntityStateChangedData{
		NotificationTemplateData:       data,
		EntityStateChangedNotification: &nData,
	}, nil
}

// EntityStateChangedData is the data for the entity_state_changed email.
type EntityStateChangedData struct {
	email.NotificationTemplateData
	*ttnpb.EntityStateChangedNotification
}
