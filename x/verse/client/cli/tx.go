package cli

import (
	"context"
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
		AddAssetToVerseCmd(),
	)
	return txCmd
}

// CreateNewVerseCmd returns a CLI command handler for creating verse
func CreateNewVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-verse [description] [icon]",
		Short: "Create new verse with auto-generated name",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := cliCtx.GetFromAddress()

			icon := ""

			if len(args) == 2 {
				icon = args[1]
			}

			msg := &types.MsgCreateVerse{
				Description: args[0],
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

// AddAssetToVerseCmd add asset to existing verse
func AddAssetToVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-asset [prefix/asset_id] [verse_name]",
		Short: "Add asset to existing verse",
		Long:  "Add asset to existing verse. CONTRACT: tx must be signed by verse creator and asset creator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := cliCtx.GetFromAddress()
			node, _ := cliCtx.GetNode()

			h := int64(16)
			res, _ := node.Block(context.Background(), &h)

			_ = res

			msg := &types.MsgCreateVerse{
				//Name:        args[0],
				Description: args[1],

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
