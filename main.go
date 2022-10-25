package main

import (
	"fmt"
	"github.com/norwd/pword/internal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
