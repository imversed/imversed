package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/types"
)

// GetSmartContractMetadata - get registered smart-contract metadata
func (k Keeper) GetSmartContractMetadata(ctx sdk.Context, address string) (types.SmartContract, bool) {
	if address == "" {
		return types.SmartContract{}, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))
	var smartcontract types.SmartContract
	bz := store.Get([]byte(address))
	if len(bz) == 0 {
		return types.SmartContract{}, false
	}

	k.cdc.MustUnmarshal(bz, &smartcontract)
	return smartcontract, true
}

// SetSmartContractMetadata set metadata for smart-contract
func (k Keeper) SetSmartContractMetadata(ctx sdk.Context, sc types.SmartContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))
	key := []byte(sc.Address)
	bz := k.cdc.MustMarshal(&sc)
	store.Set(key, bz)
}

// GetAllSmartContractsMetadata - get all smartcontracts metadata
func (k Keeper) GetAllSmartContractsMetadata(ctx sdk.Context) []types.SmartContract {
	var smartcontracts []types.SmartContract

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SmartContractKeyPrefix))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var smartContract types.SmartContract
		k.cdc.MustUnmarshal(iterator.Value(), &smartContract)

		smartcontracts = append(smartcontracts, smartContract)
	}

	return smartcontracts
}
