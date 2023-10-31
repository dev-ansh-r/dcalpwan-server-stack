package errors

// Resemble returns true iff the given errors resemble, meaning that the
// Namespace and Name of the errors are equal. A nil error only resembles nil.
// Invalid errors or definitions (including typed nil) never resemble anything.
func Resemble(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	ttnA, ok := From(a)
	if !ok {
		return false
	}
	ttnB, ok := From(b)
	if !ok {
		return false
	}
	return ttnA.Namespace() == ttnB.Namespace() &&
		ttnA.Name() == ttnB.Name()
}
