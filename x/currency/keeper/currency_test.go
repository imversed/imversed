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
		items[i].Denom = strconv.Itoa(i)

		keeper.SetCurrency(ctx, items[i])
	}
	return items
}

func TestCurrencyGet(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	items := createNCurrency(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCurrency(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestCurrencyRemove(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	items := createNCurrency(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCurrency(ctx,
			item.Denom,
		)
		_, found := keeper.GetCurrency(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestCurrencyGetAll(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	items := createNCurrency(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCurrency(ctx))
}
