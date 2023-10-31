package errors

import "github.com/gotnospirit/messageformat"

var formatter, _ = messageformat.New()

// FormatMessage formats the message using the given attributes.
func (d *Definition) FormatMessage(attributes map[string]any) string {
	if d == nil {
		return ""
	}
	if len(attributes) == 0 {
		return d.messageFormat
	}
	parsedMessageFormat := d.parsedMessageFormat
	if parsedMessageFormat == nil {
		parsedMessageFormat, _ = formatter.Parse(d.messageFormat)
	}
	if parsedMessageFormat != nil {
		formatted, err := parsedMessageFormat.FormatMap(attributes)
		if err == nil {
			return formatted
		}
	}
	return d.messageFormat
}
