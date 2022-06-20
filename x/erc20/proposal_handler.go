package erc20

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/imversed/imversed/x/erc20/keeper"
	"github.com/imversed/imversed/x/erc20/types"
)

// NewErc20ProposalHandler creates a governance handler to manage new proposal types.
// It enables RegisterTokenPairProposal to propose a registration of token mapping
func NewErc20ProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ToggleTokenRelayProposal:
			return handleToggleRelayProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleToggleRelayProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ToggleTokenRelayProposal) error {
	pair, err := k.ToggleRelay(ctx, p.Token)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeToggleTokenRelay,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
			sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
		),
	)

	return nil
}
