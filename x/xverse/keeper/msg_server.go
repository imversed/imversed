package keeper

import (
	"context"
	sdkerrors "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/imversed/imversed/x/xverse/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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

	msg.AssetId = strings.ToLower(msg.AssetId)

	switch msg.AssetType {
	case types.ContractType:
		if slices.Contains(verse.SmartContracts, msg.AssetId) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("verse already contains asset: %s", msg.AssetId))
		}
		verse.SmartContracts = append(verse.SmartContracts, msg.AssetId)

		err := k.SetContract(ctx, types.Contract{
			Hash:  msg.AssetId,
			Verse: verse.Name,
		})

		if err != nil {
			return nil, err
		}
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
		return nil, sdkerrors.Wrapf(types.ErrVerseNotfound, "verse with name %s does not exists", msg.VerseName)
	}
	// There are maybe correct-signed, but malicious tx with false msg.VerseCreator
	if verse.Owner != msg.VerseCreator {
		return nil, status.Error(codes.Unauthenticated, "verse creator in msg and chain are not same")
	}

	msg.AssetId = strings.ToLower(msg.AssetId)

	switch msg.AssetType {
	case types.ContractType:
		if !slices.Contains(verse.SmartContracts, msg.AssetId) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("verse doesn't contain asset: %s", msg.AssetId))
		}
		// remove asset from slice
		verse.SmartContracts[slices.Index(verse.SmartContracts, msg.AssetId)] = verse.SmartContracts[len(verse.SmartContracts)-1]
		verse.SmartContracts = verse.SmartContracts[:len(verse.SmartContracts)-1]

		err := k.removeContract(ctx, msg.AssetId)
		if err != nil {
			return nil, err
		}
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
		return nil, sdkerrors.Wrapf(types.ErrVerseNotfound, "verse with name \"%s\" does not exists", msg.VerseOldName)
	}
	if verse.Owner != msg.VerseCreator {
		return nil, status.Error(codes.Unauthenticated, "verse creator in msg and chain are not same")
	}

	err := k.UpdateVerseName(ctx, msg.VerseOldName, msg.VerseNewName)

	if err != nil {
		return nil, err
	}

	k.updateVerseInContracts(ctx, msg.VerseOldName, msg.VerseNewName)

	return &types.MsgRenameVerseResponse{}, nil
}

func (k Keeper) AddOracleToVerse(
	goCtx context.Context,
	msg *types.MsgAddOracleToVerse,
) (*types.MsgAddOracleToVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	verse, found := k.GetVerse(ctx, msg.VerseName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVerseNotfound, "verse with name \"%s\" does not exists", msg.VerseName)
	}

	_, err := sdk.AccAddressFromBech32(msg.Oracle)
	if err != nil {
		return nil, err
	}

	if verse.Owner != msg.GetSigners()[0].String() {
		return nil, status.Error(codes.Unauthenticated, "sender in msg and verse owner are not same")
	}

	_, err = sdk.AccAddressFromBech32(msg.Oracle)
	if err != nil {
		return nil, err
	}

	verse.Oracle = msg.Oracle

	err = k.UpdateVerse(ctx, verse)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddOracleToVerseResponse{}, nil
}

func (k Keeper) AuthorizeKeyToVerse(
	goCtx context.Context,
	msg *types.MsgAuthorizeKeyToVerse,
) (*types.MsgAuthorizeKeyToVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	verse, found := k.GetVerse(ctx, msg.VerseName)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVerseNotfound, "verse with name \"%s\" does not exists", msg.VerseName)
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	if verse.Oracle != msg.GetSigners()[0].String() && verse.Owner != msg.GetSigners()[0].String() {
		return nil, status.Error(codes.Unauthenticated, "sender has not authorized to add keys")
	}

	if slices.Contains(verse.AuthenticatedKeys, msg.Address) {
		return nil, status.Error(codes.AlreadyExists, "address already authorized")
	}

	verse.AuthenticatedKeys = append(verse.AuthenticatedKeys, msg.Address)

	err = k.UpdateVerse(ctx, verse)
	if err != nil {
		return nil, err
	}

	return &types.MsgAuthorizeKeyToVerseResponse{}, nil
}

func (k Keeper) DeauthorizeKeyToVerse(
	goCtx context.Context,
	msg *types.MsgDeauthorizeKeyToVerse,
) (*types.MsgDeauthorizeKeyToVerseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	verse, found := k.GetVerse(ctx, msg.VerseName)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVerseNotfound, "verse with name \"%s\" does not exists", msg.VerseName)
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	if verse.Oracle != msg.GetSigners()[0].String() && verse.Owner != msg.GetSigners()[0].String() {
		return nil, status.Error(codes.Unauthenticated, "sender has not authorized to add keys")
	}

	if !slices.Contains(verse.AuthenticatedKeys, msg.Address) {
		return nil, status.Error(codes.NotFound, "address not authorized")
	}

	// remove key from authorized keys
	verse.AuthenticatedKeys[slices.Index(verse.AuthenticatedKeys, msg.Address)] = verse.AuthenticatedKeys[len(verse.AuthenticatedKeys)-1]
	verse.AuthenticatedKeys = verse.AuthenticatedKeys[:len(verse.AuthenticatedKeys)-1]

	err = k.UpdateVerse(ctx, verse)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeauthorizeKeyToVerseResponse{}, nil
}
