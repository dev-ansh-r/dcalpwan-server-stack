package errors

// CorrelationID of the error.
func (e *Error) CorrelationID() string {
	if e == nil {
		return ""
	}
	return e.correlationID
}

// CorrelationID is not present in the error definition, so this just returns an empty string.
func (*Definition) CorrelationID() string { return "" }
