package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/x/currency/types"
)

func (m msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := m.Keeper.SetCurrency(ctx, types.NewCurrency(msg.Denom, sender)); err != nil {
		return nil, err
	}

	return &types.MsgIssueResponse{}, nil
}
