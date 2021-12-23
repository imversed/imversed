package currency

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/x/currency/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the currency
	for _, elem := range genState.CurrencyList {
		k.SetCurrency(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.CurrencyList = k.GetAllCurrency(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
