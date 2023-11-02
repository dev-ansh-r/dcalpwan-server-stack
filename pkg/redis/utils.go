package redis

import (
	"regexp"
	"strings"
)

// EntityRegex returns a regex that must match a given redis key format.
func EntityRegex(key string) (*regexp.Regexp, error) {
	keyRegex := strings.ReplaceAll(key, ":", "\\:")
	keyRegex = strings.ReplaceAll(keyRegex, "*", ".[^\\:]*")
	keyRegex = keyRegex + "$"
	return regexp.Compile(keyRegex)
}
