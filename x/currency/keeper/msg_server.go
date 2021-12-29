package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fulldivevr/imversed/x/currency/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

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

	mintingCost := m.GetParams(ctx).TxMintCurrencyCost
	ctx.GasMeter().ConsumeGas(mintingCost, "txMintCurrency")

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
