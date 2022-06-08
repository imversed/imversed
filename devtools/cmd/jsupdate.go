package cmd

import (
	"context"
	"fmt"

	"github.com/ignite-hq/cli/ignite/pkg/cliui/clispinner"
	"github.com/ignite-hq/cli/ignite/services/chain"
	"github.com/spf13/cobra"
)

const (
	flagPath          = "path"
	flagHome          = "home"
	flagProto3rdParty = "proto-all-modules"
	flagYes           = "yes"
)

func NewJSUpdate(ctx context.Context) *cobra.Command {

	c := &cobra.Command{
		Use:   "js-update",
		Short: "devtools make sense",
		RunE:  JSUpdateHandler,
	}

	return c
}

func JSUpdateHandler(cmd *cobra.Command, args []string) error {
	s := clispinner.New().SetText("JS Client update...")
	defer s.Stop()

	c, err := newChain(cmd)
	if err != nil {
		return err
	}

	if err := c.Generate(cmd.Context(), chain.GenerateVuex()); err != nil {
		return err
	}

	s.Stop()
	fmt.Println("JS Client updated.")

	return nil
}
