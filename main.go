package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path"
)

const version = "v0.0.0"

func main() {
	self := path.Base(os.Args[0])
	args := os.Args[1:]

	if len(args) < 1 {
		args = append(args, "simple")
	}

	for _, arg := range args {
		switch arg {
		case "-v", "--version":
			fmt.Fprintln(os.Stderr, version)
			os.Exit(0)
		case "-h", "--help":
			fmt.Fprintln(os.Stderr, self+" is a simple password generator.")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "usage: "+self+" [-v|--version] [-h--help] [OPTIONS] [TYPE]")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "options:")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "\t-v | --version")
			fmt.Fprintln(os.Stderr, "\t\tShow the version and exit.")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "\t-h | --help")
			fmt.Fprintln(os.Stderr, "\t\tShow this message and exit.")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "types:")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "\tsimple")
			fmt.Fprintln(os.Stderr, "\t\tGenerates a simple password.")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, version)
			os.Exit(0)
		case "simple":
			var buf [16]byte
			if _, err := rand.Read(buf[:]); err != nil {
				fmt.Fprintln(os.Stderr, self+": "+err.Error())
				os.Exit(1)
			}

			password := base64.RawStdEncoding.EncodeToString(buf[:])
			fmt.Println(password)
		default:
			fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: unknown flag or password type '%s'", self, arg))
			os.Exit(1)
		}
	}
}
