package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
)

func print(k, v string) {
	fmt.Printf("%-20s %s\n", k+":", v)
}

// Print version information.
func Print(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s: %s\n", root.Short, root.Name())
			print("Version", version.TTN)
			if version.BuildDate != "" {
				print("Build date", version.BuildDate)
			}
			if version.GitCommit != "" {
				print("Git commit", version.GitCommit)
			}
			print("Go version", runtime.Version())
			print("OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
		},
	}
}
