package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/imversed/imversed/x/xverse/migrations/v2/types"
	"github.com/imversed/imversed/x/xverse/types"
	"strings"
)

// MigrateStore migrates verse and create smart-contract mapping in KVStore for
// In-Place Store migration logic.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	// get verse store
	store := ctx.KVStore(storeKey)

	// get iterator for verse store
	storeIter := store.Iterator(nil, nil)
	defer storeIter.Close()

	// get prefixed store for smart-contract mapping
	contractStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefixContract)

	for ; storeIter.Valid(); storeIter.Next() {
		// get old verse
		oldBz := storeIter.Value()
		var oldVerse v2.Verse
		err := cdc.Unmarshal(oldBz, &oldVerse)
		if err != nil {
			return err
		}

		// loverCase-ing contracts hass
		var toLowerSmartContracts []string
		for _, v := range oldVerse.SmartContracts {
			toLowerSmartContracts = append(toLowerSmartContracts, strings.ToLower(v))
		}

		// migrate verse
		newVerse := types.Verse{
			Owner:             oldVerse.Owner,
			Name:              oldVerse.Name,
			Icon:              oldVerse.Icon,
			Description:       oldVerse.Description,
			SmartContracts:    toLowerSmartContracts,
			Oracle:            "",
			AuthenticatedKeys: nil,
		}

		// add all contracts that bounded to verse in contract -> mapping
		for _, v := range newVerse.SmartContracts {
			contract := types.Contract{
				Hash:  v,
				Verse: newVerse.Name,
			}
			b := cdc.MustMarshal(&contract)
			contractStore.Set(types.ContractKey(contract.Hash), b)
		}

		// set migrated verse
		bz := cdc.MustMarshal(&newVerse)
		store.Set(storeIter.Key(), bz)
	}
	return nil
}
