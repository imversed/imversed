package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "github.com/imversed/imversed/x/xverse/migrations/v3"
	"github.com/imversed/imversed/x/xverse/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	verseStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefixVerse)

	storeIter := sdk.KVStorePrefixIterator(verseStore, []byte{})
	defer storeIter.Close()

	for ; storeIter.Valid(); storeIter.Next() {
		var oldVerse v3.Verse
		err := cdc.Unmarshal(storeIter.Value(), &oldVerse)
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
			Dao:               "",
		}

		bz := cdc.MustMarshal(&newVerse)
		verseStore.Set(storeIter.Key(), bz)
	}
	return nil
}
