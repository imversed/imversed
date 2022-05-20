package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// New creates a new root command for `Imversed develop tools` with its sub commands.
func New(ctx context.Context) *cobra.Command {
	cobra.EnableCommandSorting = false

	c := &cobra.Command{
		Use:   "go run devtools/main.go",
		Short: "devtools make sense",
		// Long: ``,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.AddCommand(NewJSUpdate(ctx))

	return c
}
