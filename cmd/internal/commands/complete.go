package commands

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var errInvalidShell = errors.DefineInvalidArgument("invalid_shell", "invalid shell")

// Complete returns the auto-complete command
func Complete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "complete",
		Aliases: []string{"completion"},
		Hidden:  true,
		Short:   "Generate script for auto-completion",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name, _ := cmd.Flags().GetString("executable"); name != "" {
				cmd.Root().Use = name
			}
			switch shell, _ := cmd.Flags().GetString("shell"); shell {
			case "bash":
				return cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				return cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				// Fish does not accept `-` in variable names
				buf := new(bytes.Buffer)
				if err := cmd.Root().GenFishCompletion(buf, true); err != nil {
					return err
				}
				script := strings.Replace(buf.String(), "__ttn-lw-", "__ttn_lw_", -1)
				_, err := fmt.Print(script)
				return err
			case "powershell":
				return cmd.Root().GenPowerShellCompletion(os.Stdout)
			default:
				return errInvalidShell.WithAttributes("shell", shell)
			}
		},
	}
	cmd.Flags().String("shell", "bash", "bash|zsh|fish|powershell")
	cmd.Flags().String("executable", "", "Executable name to create generate auto completion script for")
	return cmd
}
