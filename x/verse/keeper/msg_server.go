package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/imversed/imversed/x/verse/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) CreateVerse(
	goCtx context.Context,
	msg *types.MsgCreateVerse,
) (*types.MsgCreateVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	uuid.SetRand(k.rnd)
	verse := types.Verse{Owner: msg.Sender, Name: uuid.NewString(), Icon: msg.Icon, Description: msg.Description}

	err := k.SetVerse(ctx, verse)

	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to create verse")
	}

	return &types.MsgCreateVerseResponse{}, nil
}

func (k Keeper) AddAssetToVerse(
	goCtx context.Context,
	msg *types.MsgAddAssetToVerse,
) (*types.MsgAddAssetToVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	verse, found := k.GetVerse(ctx, msg.VerseName)

	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	// There are maybe correct-signed, but malicious tx with false msg.VerseCreator
	if verse.Owner != msg.VerseCreator {
		return nil, status.Error(codes.Unauthenticated, "verse creator in msg and chain are not same")
	}

	switch msg.AssetType {
	case types.ContractType:
		if slices.Contains(verse.SmartContracts, msg.AssetId) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("verse already contains asset: %s", msg.AssetId))
		}
		verse.SmartContracts = append(verse.SmartContracts, msg.AssetId)
	default:
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid assert type: %s", msg.AssetType))
	}

	err := k.UpdateVerse(ctx, verse)

	if err != nil {
		return nil, err
	}

	return &types.MsgAddAssetToVerseResponse{}, nil
}

func (k Keeper) RemoveAssetFromVerse(
	goCtx context.Context,
	msg *types.MsgRemoveAssetFromVerse,
) (*types.MsgRemoveAssetFromVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	verse, found := k.GetVerse(ctx, msg.VerseName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name %s does not exists", msg.VerseName)
	}
	// There are maybe correct-signed, but malicious tx with false msg.VerseCreator
	if verse.Owner != msg.VerseCreator {
		return nil, status.Error(codes.Unauthenticated, "verse creator in msg and chain are not same")
	}

	switch msg.AssetType {
	case types.ContractType:
		if !slices.Contains(verse.SmartContracts, msg.AssetId) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("verse doesn't contain asset: %s", msg.AssetId))
		}
		// remove asset from slice
		verse.SmartContracts[slices.Index(verse.SmartContracts, msg.AssetId)] = verse.SmartContracts[len(verse.SmartContracts)-1]
		verse.SmartContracts = verse.SmartContracts[:len(verse.SmartContracts)-1]
	default:
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid assert type: %s", msg.AssetType))
	}
	err := k.UpdateVerse(ctx, verse)

	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveAssetFromVerseResponse{}, nil
}

func (k Keeper) RenameVerse(
	goCtx context.Context,
	msg *types.MsgRenameVerse,
) (*types.MsgRenameVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	verse, found := k.GetVerse(ctx, msg.VerseOldName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name %s does not exists", msg.VerseOldName)
	}
	if verse.Owner != msg.VerseCreator {
		return nil, status.Error(codes.Unauthenticated, "verse creator in msg and chain are not same")
	}

	err := k.UpdateVerseName(ctx, msg.VerseOldName, msg.VerseNewName)

	if err != nil {
		return nil, err
	}

	return &types.MsgRenameVerseResponse{}, nil
}
