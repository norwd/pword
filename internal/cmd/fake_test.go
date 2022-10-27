package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestThatAlwaysFails(t *testing.T) {
	t.Fail("Argghhhh")
}
