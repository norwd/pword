package main

import (
	"strconv"
	"fmt"
	"os"
	"path"
)

const version = "v0.0.0"

type passworder interface {
	Password() (password string, err error)
}

func usage(self string) {
	fmt.Fprintln(os.Stderr, `
`+self+` - A small, command-line password generator.

usage: `+self+` [-v|--version] [-h|--help] [BACKEND] [OPTIONS]

global options:

	-v | --version

		Show the version and exit.

	-h | --help

		Show this message and exit.

backends:

	simple - Generates a simple password, just a random string of characters.

	usage: simple [-l|--length <LENGTH>]

	options:

		-l | --length <LENGTH>

			Sets the length of the password. Defaults to 16.

`+version)
}

func main() {
	self := path.Base(os.Args[0])
	args := os.Args[1:]

	if len(args) < 1 {
		args = append(args, "simple")
	}

	var generator passworder

	for i := 0; i < len(args); i++ {
		switch arg := args[i]; arg {
		case "-v", "--version":
			fmt.Fprintln(os.Stderr, version)
			os.Exit(0)
		case "-h", "--help":
			usage(self)
			os.Exit(0)
		case "simple":
			s := simple{}

			for i++; i < len(args); i++ {
				switch arg := args[i]; arg {
				case "-l", "--length":
					if i++; i >= len(args) {
						fmt.Fprintln(os.Stderr, self+": argument '"+arg+"' requires parameter")
						usage(self)
						os.Exit(1)
					} else if length, err := strconv.Atoi(args[i]); err != nil {
						fmt.Fprintln(os.Stderr, self+": invalid length '"+args[i]+"': "+err.Error())
						usage(self)
						os.Exit(1)
					} else {
						s.length = length
					}
				default:
					fmt.Fprintln(os.Stderr, self+": unknown flag '"+arg+"'")
					usage(self)
					os.Exit(1)
				}
			}

			generator = s
		default:
			fmt.Fprintln(os.Stderr, self+": unknown flag or password type '"+arg+"'")
			usage(self)
			os.Exit(1)
		}

		if password, err := generator.Password(); err != nil {
			fmt.Fprintln(os.Stderr, self+": "+err.Error())
			os.Exit(1)
		} else {
			fmt.Println(password)
		}
	}
}
