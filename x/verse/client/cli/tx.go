package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/imversed/imversed/x/verse/types"
	"github.com/spf13/cobra"
)

// NewTxCmd returns a root CLI command handler for certain modules/verse transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "verse subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		CreateNewVerseCmd(),
	)
	return txCmd
}

// CreateNewVerseCmd returns a CLI command handler for creating verse
func CreateNewVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-verse [name] [description] [icon]",
		Short: "Create new verse",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := cliCtx.GetFromAddress()

			icon := ""

			if len(args) == 3 {
				icon = args[2]
			}

			msg := &types.MsgCreateVerse{
				Name:        args[0],
				Description: args[1],
				Icon:        icon,

				Sender: sender.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
