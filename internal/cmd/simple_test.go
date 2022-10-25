package cmd

import (
	"testing"
	"fmt"
	"github.com/spf13/cobra"
)

func TestRunSimpleCmd(t *testing.T) {
	tests := []struct{
		name string
		len int
		err func (*testing.T, error)
	}{
		{ name: "Zero Length", len: 0, err: nil },
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func (t *testing.T) {
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

