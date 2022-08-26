package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path"
)

const version = "v0.0.0"

func usage(self string) {
	fmt.Fprintln(os.Stderr, `
usage: ` + self + ` [-v|--version] [-h|--help] [simple]

options:

	-v | --version

		Show the version and exit.

	-h | --help

		Show this message and exit.

backends:

	simple

		Generates a simple, 16 character password.

` + version)
}

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
			usage(self)
			os.Exit(0)
		case "simple":
			var buf [16]byte
			if _, err := rand.Read(buf[:]); err != nil {
				fmt.Fprintln(os.Stderr, self+": "+err.Error())
				os.Exit(1)
			}

			fmt.Println(base64.RawStdEncoding.EncodeToString(buf[:]))
		default:
			fmt.Fprintln(os.Stderr, self + ": unknown flag or password type '" + arg + "'")
			usage(self)
			os.Exit(1)
		}
	}
}

