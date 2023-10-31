package ttnmage

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

// Styl namespace.
type Styl mg.Namespace

// Lint runs stylint over frontend styl files.
func (styl Styl) Lint() error {
	if mg.Verbose() {
		fmt.Println("Running stylint")
	}
	return runYarnV("stylint",
		"--config", "config/stylintrc.json",
		"./pkg/webui",
	)
}
