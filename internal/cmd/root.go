package cmd

import (
	"github.com/spf13/cobra"
)

const version = "v0.2.0"

var rootCmd = &cobra.Command{
	Use:     "pword",
	Short:   "A small, command line password generator.",
	Long:    "Generate passwords for any occasion, using a variety of backends.",
	Version: version,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() error {
	return rootCmd.Execute()
}
