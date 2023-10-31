package telemetry

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
)

// GenerateHash receives a string slice and returns a string of a UID.
func GenerateHash(_ context.Context, elems ...string) string {
	s := sha512.New()
	for _, elem := range elems {
		s.Write([]byte(elem + ":"))
	}
	return base64.StdEncoding.EncodeToString(s.Sum(nil))
}
