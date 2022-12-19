package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pword",
	Short:   "A small, command line password generator.",
	Long:    "Generate passwords for any occasion, using a variety of backends.",
	Version: Version,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

// Execute the command
func Execute() error {
	return rootCmd.Execute()
}
