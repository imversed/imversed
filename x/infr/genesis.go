package infr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/keeper"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
	"github.com/imversed/imversed/x/infr/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)
	ndata := k.GetParams(ctx)
	ndata.MinGasPrices = minGasPriceHelper.Helper.Get()
	k.SetParams(ctx, ndata)
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
