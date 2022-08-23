package myctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "myctl",
	Version: version,
	Short:   "myctl is ",
	Long: `myctl is a super fancy CLI
				hello
	`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "whoops. thete was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
