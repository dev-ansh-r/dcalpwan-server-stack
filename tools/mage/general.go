package ttnmage

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

var initDeps []any

// Init initializes the tooling.
func Init() {
	mg.Deps(initDeps...)
}

func targetError(err error) error {
	if err != nil {
		return fmt.Errorf("failed checking modtime: %w", err)
	}
	return nil
}
