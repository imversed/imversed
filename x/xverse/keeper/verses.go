package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/imversed/imversed/x/xverse/types"
)

func (k Keeper) HasVerse(ctx sdk.Context, verse types.Verse) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)
	return store.Has(types.VerseKey(verse.Name))
}

// SetVerse set a specific verse in the store from its index
func (k Keeper) SetVerse(ctx sdk.Context, verse types.Verse) error {
	if k.HasVerse(ctx, verse) {
		return sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name %s has already exists", verse.Name)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)
	b := k.cdc.MustMarshal(&verse)
	store.Set(types.VerseKey(verse.Name), b)

	return nil
}

// UpdateVerse set a specific verse in the store from its index
func (k Keeper) UpdateVerse(ctx sdk.Context, verse types.Verse) error {
	if !k.HasVerse(ctx, verse) {
		return sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name %s does not exists", verse.Name)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)
	b := k.cdc.MustMarshal(&verse)
	store.Set(types.VerseKey(verse.Name), b)

	return nil
}

// UpdateVerseName set a specific verse in the store from its index
func (k Keeper) UpdateVerseName(ctx sdk.Context, oldName string, newName string) error {
	verse, found := k.GetVerse(ctx, newName)
	if found {
		return sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name \"%s\" already exists", newName)
	}
	verse, found = k.GetVerse(ctx, oldName)
	if !found {
		return sdkerrors.Wrapf(types.ErrVerseAlreadyExists, "verse with name %s does not exists", oldName)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)
	store.Delete(types.VerseKey(verse.Name))

	verse.Name = newName
	b := k.cdc.MustMarshal(&verse)
	store.Set(types.VerseKey(verse.Name), b)

	renameCost := k.GetParams(ctx).TxRenameVerseCost
	ctx.GasMeter().ConsumeGas(renameCost, "txRenameVerse")

	return nil
}

// GetVerse returns a verse from its index
func (k Keeper) GetVerse(
	ctx sdk.Context,
	name string,
) (val types.Verse, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)

	b := store.Get(types.VerseKey(name))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllVerses returns all verse
func (k Keeper) GetAllVerses(ctx sdk.Context) (list []types.Verse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixVerse)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Verse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
