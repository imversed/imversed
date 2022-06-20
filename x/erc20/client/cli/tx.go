package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/ethereum/go-ethereum/common"

	imversed "github.com/tharsis/ethermint/types"

	"github.com/imversed/imversed/x/erc20/types"
)

// NewTxCmd returns a root CLI command handler for certain modules/erc20 transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "erc20 subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConvertCoinCmd(),
		NewConvertERC20Cmd(),
		NewUpdateTokenPairERC20Cmd(),
		NewRegisterCoinCmd(),
		ToggleTokenRelayCmd(),
		NewRegisterERC20Cmd(),
	)
	return txCmd
}

// NewConvertCoinCmd returns a CLI command handler for converting cosmos coins
func NewConvertCoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-coin [coin] [receiver_hex]",
		Short: "Convert a Cosmos coin to ERC20",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			var receiver string
			sender := cliCtx.GetFromAddress()

			if len(args) == 2 {
				receiver = args[1]
				if err := imversed.ValidateAddress(receiver); err != nil {
					return fmt.Errorf("invalid receiver hex address %w", err)
				}
			} else {
				receiver = common.BytesToAddress(sender).Hex()
			}

			msg := &types.MsgConvertCoin{
				Coin:     coin,
				Receiver: receiver,
				Sender:   sender.String(),
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

// NewConvertERC20Cmd returns a CLI command handler for converting ERC20s
func NewConvertERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-erc20 [contract-address] [amount] [receiver]",
		Short: "Convert an ERC20 token to Cosmos coin",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contract := args[0]
			if err := imversed.ValidateAddress(contract); err != nil {
				return fmt.Errorf("invalid ERC20 contract address %w", err)
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount %s", args[1])
			}

			from := common.BytesToAddress(cliCtx.GetFromAddress().Bytes())

			receiver := cliCtx.GetFromAddress()
			if len(args) == 3 {
				receiver, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
			}

			msg := &types.MsgConvertERC20{
				ContractAddress: contract,
				Amount:          amount,
				Receiver:        receiver.String(),
				Sender:          from.Hex(),
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

// NewRegisterCoinProposalCmd implements the command to submit a community-pool-spend proposal
func NewRegisterCoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-coin [metadata]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a register coin proposal",
		Long: `Submit a proposal to register a Cosmos coin to the erc20 along with an initial deposit.
Upon passing, the
The proposal details must be supplied via a JSON file.`,
		Example: fmt.Sprintf(`$ %s tx gov submit-proposal register-coin <path/to/metadata.json> --from=<key_or_address>

Where metadata.json contains (example):

{
  "description": "staking, gas and governance token of the Evmos testnets"
  "denom_units": [
		{
			"denom": "aevmos",
			"exponent": 0,
			"aliases": ["atto evmos"]
		},
		{
			"denom": "evmos",
			"exponent": 18
		}
	],
	"base": "aevmos",
	"display: "evmos",
	"name": "Evmos",
	"symbol": "EVMOS"
}`, version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			metadata, err := ParseMetadata(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			msg := &types.MsgRegisterCoin{
				Metadata: metadata,
				Sender:   from.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewRegisterERC20Cmd implements the command to submit a community-pool-spend proposal
func NewRegisterERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-erc20 [erc20-address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Submit a proposal to register an ERC20 token",
		Long:    "Submit a proposal to register an ERC20 token to the erc20 along with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx gov submit-proposal register-erc20 <path/to/proposal.json> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			erc20Addr := args[0]

			from := clientCtx.GetFromAddress()

			msg := &types.MsgRegisterERC20{
				Erc20Address: erc20Addr,
				Sender:       from.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// ToggleTokenRelayCmd implements the command to toggle token relay
func ToggleTokenRelayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggle-token-relay [token]",
		Args:    cobra.ExactArgs(1),
		Short:   "Submit a toggle token relay proposal",
		Long:    "Submit a proposal to toggle the relaying of a token pair along with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx erc20 toggle-token-relay <denom_or_contract> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			token := args[0]

			msg := &types.MsgToggleTokenRelay{
				Token:  token,
				Sender: from.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpdateTokenPairERC20Cmd implements the command to update erc20 token pair
func NewUpdateTokenPairERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-token-pair-erc20 [erc20_address] [new_erc20_address]",
		Args:    cobra.ExactArgs(2),
		Short:   "Update token pair ERC20",
		Long:    `Update the ERC20 address of a token pair.`,
		Example: fmt.Sprintf("$ %s tx erc20 update-token-pair-erc20 <erc20_address> <new_erc20_address> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			erc20Addr := args[0]
			if err := imversed.ValidateAddress(erc20Addr); err != nil {
				return fmt.Errorf("invalid current ERC20 contract address %w", err)
			}

			newERC20Addr := args[1]
			if err := imversed.ValidateAddress(newERC20Addr); err != nil {
				return fmt.Errorf("invalid new ERC20 contract address %w", err)
			}

			from := clientCtx.GetFromAddress()

			msg := &types.MsgUpdateTokenPairERC20{
				Erc20Address:    erc20Addr,
				NewErc20Address: newERC20Addr,
				Sender:          from.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
