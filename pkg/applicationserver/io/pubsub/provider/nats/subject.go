package nats

import (
	"fmt"
	"strings"
)

func combineSubjects(s1, s2 string) string {
	s1 = strings.Trim(s1, ".")
	s2 = strings.Trim(s2, ".")
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	return fmt.Sprintf("%s.%s", s1, s2)
}
