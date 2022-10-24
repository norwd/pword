package cmd

import (
	"testing"
	"sync"
	"github.com/spf13/cobra"
)

var mu sync.Mutex

func TestExecuteCallsRootCmd(t *testing.T) {
	mu.Lock()

	defer func(cmd *cobra.Command) {
		rootCmd = cmd
		mu.Unlock()
	}(rootCmd)

	var count int

	rootCmd = &cobra.Command{
		RunE: func (cmd *cobra.Command, args []string) error {
			count++
			return nil
		},
	}

	Execute()

	if have, want := count, 1; have != want {
		t.Errorf("rootCmd.RunE() call count: have %d, want %d", have, want)
	}
}

