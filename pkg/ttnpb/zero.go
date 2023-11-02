package ttnpb

func fieldsAreZero(v interface{ FieldIsZero(string) bool }, paths ...string) bool {
	for _, p := range paths {
		if !v.FieldIsZero(p) {
			return false
		}
	}
	return true
}
