package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fulldivevr/imversed/x/currency/types"
)

func (m msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom := msg.Coin.Denom
	currency, found := m.Keeper.GetCurrency(ctx, denom)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCurrency, "currency with denom %s does not exists", currency.Denom)
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(currency.Owner)
	if err != nil {
		return nil, err
	}

	if !owner.Equals(sender) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCurrency, "sender is not owner")
	}

	if err := m.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Coin)); err != nil {
		return nil, err
	}

	if err := m.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(msg.Coin)); err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
