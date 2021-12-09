package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/fulldivevr/imversed/x/currency/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdMintTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount]",
		Short: "mint new tokens",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintTokens(
				clientCtx.GetFromAddress().String(),
				coins,
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
