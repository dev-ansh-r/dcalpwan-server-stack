package smtp

import (
	"context"
	"net"
	"strconv"

	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	gomail "gopkg.in/mail.v2"
)

type sendTask struct {
	message *gomail.Message
	result  chan error
}

// SMTP is the type that implements SMTP as email provider.
type SMTP struct {
	ctx         context.Context
	logger      log.Interface
	emailConfig email.Config
	smtpConfig  Config
	dialer      *gomail.Dialer
	tasks       chan sendTask
}

func (s *SMTP) handle() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case task := <-s.tasks:
			task.result <- s.dialer.DialAndSend(task.message)
		}
	}
}

var buffer = 8 // send buffer per connection

// New creates a SMTP email provider.
func New(ctx context.Context, emailConfig email.Config, smtpConfig Config) (email.Sender, error) {
	host, portStr, _ := net.SplitHostPort(smtpConfig.Address)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	if smtpConfig.Connections == 0 {
		smtpConfig.Connections = 1
	}
	s := &SMTP{
		ctx:         ctx,
		logger:      log.FromContext(ctx).WithField("email_provider", "SMTP"),
		emailConfig: emailConfig,
		smtpConfig:  smtpConfig,
		dialer:      gomail.NewDialer(host, port, smtpConfig.Username, smtpConfig.Password),
		tasks:       make(chan sendTask, smtpConfig.Connections*buffer),
	}
	if smtpConfig.TLSConfig != nil {
		s.dialer.TLSConfig = smtpConfig.TLSConfig.Clone()
		s.dialer.TLSConfig.ServerName, _, _ = net.SplitHostPort(s.smtpConfig.Address)
	}
	for i := 0; i < smtpConfig.Connections; i++ {
		go s.handle()
	}
	return s, nil
}

// Send an email message.
func (s *SMTP) Send(message *email.Message) error {
	logger := s.logger.WithFields(log.Fields(
		"template_name", message.TemplateName,
		"recipient_name", message.RecipientName,
		"recipient_address", message.RecipientAddress,
	))

	m := gomail.NewMessage()
	m.SetAddressHeader("From", s.emailConfig.SenderAddress, s.emailConfig.SenderName)
	m.SetAddressHeader("To", message.RecipientAddress, message.RecipientName)
	m.SetHeader("Subject", message.Subject)
	if message.TextBody != "" {
		m.AddAlternative("text/plain", message.TextBody)
	}
	if message.HTMLBody != "" {
		m.AddAlternative("text/html", message.HTMLBody)
	}

	sendResult := make(chan error)
	logger.Debug("Sending email...")
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	case s.tasks <- sendTask{message: m, result: sendResult}:
		err := <-sendResult
		if err != nil {
			logger.WithError(err).Error("Could not send email")
			return err
		}
		logger.Info("Sent email")
		return nil
	}
}
