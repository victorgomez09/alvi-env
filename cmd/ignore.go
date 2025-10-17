package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var ignoreCmd = &cobra.Command{
	Use:     "ignore",
	Aliases: []string{"ignore", "pass", "skip"},
	Short:   "Ignore current directory from read .envrc file",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		actions.IgnoreEnvrc()
	},
}

func init() {
	rootCmd.AddCommand(ignoreCmd)
}
