package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alenv",
	Short: "alenv is a cli tool for managing environment variables",
	Long:  "alenv is a cli tool for managing environment variables isolated by environments like 'development', 'testing', 'production', etc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alenv is a cli tool for managing environment variables. Use with --help flag to see commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf(os.Stderr.Name(), "Oops. An error while executing Alenv '%s'\n", err)
		os.Exit(1)
	}
}
