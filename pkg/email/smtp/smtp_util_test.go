package smtp

import (
	"io"

	"github.com/emersion/go-smtp"
)

type backend struct {
	messages chan *message
}

func (bkd *backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	s := &session{
		msgs: bkd.messages,
	}
	return s, nil
}

type message struct {
	Sender     string
	Recipients []string
	Data       []byte
	Opts       *smtp.MailOptions
}

type session struct {
	msg  *message
	msgs chan *message
}

func (*session) AuthPlain(_, _ string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	s.Reset()
	s.msg.Sender = from
	s.msg.Opts = opts
	return nil
}

func (s *session) Rcpt(to string, _ *smtp.RcptOptions) error {
	s.msg.Recipients = append(s.msg.Recipients, to)
	return nil
}

func (s *session) Data(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	s.msg.Data = b
	s.msgs <- s.msg
	return nil
}

func (s *session) Reset() {
	s.msg = &message{}
}

func (*session) Logout() error {
	return nil
}
