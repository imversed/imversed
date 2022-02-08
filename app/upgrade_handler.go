package app

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	currencymoduletypes "github.com/fulldivevr/imversed/x/currency/types"
)

func (app ImversedApp) setUpgradeHandler(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		"v2.1",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			err := cfg.RegisterMigration("nft", 2, app.NFTKeeper.MigrationAddOracleUrl)
			if err != nil {
				return nil, err
			}

			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == "v2.2" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{currencymoduletypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	app.UpgradeKeeper.SetUpgradeHandler(
		"v2.3",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			// does nothing, is necessary only because ob nft version was changed
			err := cfg.RegisterMigration("nft", 3, func(ctx sdk.Context) error { return nil })
			if err != nil {
				return nil, err
			}
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
}
