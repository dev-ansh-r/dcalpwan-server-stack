package smtp

import (
	"net"
	"testing"

	"github.com/emersion/go-smtp"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestSMTP(t *testing.T) {
	t.Parallel()

	a := assertions.New(t)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()

	bkd := &backend{
		messages: make(chan *message, 1),
	}
	server := smtp.NewServer(bkd)

	go server.Serve(lis) // nolint:errcheck

	smtpAddress := lis.Addr().String()

	ctx := test.Context()
	ctx = log.NewContext(ctx, test.GetLogger(t))

	smtp, err := New(
		ctx,
		email.Config{
			SenderName:    "Unit Test",
			SenderAddress: "unit@test.local",
		},
		Config{
			Address: smtpAddress,
		},
	)
	a.So(err, should.BeNil)

	mail := &email.Message{
		TemplateName:     "test",
		RecipientName:    "John Doe",
		RecipientAddress: "john.doe@example.com",
		Subject:          "Testing SMTP",
		HTMLBody:         "<h1>Testing SMTP</h1><p>We are testing SMTP</p>",
		TextBody:         "****************\r\nTesting SMTP\r\n****************\r\n\r\nWe are testing SMTP",
	}

	err = smtp.Send(mail)
	a.So(err, should.BeNil)

	received := <-bkd.messages

	a.So(received.Sender, should.Equal, "unit@test.local")
	a.So(received.Recipients, should.Contain, mail.RecipientAddress)

	dataString := string(received.Data)

	a.So(dataString, should.ContainSubstring, mail.Subject)
	a.So(dataString, should.ContainSubstring, mail.HTMLBody)
	a.So(dataString, should.ContainSubstring, mail.TextBody)
}
