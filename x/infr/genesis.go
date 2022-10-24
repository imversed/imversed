package infr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/keeper"
	"github.com/imversed/imversed/x/infr/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)

	for _, sc := range data.SmartContracts {
		k.SetSmartContractMetadata(ctx, sc)
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:         k.GetParams(ctx),
		SmartContracts: k.GetAllSmartContractsMetadata(ctx),
	}
}
