package lorawan

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	errDecode                       = errors.Define("decode", "decode `{lorawan_field}`")
	errEncode                       = errors.Define("encode", "encode `{lorawan_field}`")
	errEncodedFieldLengthBound      = errors.DefineInvalidArgument("encoded_field_length_bound", "`{lorawan_field}` encoded length `{value}` should between `{min}` and `{max}`", valueKey)
	errEncodedFieldLengthEqual      = errors.DefineInvalidArgument("encoded_field_length_equal", "`{lorawan_field}` encoded length `{value}` should be `{expected}`", valueKey)
	errEncodedFieldLengthTwoChoices = errors.DefineInvalidArgument("encoded_field_length_multiple_choices", "`{lorawan_field}` encoded length `{value}` should be equal to `{expected_1}` or `{expected_2}`", valueKey)
	errFieldBound                   = errors.DefineInvalidArgument("field_bound", "`{lorawan_field}` should be between `{min}` and `{max}`", valueKey)
	errFieldHasMax                  = errors.DefineInvalidArgument("field_with_max", "`{lorawan_field}` value `{value}` should be lower or equal to `{max}`", valueKey)
	errFieldLengthEqual             = errors.DefineInvalidArgument("field_length_equal", "`{lorawan_field}` length `{value}` should be equal to `{expected}`", valueKey)
	errFieldLengthHasMax            = errors.DefineInvalidArgument("field_length_with_max", "`{lorawan_field}` length `{value}` should be lower or equal to `{max}`", valueKey)
	errFieldLengthHasMin            = errors.DefineInvalidArgument("field_length_with_min", "`{lorawan_field}` length `{value}` should be higher or equal to `{min}`", valueKey)
	errFieldLengthTwoChoices        = errors.DefineInvalidArgument("field_length_with_multiple_choices", "`{lorawan_field}` length `{value}` should be equal to `{expected_1}` or `{expected_2}`", valueKey)
	errUnknownField                 = errors.DefineInvalidArgument("unknown_field", "unknown `{lorawan_field}`", valueKey)
	errProprietary                  = errors.DefineInvalidArgument("proprietary", "expected standard LoRaWAN frame")
)

const valueKey = "value"

type valueErr func(any) *errors.Error

func unexpectedValue(err interface {
	WithAttributes(kv ...any) *errors.Error
},
) valueErr {
	return func(value any) *errors.Error {
		return err.WithAttributes(valueKey, value)
	}
}

func errExpectedLowerOrEqual(lorawanField string, max any) valueErr {
	return unexpectedValue(errFieldHasMax.WithAttributes("lorawan_field", lorawanField, "max", max))
}

func errExpectedLengthEqual(lorawanField string, expected any) valueErr {
	return unexpectedValue(errFieldLengthEqual.WithAttributes("lorawan_field", lorawanField, "expected", expected))
}

func errExpectedLengthLowerOrEqual(lorawanField string, max any) valueErr {
	return unexpectedValue(errFieldLengthHasMax.WithAttributes("lorawan_field", lorawanField, "max", max))
}

func errExpectedLengthHigherOrEqual(lorawanField string, min any) valueErr {
	return unexpectedValue(errFieldLengthHasMin.WithAttributes("lorawan_field", lorawanField, "min", min))
}

func errExpectedLengthTwoChoices(lorawanField string, expected1, expected2 any) valueErr {
	return unexpectedValue(errFieldLengthTwoChoices.WithAttributes("lorawan_field", lorawanField, "expected_1", expected1, "expected_2", expected2))
}

func errExpectedLengthEncodedBound(lorawanField string, min, max any) valueErr {
	return unexpectedValue(errEncodedFieldLengthBound.WithAttributes("lorawan_field", lorawanField, "min", min, "max", max))
}

func errExpectedLengthEncodedEqual(lorawanField string, expected any) valueErr {
	return unexpectedValue(errEncodedFieldLengthEqual.WithAttributes("lorawan_field", lorawanField, "expected", expected))
}

func errExpectedLengthEncodedTwoChoices(lorawanField string, expected1, expected2 any) valueErr {
	return unexpectedValue(errEncodedFieldLengthTwoChoices.WithAttributes("lorawan_field", lorawanField, "expected_1", expected1, "expected_2", expected2))
}

func errFailedEncoding(lorawanField string) *errors.Error {
	return errEncode.WithAttributes("lorawan_field", lorawanField)
}

func errFailedDecoding(lorawanField string) *errors.Error {
	return errDecode.WithAttributes("lorawan_field", lorawanField)
}

func errUnknown(lorawanField string) valueErr {
	return unexpectedValue(errUnknownField.WithAttributes("lorawan_field", lorawanField))
}

func errMissing(lorawanField string) *errors.Error {
	return errUnknownField.WithAttributes("lorawan_field", lorawanField)
}

func errExpectedBetween(lorawanField string, min, max any) valueErr {
	return unexpectedValue(errFieldBound.WithAttributes("lorawan_field", lorawanField, "min", min, "max", max))
}
