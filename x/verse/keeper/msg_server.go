package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/imversed/imversed/x/verse/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) CreateVerse(
	goCtx context.Context,
	msg *types.MsgCreateVerse,
) (*types.MsgCreateVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	verse := types.Verse{Owner: msg.Sender, Name: uuid.NewString(), Icon: msg.Icon, Description: msg.Description}

	err := k.SetVerse(ctx, verse)

	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to create verse")
	}

	return &types.MsgCreateVerseResponse{}, nil
}

func (k Keeper) AddAssetToVerse(goCtx context.Context, msg *types.MsgAddAssetToVerse) (*types.MsgAddAssetToVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	verse, found := k.GetVerse(ctx, msg.VerseName)

	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	switch msg.AssetType {
	case types.ContractType:
		verse.SmartContracts = append(verse.SmartContracts, msg.AssetId)
	}
	err := k.SetVerse(ctx, verse)

	if err != nil {
		return nil, err
	}

	return &types.MsgAddAssetToVerseResponse{}, nil
}
