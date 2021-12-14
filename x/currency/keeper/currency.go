package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/x/currency/types"
)

// SetCurrency set a specific currency in the store from its index
func (k Keeper) SetCurrency(ctx sdk.Context, currency types.Currency) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrencyKeyPrefix))
	b := k.cdc.MustMarshal(&currency)
	store.Set(types.CurrencyKey(
		currency.Denom,
	), b)
}

// GetCurrency returns a currency from its index
func (k Keeper) GetCurrency(
	ctx sdk.Context,
	denom string,

) (val types.Currency, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrencyKeyPrefix))

	b := store.Get(types.CurrencyKey(
		denom,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCurrency removes a currency from the store
func (k Keeper) RemoveCurrency(
	ctx sdk.Context,
	denom string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrencyKeyPrefix))
	store.Delete(types.CurrencyKey(
		denom,
	))
}

// GetAllCurrency returns all currency
func (k Keeper) GetAllCurrency(ctx sdk.Context) (list []types.Currency) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrencyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Currency
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
