package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	versetypes "github.com/imversed/imversed/x/verse/types"
)

func (app ImversedApp) setUpgradeHandler(cfg module.Configurator) {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == "v3.8" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{versetypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.1",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.2",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.4",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.5",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.6",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.8",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

}
