package dir

import (
	"os"
	"path/filepath"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestMailDir(t *testing.T) {
	a, ctx := test.New(t)
	tempDir := t.TempDir()

	mailer, err := New(ctx, email.Config{
		SenderName:    "Sender Name",
		SenderAddress: "sender@example.com",
		Provider:      "dir",
		Network:       email.NetworkConfig{
			// Not used
		},
	}, tempDir)
	a.So(err, should.BeNil)

	err = mailer.Send(&email.Message{
		TemplateName:     "irrelevant",
		RecipientName:    "John Doe",
		RecipientAddress: "john.doe@example.com",
		Subject:          "Email Subject",
		HTMLBody:         "<h1>Title</h1><p>Body</p>",
		TextBody:         "Title\n-----\n\nBody",
	})
	a.So(err, should.BeNil)

	entries, err := os.ReadDir(tempDir)
	if a.So(err, should.BeNil) && a.So(entries, should.HaveLength, 1) {
		data, err := os.ReadFile(filepath.Join(tempDir, entries[0].Name()))
		a.So(err, should.BeNil)
		a.So(string(data), should.ContainSubstring, "<h1>Title</h1><p>Body</p>")
		a.So(string(data), should.ContainSubstring, "Title\n-----\n\nBody")
	}
}
