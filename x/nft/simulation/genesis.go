package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/fulldivevr/metachain/x/nft/types"
)

// RandomizedGenState generates a random GenesisState for nft
func RandomizedGenState(simState *module.SimulationState) {
	nftGenesis := types.NewGenesisState()

	bz, err := json.MarshalIndent(nftGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(nftGenesis)
}
