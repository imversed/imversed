package keeper_test

import (
	"testing"

	testkeeper "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CurrencyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.TxMintCurrencyCost, k.TxMintCurrencyCost(ctx))
}
