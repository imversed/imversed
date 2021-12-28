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

	mintingCost := m.GetParams(ctx).TxMintCurrencyCost
	ctx.GasMeter().ConsumeGas(mintingCost, "txMintCurrency")

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

func (m msgServer) UpdateCurrency(goCtx context.Context, msg *types.MsgUpdateCurrency) (*types.MsgUpdateCurrencyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := m.GetCurrency(
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

	m.SetCurrency(ctx, currency)

	return &types.MsgUpdateCurrencyResponse{}, nil
}
