package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
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

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.9",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.10",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.11",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.12",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
}
