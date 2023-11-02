package ttnpb

import "fmt"

// FieldIsZero returns whether path p is zero.
func (v *Picture_Embedded) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "data":
		return v.Data == nil
	case "mime_type":
		return v.MimeType == ""
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *Picture) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "embedded":
		return v.Embedded == nil
	case "embedded.data":
		return v.Embedded.FieldIsZero("embedded.data")
	case "embedded.mime_type":
		return v.Embedded.FieldIsZero("embedded.mime_type")
	case "sizes":
		return v.Sizes == nil
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}
