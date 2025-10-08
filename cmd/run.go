package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var runCommand = &cobra.Command{
	Use:     "run",
	Aliases: []string{"start"},
	Short:   "Run environment variables for an environment",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		actions.RunEnvironment(args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCommand)
}
