package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	infrtypes "github.com/imversed/imversed/x/infr/types"
)

// Logger returns a module-specific logger.
func (k Keeper) BurnCoinsFromModuleAccount(ctx sdk.Context) error {
	moduleAddress := k.accountKeeper.GetModuleAddress(infrtypes.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, "aimv")

	if balance.IsZero() {
		return nil
	}

	err := k.bankKeeper.BurnCoins(ctx, infrtypes.ModuleName, sdk.Coins{balance})
	if err != nil {
		return err
	}

	return nil
}
