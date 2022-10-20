package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	simpleCmd.Flags().IntP("length", "l", 16, "How long the password should be.")
	rootCmd.AddCommand(simpleCmd)
}

var simpleCmd = &cobra.Command{
	Use: "simple",
	Short: "Generates a simple password.",
	Long: "Generate a simple base64 encoded password from random data.",
	RunE: runSimpleCmd,
}

func runSimpleCmd(cmd *cobra.Command, args []string) (err error) {
	var length int

	if length, err = cmd.Flags().GetInt("length"); err != nil {
		return
	}

	buf := make([]byte, length)

	if _, err = rand.Read(buf[:]); err != nil {
		return
	}

	fmt.Println(base64.RawStdEncoding.EncodeToString(buf[:])[:length])

	return
}
