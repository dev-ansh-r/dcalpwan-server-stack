package email_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

// welcomeEmailData is the data specifically for the welcome email.
type welcomeEmailData struct {
	email.TemplateData
	ActivationToken string
}

func TestEmail(t *testing.T) {
	t.Parallel()

	a, ctx := test.New(t)

	registry := email.NewTemplateRegistry()

	welcomeEmailTemplate, err := email.NewTemplateFS(
		os.DirFS("testdata"), "welcome",
		email.FSTemplate{
			SubjectTemplate:      "Welcome to {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html",
			HTMLTemplateFile:     "welcome.html",
			TextTemplateBaseFile: "base.txt",
			TextTemplateFile:     "welcome.txt",
			IncludePatterns:      []string{"header.html", "footer.html", "header.txt", "footer.txt"},
		},
	)
	a.So(err, should.BeNil)

	registry.RegisterTemplate(welcomeEmailTemplate)

	a.So(registry.RegisteredTemplates(), should.Contain, "welcome")
	returnedTemplate := registry.GetTemplate(ctx, "welcome")

	for i, template := range []*email.Template{welcomeEmailTemplate, returnedTemplate} {
		template := template
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			a := assertions.New(t)

			message, err := template.Execute(&welcomeEmailData{
				TemplateData: email.NewTemplateData(&email.NetworkConfig{
					Name:              "The Things Network",
					IdentityServerURL: "https://eu1.cloud.thethings.network/oauth",
					ConsoleURL:        "https://console.cloud.thethings.network",
				}, &ttnpb.User{
					Name:                "John Doe",
					PrimaryEmailAddress: "john.doe@example.com",
				}),
			})

			if a.So(err, should.BeNil) && a.So(message, should.NotBeNil) {
				a.So(message.Subject, should.Equal, "Welcome to The Things Network")
				a.So(message.HTMLBody, should.ContainSubstring, `<div class="header">`)
				a.So(message.HTMLBody, should.ContainSubstring, "Welcome to The Things Network, John Doe!")
				a.So(message.HTMLBody, should.ContainSubstring, `<div class="footer">`)
				a.So(message.TextBody, should.ContainSubstring, "==================")
				a.So(message.TextBody, should.ContainSubstring, "Welcome to The Things Network, John Doe!")
			}
		})
	}
}
