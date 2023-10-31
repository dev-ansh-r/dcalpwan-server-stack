package sendgrid

import (
	"os"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestSendGrid(t *testing.T) {
	a := assertions.New(t)

	apiKey := os.Getenv("SENDGRID_API_KEY")

	sg, err := New(
		log.NewContext(test.Context(), test.GetLogger(t)),
		email.Config{
			SenderName:    "Unit Test",
			SenderAddress: "unit@test.local",
		},
		Config{
			APIKey:      apiKey,
			SandboxMode: true,
		},
	)
	a.So(err, should.BeNil)

	err = sg.Send(&email.Message{
		TemplateName:     "test",
		RecipientName:    "John Doe",
		RecipientAddress: "john.doe@example.com",
		Subject:          "Testing SendGrid",
		HTMLBody:         "<h1>Testing SendGrid</h1><p>We are testing SendGrid</p>",
		TextBody:         "****************\nTesting SendGrid\n****************\n\nWe are testing SendGrid",
	})

	if apiKey == "" {
		a.So(err, should.NotBeNil)
	} else {
		a.So(err, should.BeNil)
	}
}
