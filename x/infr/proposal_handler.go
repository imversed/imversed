package infr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/imversed/imversed/x/infr/keeper"
	"github.com/imversed/imversed/x/infr/types"
)

// NewInfrProposalHandler creates a governance handler to manage new proposal types.
// It enables RegisterTokenPairProposal to propose a registration of token mapping
func NewInfrProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ChangeMinGasPricesProposal:
			return handleChangeMinGasPricesProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleChangeMinGasPricesProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ChangeMinGasPricesProposal) error {

	pa := k.GetParams(ctx)
	pa.MinGasPrices = p.MinGasPrices
	k.SetParams(ctx, pa)
	//pair, err := k.UpdateTokenPairERC20(ctx, p.GetERC20Address(), p.GetNewERC20Address())
	//if err != nil {
	//	return err
	//}

	//ctx.EventManager().EmitEvent(
	//	sdk.NewEvent(
	//		types.EventTypeUpdateTokenPairERC20,
	//		sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
	//		sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
	//	),
	//)

	return nil
}
