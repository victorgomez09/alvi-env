package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var runCommand = &cobra.Command{
	Use:     "run",
	Aliases: []string{"start"},
	Short:   "Run environment variables for an environment",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// actions.RunEnvironment(args[0])
		if err := actions.ExecuteEnvrc(); err != nil {
			fmt.Fprintf(os.Stderr, "Error [alenv]: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCommand)
}
