package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/x/infr/types"
)

// GetParams returns the total set of infr parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the infr parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
