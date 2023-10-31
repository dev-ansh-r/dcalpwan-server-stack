// Package mock provides a test email provider that is used in tests.
package mock

import "go.thethings.network/lorawan-stack/v3/pkg/email"

// Mock implements the email.Sender interface and stores sent emails internally.
type Mock struct {
	Messages []*email.Message
	Error    error
}

// New returns a new mock email.Sender.
func New() *Mock {
	return &Mock{}
}

// Send implements email.Sender.
// It appends the messages to Messages field and returns the Error field.
func (m *Mock) Send(message *email.Message) error {
	m.Messages = append(m.Messages, message)
	return m.Error
}
