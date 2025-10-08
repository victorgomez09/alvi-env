package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/victorgomez09/alvi-env/internal/actions"
)

var initCommand = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialize"},
	Short:   "Initialize alenv in current directory",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Creating Alenv configuration file")
		actions.InitializeConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
