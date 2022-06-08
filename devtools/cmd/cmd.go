package cmd

import (
	"context"
	"path/filepath"

	"github.com/ignite-hq/cli/ignite/services/chain"
	"github.com/spf13/cobra"
)

// New creates a new root command for `Imversed develop tools` with its sub commands.
func New(ctx context.Context) *cobra.Command {
	cobra.EnableCommandSorting = false

	c := &cobra.Command{
		Use:   "devtools [command]",
		Short: "devtools make sense",
		// Long: ``,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.AddCommand(NewJSUpdate(ctx))
	c.AddCommand(NewDiscoverModules(ctx))

	return c
}

func newChain(cmd *cobra.Command) (*chain.Chain, string, error) {
	absPath, err := filepath.Abs("")
	if err != nil {
		return nil, "", err
	}

	c, err := chain.New(absPath, chain.EnableThirdPartyModuleCodegen())
	if err != nil {
		return nil, "", err
	}

	return c, absPath, nil
}
