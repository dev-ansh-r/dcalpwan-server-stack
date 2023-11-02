package dir

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/oklog/ulid/v2"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	gomail "gopkg.in/mail.v2"
)

// MailDir is an email.Sender implementation that writes emails to a directory.
type MailDir struct {
	emailConfig     email.Config
	dir             string
	messageSettings []gomail.MessageSetting
}

// New returns a new email.Sender that writes emails to dir.
func New(_ context.Context, emailConfig email.Config, dir string) (*MailDir, error) {
	info, err := os.Stat(dir)
	switch {
	case errors.Is(err, os.ErrNotExist):
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	case !info.IsDir():
		return nil, fmt.Errorf("%q is not a dir", dir)
	}
	return &MailDir{
		emailConfig: emailConfig,
		dir:         dir,
		messageSettings: []gomail.MessageSetting{
			gomail.SetEncoding(gomail.Unencoded),
		},
	}, nil
}

func (md *MailDir) ulid() ulid.ULID {
	return ulid.MustNew(ulid.Now(), rand.Reader)
}

// Send implements email.Sender.
func (md *MailDir) Send(message *email.Message) (err error) {
	m := gomail.NewMessage(md.messageSettings...)
	m.SetAddressHeader("From", md.emailConfig.SenderAddress, md.emailConfig.SenderName)
	m.SetAddressHeader("To", message.RecipientAddress, message.RecipientName)
	m.SetHeader("Subject", message.Subject)
	if message.TextBody != "" {
		m.AddAlternative("text/plain", message.TextBody)
	}
	if message.HTMLBody != "" {
		m.AddAlternative("text/html", message.HTMLBody)
	}
	f, err := os.OpenFile(
		filepath.Join(md.dir, fmt.Sprintf("%s.eml", md.ulid().String())),
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0o644,
	)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	_, err = m.WriteTo(f)
	return err
}
