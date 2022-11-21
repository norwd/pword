package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/norwd/pword/internal/dat"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func getSimpleCmdFlags() []string {
	flags := make([]string, 0)

	simpleCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		flags = append(flags, flag.Name)
	})

	return flags
}

func TestSimpleCmdDefinesFlagsForAllCharacterClasses(t *testing.T) {
	for name := range dat.CharacterClasses {
		if simpleCmd.Flags().Lookup("include-"+name) == nil {
			t.Errorf("expected --include-%s flag, but none found", name)
		}
		if simpleCmd.Flags().Lookup("no-include-"+name) == nil {
			t.Errorf("expected --no-include-%s flag, but none found", name)
		}
	}
}

func TestSimpleCmdOnlyDefinesFlagsForValidCharacterClasses(t *testing.T) {
	for _, flag := range getSimpleCmdFlags() {
		if !strings.Contains(flag, "include-") {
			continue // Not a character inclusion/exclusion flag
		}

		parts := strings.Split(flag, "-")
		last := parts[len(parts)-1]

		if last == "custom" {
			continue // Custom character set is not in predefined list
		}

		if _, ok := dat.CharacterClasses[last]; !ok {
			t.Errorf("unexpected flag %q, no matching character class", flag)
		}
	}
}

func TestRunSimpleCmd(t *testing.T) {
	tests := []struct {
		name string
		len  int
		err  func(*testing.T, error)
	}{
		{name: "Zero Length", len: 0, err: nil},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			cmd := &cobra.Command{}

			cmd.Flags().IntP("length", "l", test.len, "")
			cmd.Flags().Set("length", fmt.Sprintf("%d", test.len))

			err := runSimpleCmd(cmd, []string{})

			if err == nil && test.err != nil {
				t.Error("want error, got nil")
			} else if err != nil && test.err == nil {
				t.Errorf("want nil error, got %q", err)
			} else if err != nil && test.err != nil {
				test.err(t, err)
			}
		})
	}
}
