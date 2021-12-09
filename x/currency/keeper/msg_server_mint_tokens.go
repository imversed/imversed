package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/x/currency/types"
)

func (k msgServer) MintTokens(goCtx context.Context, msg *types.MsgMintTokens) (*types.MsgMintTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintTokensResponse{}, nil
}
