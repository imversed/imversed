package currency_test

import (
	"testing"

	keepertest "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency"
	"github.com/fulldivevr/imversed/x/currency/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		CurrencyList: []types.Currency{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CurrencyKeeper(t)
	currency.InitGenesis(ctx, *k, genesisState)
	got := currency.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.CurrencyList, len(genesisState.CurrencyList))
	require.Subset(t, genesisState.CurrencyList, got.CurrencyList)
	// this line is used by starport scaffolding # genesis/test/assert
}
