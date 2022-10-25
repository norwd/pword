package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
	"testing"
)

var mu sync.Mutex

func TestExecuteCallsRootCmd(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "Without Error", err: nil},
		{name: "With Error", err: fmt.Errorf("expected error")},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			mu.Lock()

			defer func(cmd *cobra.Command) {
				rootCmd = cmd
				mu.Unlock()
			}(rootCmd)

			var count int

			rootCmd = &cobra.Command{
				SilenceErrors: true,
				SilenceUsage:  true,
				RunE: func(cmd *cobra.Command, args []string) error {
					count++
					return test.err
				},
			}

			err := Execute()

			if have, want := count, 1; have != want {
				t.Errorf("call count: have %d, want %d", have, want)
			}

			if have, want := err, test.err; have != want {
				t.Errorf("error: have %q, want %q", have, want)
			}
		})
	}
}
