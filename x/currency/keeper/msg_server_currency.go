package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fulldivevr/imversed/x/currency/types"
)

func (k msgServer) CreateCurrency(goCtx context.Context, msg *types.MsgCreateCurrency) (*types.MsgCreateCurrencyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCurrency(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var currency = types.Currency{
		Owner: msg.Owner,
		Denom: msg.Denom,
	}

	k.SetCurrency(
		ctx,
		currency,
	)
	return &types.MsgCreateCurrencyResponse{}, nil
}

func (k msgServer) UpdateCurrency(goCtx context.Context, msg *types.MsgUpdateCurrency) (*types.MsgUpdateCurrencyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCurrency(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var currency = types.Currency{
		Owner: msg.Owner,
		Denom: msg.Denom,
	}

	k.SetCurrency(ctx, currency)

	return &types.MsgUpdateCurrencyResponse{}, nil
}

func (k msgServer) DeleteCurrency(goCtx context.Context, msg *types.MsgDeleteCurrency) (*types.MsgDeleteCurrencyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCurrency(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCurrency(
		ctx,
		msg.Denom,
	)

	return &types.MsgDeleteCurrencyResponse{}, nil
}
