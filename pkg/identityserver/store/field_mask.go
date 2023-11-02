package store

import "strings"

// FieldMask is used to specify applicable fields in SELECT or UPDATE queries.
type FieldMask []string

// Contains returns true if the given field is present in the field mask.
func (fm FieldMask) Contains(search string) bool {
	for _, f := range fm {
		if f == search {
			return true
		}
	}
	return false
}

// TopLevel returns the top-level fields from the field mask.
func (fm FieldMask) TopLevel() FieldMask {
	out := make(FieldMask, 0, len(fm))
	for _, f := range fm {
		before, _, _ := strings.Cut(f, ".")
		if !out.Contains(before) {
			out = append(out, before)
		}
	}
	return out
}

// TrimPrefix returns a field mask with all fields of fm that contain prefix, but then having that prefix removed.
func (fm FieldMask) TrimPrefix(prefix string) FieldMask {
	if !strings.HasSuffix(prefix, ".") {
		prefix += "."
	}
	out := make([]string, 0, len(fm))
	for _, f := range fm {
		if strings.HasPrefix(f, prefix) {
			out = append(out, strings.TrimPrefix(f, prefix))
		}
	}
	return out
}
