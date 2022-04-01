package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/currency/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.TxMintCurrencyCost(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// TxMintCurrencyCost returns the TxMintCurrencyCost param
func (k Keeper) TxMintCurrencyCost(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyTxMintCurrencyCost, &res)
	return
}
