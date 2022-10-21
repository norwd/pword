package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const version = "v0.0.0"

var rootCmd = &cobra.Command{
	Use:     "pword",
	Short:   "A small, command line password generator.",
	Long: "Generate passwords for any occasion, using a variety of backends.",
	Version: version,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
