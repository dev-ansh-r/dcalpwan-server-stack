package goproto

import (
	"fmt"
	"math"

	"google.golang.org/protobuf/types/known/structpb"
)

func validateNumber(n float64, prefix string) []string {
	if math.IsNaN(n) {
		return []string{fmt.Sprintf("%s: invalid NaN value", prefix)}
	}
	if math.IsInf(n, 1) {
		return []string{fmt.Sprintf("%s: invalid Infinity value", prefix)}
	}
	if math.IsInf(n, -1) {
		return []string{fmt.Sprintf("%s: invalid -Infinity value", prefix)}
	}
	return nil
}

func validateList(l *structpb.ListValue, prefix string) []string {
	if l == nil {
		return nil
	}
	total := make([]string, 0)
	for i, v := range l.Values {
		prefix := fmt.Sprintf("%s[%d]", prefix, i)
		switch vv := v.GetKind().(type) {
		case *structpb.Value_NumberValue:
			if vv == nil {
				break
			}
			total = append(total, validateNumber(vv.NumberValue, prefix)...)
		case *structpb.Value_StructValue:
			if vv == nil {
				break
			}
			total = append(total, validateStruct(vv.StructValue, prefix)...)
		case *structpb.Value_ListValue:
			if vv == nil {
				break
			}
			total = append(total, validateList(vv.ListValue, prefix)...)
		}
	}
	return total
}

func validateStruct(st *structpb.Struct, prefix string) []string {
	if st == nil {
		return nil
	}
	total := make([]string, 0)
	for k, v := range st.Fields {
		prefix := fmt.Sprintf("%s.%s", prefix, k)
		switch vv := v.GetKind().(type) {
		case *structpb.Value_NumberValue:
			if vv == nil {
				break
			}
			total = append(total, validateNumber(vv.NumberValue, prefix)...)
		case *structpb.Value_StructValue:
			if vv == nil {
				break
			}
			total = append(total, validateStruct(vv.StructValue, prefix)...)
		case *structpb.Value_ListValue:
			if vv == nil {
				break
			}
			total = append(total, validateList(vv.ListValue, prefix)...)
		}
	}
	return total
}

// ValidateStruct recursively verifies if the struct contains any invalid values (NaN, -Infinity, Infinity)
// and emits warning messages for such fields.
func ValidateStruct(st *structpb.Struct) []string {
	return validateStruct(st, "")
}
