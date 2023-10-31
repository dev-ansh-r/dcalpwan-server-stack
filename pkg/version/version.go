// Package version contains version and build variables set by the CI process.
package version

import (
	"fmt"
)

// String returns the version string.
func String() string {
	version := TTN
	if GitCommit != "" && BuildDate != "" {
		version += fmt.Sprintf(" (%s, %s)", GitCommit, BuildDate)
	}
	return version
}
