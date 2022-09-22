package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) SmartContract(gctx context.Context, req *types.QuerySmartContractRequest) (*types.QuerySmartContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(gctx)

	val, found := k.GetSmartContractMetadata(ctx, req.Address)

	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QuerySmartContractResponse{Sc: &val}, nil
}
