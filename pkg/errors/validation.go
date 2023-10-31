package errors

type validationError interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

var errValidation = DefineInvalidArgument("validation", "invalid `{field}`: {reason}", "name")
