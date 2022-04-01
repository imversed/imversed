package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/fulldivevr/imversed/x/currency/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Issue(ctx sdk.Context, denom string, owner sdk.AccAddress, icon string) error {
	return k.SetCurrency(ctx, types.NewCurrency(denom, owner, icon))
}

func (k Keeper) Mint(ctx sdk.Context, coin sdk.Coin, owner sdk.AccAddress) error {
	denom := coin.Denom
	currency, found := k.GetCurrency(ctx, denom)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidCurrency, "currency with denom %s does not exists", denom)
	}

	expected, err := sdk.AccAddressFromBech32(currency.Owner)
	if err != nil {
		return err
	}

	if !owner.Equals(expected) {
		return sdkerrors.Wrapf(types.ErrInvalidCurrency, "sender is not owner")
	}

	mintingCost := k.GetParams(ctx).TxMintCurrencyCost
	ctx.GasMeter().ConsumeGas(mintingCost, "txMintCurrency")

	coins := sdk.NewCoins(coin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coins); err != nil {
		return err
	}

	return nil
}

func (k Keeper) MigrationAddIcon(ctx sdk.Context) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrencyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var currency types.Currency
		k.cdc.MustUnmarshal(iterator.Value(), &currency)
		currency.Icon = ""
		updates := k.cdc.MustMarshal(&currency)

		store.Set(types.CurrencyKey(
			currency.Denom,
		), updates)
	}
	return nil
}
