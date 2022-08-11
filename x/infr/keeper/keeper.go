package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/imversed/imversed/x/infr/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of this module maintains collections of erc20.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	evmKeeper     types.EVMKeeper
	erc20Keeper   types.Erc20Keeper
}

// NewKeeper creates new instances of the erc20 Keeper
func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	ek types.EVMKeeper,
	er types.Erc20Keeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramstore: ps,

		accountKeeper: ak,
		evmKeeper:     ek,
		erc20Keeper:   er,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
