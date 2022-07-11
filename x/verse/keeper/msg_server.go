package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/imversed/imversed/x/verse/types"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) CreateVerse(
	goCtx context.Context,
	msg *types.MsgCreateVerse,
) (*types.MsgCreateVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	verse := types.Verse{Owner: msg.Sender, Name: msg.Name, Icon: msg.Icon, Description: msg.Description}

	err := k.SetVerse(ctx, verse)

	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to create verse")
	}

	return &types.MsgCreateVerseResponse{}, nil
}
