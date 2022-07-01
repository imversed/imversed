package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	gravitytypes "github.com/peggyjv/gravity-bridge/module/v2/x/gravity/types"
)

func (app ImversedApp) setUpgradeHandler(cfg module.Configurator) {
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
}

func (app ImversedApp) setUpgradeStore() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	if upgradeInfo.Name == "v3.4" {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{gravitytypes.ModuleName},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
