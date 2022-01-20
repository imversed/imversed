package app

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	currencymoduletypes "github.com/fulldivevr/imversed/x/currency/types"
)

const upgradeName = "v2.2"

func (app ImversedApp) setUpgradeHandler(cfg module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgradeName,
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

	if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{currencymoduletypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
