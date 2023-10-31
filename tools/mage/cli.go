package ttnmage

import (
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
)

// Cli namespace.
type Cli mg.Namespace

// Autocomplete generates scripts for auto-completion.
func (Cli) Autocomplete() error {
	for _, cfg := range []struct {
		which      string
		shell      string
		filename   string
		executable string
	}{
		{
			"ttn-lw-cli", "bash", "ttn-lw-cli-snap", "ttn-lw-stack.ttn-lw-cli",
		},
		{
			"ttn-lw-cli", "bash", "ttn-lw-cli", "ttn-lw-cli",
		},
		{
			"ttn-lw-cli", "fish", "ttn-lw-cli.fish", "ttn-lw-cli",
		},
		{
			"ttn-lw-cli", "zsh", "_ttn-lw-cli", "ttn-lw-cli",
		},
		{
			"ttn-lw-stack", "bash", "ttn-lw-stack", "ttn-lw-stack",
		},
		{
			"ttn-lw-stack", "fish", "ttn-lw-stack.fish", "ttn-lw-stack",
		},
		{
			"ttn-lw-stack", "zsh", "_ttn-lw-stack", "ttn-lw-stack",
		},
	} {
		f, err := os.Create(path.Join("config", "completion", cfg.shell, cfg.filename))
		if err != nil {
			return err
		}
		if err = execGo(
			f, os.Stderr,
			"run", fmt.Sprintf("./cmd/%s", cfg.which), "complete",
			"--shell", cfg.shell,
			"--executable", cfg.executable); err != nil {
			return err
		}
	}
	return nil
}
