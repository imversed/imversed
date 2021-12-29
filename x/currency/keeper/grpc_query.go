package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/fulldivevr/imversed/x/currency/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CurrencyAll(c context.Context, req *types.QueryAllCurrencyRequest) (*types.QueryAllCurrencyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var currencys []types.Currency
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	currencyStore := prefix.NewStore(store, types.KeyPrefix(types.CurrencyKeyPrefix))

	pageRes, err := query.Paginate(currencyStore, req.Pagination, func(key []byte, value []byte) error {
		var currency types.Currency
		if err := k.cdc.Unmarshal(value, &currency); err != nil {
			return err
		}

		currencys = append(currencys, currency)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCurrencyResponse{Currency: currencys, Pagination: pageRes}, nil
}

func (k Keeper) Currency(c context.Context, req *types.QueryGetCurrencyRequest) (*types.QueryGetCurrencyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCurrency(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCurrencyResponse{Currency: val}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
