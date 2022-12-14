package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/xverse/types"
)

// GetParams returns the total set of verse parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the verse parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
