package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/imversed/imversed/x/xverse/migrations/v2/types"
	"github.com/imversed/imversed/x/xverse/types"
)

// MigrateStore migrates the BaseFee value from the store to the params for
// In-Place Store migration logic.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	storeIter := store.Iterator(nil, nil)
	defer storeIter.Close()

	for ; storeIter.Valid(); storeIter.Next() {
		oldBz := storeIter.Value()
		var oldVerse v2.Verse
		err := cdc.Unmarshal(oldBz, &oldVerse)
		if err != nil {
			return err
		}

		newVerse := types.Verse{
			Owner:             oldVerse.Owner,
			Name:              oldVerse.Name,
			Icon:              oldVerse.Icon,
			Description:       oldVerse.Description,
			SmartContracts:    oldVerse.SmartContracts,
			Oracle:            "",
			AuthenticatedKeys: nil,
		}

		bz := cdc.MustMarshal(&newVerse)

		store.Set(storeIter.Key(), bz)
	}
	return nil
}
