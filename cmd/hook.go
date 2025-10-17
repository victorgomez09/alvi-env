package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var hookCommand = &cobra.Command{
	Use:     "hook",
	Aliases: []string{"hook"},
	Short:   "Prints the shell function needed for integration",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		actions.PrintShellHook()
	},
}

func init() {
	rootCmd.AddCommand(hookCommand)
}
