package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/fulldivevr/imversed/x/currency/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdIssue())
	cmd.AddCommand(CmdMint())
	// this line is used by starport scaffolding # 1

	return cmd
}

var _ = strconv.Itoa(0)

func CmdIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "issue [denom]",
		Long:    "Issue a new token denomination",
		Example: fmt.Sprintf("$ %s tx %s issue <denom> --from=<key-name>", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssue(
				clientCtx.GetFromAddress().String(),
				argDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint [token]",
		Short:   "Mint tokens for particular denomination",
		Example: fmt.Sprintf("$ %s tx %s mint 100500tokenname --from=<key-name>", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
				clientCtx.GetFromAddress().String(),
				coin,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
