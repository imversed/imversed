package xverse

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/imversed/imversed/x/xverse/keeper"
	"github.com/imversed/imversed/x/xverse/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// Set all the verse
	for _, elem := range genState.VerseList {
		_ = k.SetVerse(ctx, elem)
		for _, v := range elem.SmartContracts {
			err := k.SetContract(ctx, types.Contract{
				Hash:  v,
				Verse: elem.Name,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetAllVerses(ctx), k.GetParams(ctx))
}
