// Package email provides an interface to send messages over email.
package email

// Message for sending over email.
type Message struct {
	TemplateName string

	RecipientName    string
	RecipientAddress string

	Subject  string
	HTMLBody string
	TextBody string
}

// Sender is the interface for sending messages over email.
type Sender interface {
	Send(message *Message) error
}
