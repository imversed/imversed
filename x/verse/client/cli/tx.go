package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	infrtypes "github.com/imversed/imversed/x/infr/types"
	"github.com/imversed/imversed/x/verse/types"
	"github.com/spf13/cobra"
	"strings"
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
		RenameVerseCmd(),
		RemoveAssetFromVerseCmd(),
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

			splitted := strings.Split(args[0], "/")
			if len(splitted) != 2 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid asset '%s'. use [type/id] expression", args[0])
			}

			queryClient := types.NewQueryClient(cliCtx)

			verseParams := &types.QueryGetVerseRequest{
				VerseName: args[1],
			}

			verseResp, err := queryClient.Verse(context.Background(), verseParams)
			if err != nil {
				return err
			}

			assetParams := &infrtypes.QuerySmartContractRequest{
				Address: splitted[1],
			}
			infrQueryClient := infrtypes.NewQueryClient(cliCtx)

			assetResponse, err := infrQueryClient.SmartContract(context.Background(), assetParams)
			if err != nil {
				return err
			}

			msg := &types.MsgAddAssetToVerse{
				Sender:       sender.String(),
				VerseName:    args[1],
				AssetType:    splitted[0],
				AssetId:      splitted[1],
				AssetCreator: assetResponse.Sc.Creator,
				VerseCreator: verseResp.Verse.Owner,
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

// RenameVerseCmd rename existing verse
func RenameVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename-verse [verse_old_name] [verse_new_name]",
		Short: "Rename verse, new name must be unique",
		Long:  "Rename verse, new name must be unique. CONTRACT: tx must be signed by verse creator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := cliCtx.GetFromAddress()

			queryClient := types.NewQueryClient(cliCtx)

			verseParams := &types.QueryGetVerseRequest{
				VerseName: args[0],
			}

			verseResp, err := queryClient.Verse(context.Background(), verseParams)
			if err != nil {
				return err
			}

			msg := &types.MsgRenameVerse{
				Sender:       sender.String(),
				VerseCreator: verseResp.Verse.Owner,
				VerseOldName: args[0],
				VerseNewName: args[1],
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

// RemoveAssetFromVerseCmd remove existing asset from verse
func RemoveAssetFromVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-asset [prefix/asset_id] [verse_name]",
		Short: "Remove asset from verse",
		Long:  "Remove asset from verse. CONTRACT: tx must be signed by verse creator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := cliCtx.GetFromAddress()

			splitted := strings.Split(args[0], "/")
			if len(splitted) != 2 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid asset '%s'. use [type/id] expression", args[0])
			}

			queryClient := types.NewQueryClient(cliCtx)

			verseParams := &types.QueryGetVerseRequest{
				VerseName: args[1],
			}

			verseResp, err := queryClient.Verse(context.Background(), verseParams)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveAssetFromVerse{
				Sender:       sender.String(),
				VerseName:    args[1],
				AssetType:    splitted[0],
				AssetId:      splitted[1],
				VerseCreator: verseResp.Verse.Owner,
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
