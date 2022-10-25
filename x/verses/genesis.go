package verses

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/imversed/imversed/x/verses/keeper"
	"github.com/imversed/imversed/x/verses/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// Set all the verse
	for _, elem := range genState.VerseList {
		_ = k.SetVerse(ctx, elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetAllVerses(ctx), k.GetParams(ctx))
}
