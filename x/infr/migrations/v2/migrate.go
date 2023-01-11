package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/types"
	"strings"
)

// MigrateStore migrates the BaseFee value from the store to the params for
// In-Place Store migration logic.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))

	storeIter := store.Iterator(nil, nil)
	defer storeIter.Close()

	for ; storeIter.Valid(); storeIter.Next() {
		bz := storeIter.Value()
		var contract types.SmartContract

		err := cdc.Unmarshal(bz, &contract)
		if err != nil {
			return err
		}

		contract.Address = strings.ToLower(contract.Address)
		bz = cdc.MustMarshal(&contract)

		key := []byte(contract.Address)
		store.Set(key, bz)
		store.Delete(storeIter.Key())
	}
	return nil
}
