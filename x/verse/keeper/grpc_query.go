package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/imversed/imversed/x/verse/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Verse(c context.Context, req *types.QueryGetVerseRequest) (*types.QueryGetVerseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVerse(
		ctx,
		req.VerseName,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, " verse not found")
	}

	return &types.QueryGetVerseResponse{Verse: val}, nil
}

func (k Keeper) VerseAll(c context.Context, req *types.QueryAllVerseRequest) (*types.QueryAllVerseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var verses []types.Verse
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	verseStore := prefix.NewStore(store, types.KeyPrefixVerse)

	pageRes, err := query.Paginate(verseStore, req.Pagination, func(key []byte, value []byte) error {
		var verse types.Verse
		if err := k.cdc.Unmarshal(value, &verse); err != nil {
			return err
		}

		verses = append(verses, verse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVerseResponse{Verse: verses, Pagination: pageRes}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) HasAsset(c context.Context, req *types.QueryHasAssetRequest) (*types.QueryHasAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVerse(
		ctx,
		req.VerseName,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "verse not found")
	}

	switch req.AssetType {
	case types.ContractType:
		return &types.QueryHasAssetResponse{HasAsset: slices.Contains(val.SmartContracts, req.AssetId)}, nil
	default:
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid assert type: %s", req.AssetType))
	}
}

func (k Keeper) GetAssets(c context.Context, req *types.QueryGetVerseAssetsRequest) (*types.QueryGetVerseAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVerse(
		ctx,
		req.VerseName,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "verse not found")
	}

	return &types.QueryGetVerseAssetsResponse{
		Assets: val.SmartContracts,
	}, nil
}
