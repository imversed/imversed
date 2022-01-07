package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCurrency(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Currency {
	items := make([]types.Currency, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i * 100)

		keeper.SetCurrency(ctx, items[i])
	}
	return items
}

func TestCurrencyGet(t *testing.T) {
	k, ctx := keepertest.CurrencyKeeper(t, nil)
	items := createNCurrency(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetCurrency(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestCurrencyGetAll(t *testing.T) {
	k, ctx := keepertest.CurrencyKeeper(t, nil)
	items := createNCurrency(k, ctx, 10)
	saved := k.GetAllCurrency(ctx)
	require.ElementsMatch(t, items, saved)
}
