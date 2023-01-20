package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/imversed/imversed/x/xverse/types"
	"golang.org/x/exp/slices"
)

type Hooks struct {
	k Keeper
}

var _ evmtypes.EvmHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) PostTxProcessing(
	ctx sdk.Context,
	msg core.Message,
	receipt *ethtypes.Receipt,
) error {
	// if smart-contract deploying
	if msg.To() == nil {
		return nil
	}

	contract, found := h.k.GetContract(ctx, msg.To().Hex())
	if !found {
		return nil
	}

	verse, found := h.k.GetVerse(ctx, contract.Verse)

	if !found {
		return nil
	}

	cosmosAddress := sdk.AccAddress(msg.From().Bytes())

	if verse.Owner == cosmosAddress.String() {
		return nil
	}

	if !slices.Contains(verse.AuthenticatedKeys, cosmosAddress.String()) {
		return sdkerrors.Wrapf(types.ErrNotAuthenticated, "%s", cosmosAddress.String())
	}

	return nil
}
