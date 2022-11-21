package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	simpleCmd.Flags().IntP("length", "l", 16, "set how long the password should be.")
	simpleCmd.Flags().String("include-custom", "", "include this character set.")
	simpleCmd.Flags().String("no-include-custom", "", "exclude this character set.")

	for _, class := range []struct{
		name string
		desc string
		warn string
	}{
		{ name: "xdigit", desc: "hexadecimal digits, upper and lower case 'A' through 'F' and digits '0' through '9'.", },
		{ name: "upper", desc: "upper case letters 'A' through 'Z'.", },
		{ name: "space", desc: "white space characters, empty space (' '), horizontal tab ('\\t'), carriage return ('\\r'), line feed ('\\n'), vertical tab ('\\v'), and form feed ('\\f').", warn: "you probably meant to use --[no--]include--blank.", },
		{ name: "punct", desc: "punctuation characters, such as various brackets '[]()<>{}', mathematical symbols '+-/*', general punctuation '.,?!', and others.", },
		{ name: "print", desc: "all printable characters, including the space character.", },
		{ name: "lower", desc: "lower case letters 'a' through 'z'.", },
		{ name: "graph", desc: "all printable characters, except the space character.", },
		{ name: "digit", desc: "digits '0' through '0'.", },
		{ name: "cntrl", desc: "control characters, such as '\\a' or '\\e'.", warn: "these characters may not be displayed properly and may not even be supported by the application the password is for.", },
		{ name: "blank", desc: "the space and tab characters ' ' and '\\t'.", },
		{ name: "alpha", desc: "both upper and lower case letters 'A' through 'Z'.", },
		{ name: "alnum", desc: "both upper and lower case letters 'A' through 'Z', and digits '0' through '9'.", },
	} {
		simpleCmd.Flags().Bool("include-" + class.name, false, "include " + class.desc)
		simpleCmd.Flags().Bool("no-include-" + class.name, false, "exclude " + class.desc)

		if class.warn != "" {
			simpleCmd.Flags().MarkDeprecated("include-" + class.name, class.warn)
		}
	}

	rootCmd.AddCommand(simpleCmd)
}

var simpleCmd = &cobra.Command{
	Use:   "simple",
	Short: "Generates a simple password.",
	Long:  "Generate a simple base64 encoded password from random data.",
	RunE:  runSimpleCmd,
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
