package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/imversed/imversed/x/verse/types"
	"github.com/spf13/cobra"
	"strings"
)

// GetQueryCmd returns the parent command for all verse CLI query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the verse module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetVersesCmd(),
		GetVerseCmd(),
		GetParamsCmd(),
		HasAssetCmd(),
		GetAssetsCmd(),
	)
	return cmd
}

// GetVersesCmd queries a verses registered
func GetVersesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verses",
		Short: "Gets all verses registered",
		Long:  "Gets all verses registered",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryAllVerseRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VerseAll(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetVerseCmd queries a verse registered
func GetVerseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verse [name]",
		Short: "Get a verse registered",
		Long:  "Get a verse registered",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryGetVerseRequest{
				VerseName: args[0],
			}

			res, err := queryClient.Verse(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetParamsCmd queries hub info
func GetParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Gets verse module params",
		Long:  "Gets verse module params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// HasAssetCmd returns true if verse has asset
func HasAssetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "has-asset [type/prefix] [verse_name]",
		Short: "Gets verse module params",
		Long:  "Gets verse module params",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			splitted := strings.Split(args[0], "/")
			if len(splitted) != 2 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid asset '%s'. use [type/id] expression", args[0])
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryHasAssetRequest{
				VerseName: args[1],
				AssetType: splitted[0],
				AssetId:   splitted[1],
			}

			res, err := queryClient.HasAsset(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetAssetsCmd returns all assets for verse
func GetAssetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets [verse_name]",
		Short: "Gets all verses' assets",
		Long:  "Gets all verses' assets",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryGetVerseAssetsRequest{VerseName: args[0]}

			res, err := queryClient.GetAssets(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
