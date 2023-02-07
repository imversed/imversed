package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	infrtypes "github.com/imversed/imversed/x/infr/types"
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

			// re-create infr module account
			_, perms := app.AccountKeeper.GetModuleAddressAndPermissions(infrtypes.ModuleName)
			macc := types.NewEmptyModuleAccount(infrtypes.ModuleName, perms...)
			maccI := (app.AccountKeeper.NewAccount(ctx, macc)).(types.ModuleAccountI) // set the account number
			app.AccountKeeper.SetModuleAccount(ctx, maccI)

			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.13",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.14",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
	app.UpgradeKeeper.SetUpgradeHandler(
		"v3.15",
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, cfg, vm)
		},
	)
}
