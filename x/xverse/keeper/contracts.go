package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/xverse/types"
	"strings"
)

func (k Keeper) HasContract(ctx sdk.Context, contract types.Contract) bool {
	contract.Hash = strings.ToLower(contract.Hash)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)
	return store.Has(types.ContractKey(contract.Hash))
}

func (k Keeper) SetContract(ctx sdk.Context, contract types.Contract) error {
	contract.Hash = strings.ToLower(contract.Hash)
	if k.HasContract(ctx, contract) {
		return sdkerrors.Wrapf(types.ErrContractAlreadyMapped, "contract with hash %s has already mapped to verse", contract.Hash)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)
	b := k.cdc.MustMarshal(&contract)
	store.Set(types.ContractKey(contract.Hash), b)

	return nil
}

func (k Keeper) GetContract(ctx sdk.Context, hash string) (val types.Contract, found bool) {
	hash = strings.ToLower(hash)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)

	b := store.Get(types.ContractKey(hash))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// removeContract removes mapping for contract <-> verse pair.
// Must be used only in RemoveAsset method
func (k Keeper) removeContract(ctx sdk.Context, hash string) error {
	hash = strings.ToLower(hash)
	if _, found := k.GetContract(ctx, hash); !found {
		return sdkerrors.Wrapf(types.ErrContractNotFound, "contract with hash %s was not mapped", hash)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)

	store.Delete(types.ContractKey(hash))

	return nil
}

func (k Keeper) updateVerseInContracts(ctx sdk.Context, verseOldName, verseNewName string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var contract types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)

		if contract.Verse == verseOldName {
			contract.Verse = verseNewName
			updates := k.cdc.MustMarshal(&contract)
			store.Set(types.ContractKey(contract.Hash), updates)
		}
	}
}

// GetAllContracts returns all contracts
func (k Keeper) GetAllContracts(ctx sdk.Context) (list []types.Contract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixContract)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
