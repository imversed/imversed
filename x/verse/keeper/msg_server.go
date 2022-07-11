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

	verse := types.Verse{Owner: msg.Sender, Name: msg.Name}

	// check if the denomination already registered
	if k.HasVerse(ctx, verse) {
		return nil, sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse already registered: %s", verse.Name)
	}
	err := k.SetVerse(ctx, verse)

	_ = err

	return &types.MsgCreateVerseResponse{}, nil
}
