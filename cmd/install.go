package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the cli in PATH",
	Run: func(cmd *cobra.Command, args []string) {
		actions.InstallEnvrc()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
