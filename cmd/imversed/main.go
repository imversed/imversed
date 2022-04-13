package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/imversed/imversed/app"
	cmdcfg "github.com/imversed/imversed/cmd/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := NewRootCmd()

	//rootCmd, _ := cosmoscmd.NewRootCmd(
	//	app.Name,
	//	app.AccountAddressPrefix,
	//	app.DefaultNodeHome,
	//	app.Name,
	//	app.ModuleBasics,
	//	AppCtorAdapter,
	//	// this line is used by starport scaffolding # root/arguments
	//)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

//func AppCtorAdapter(
//	logger log.Logger,
//	db dbm.DB,
//	traceStore io.Writer,
//	loadLatest bool,
//	skipUpgradeHeights map[int64]bool,
//	homePath string,
//	invCheckPeriod uint,
//	encodingConfig cosmoscmd.EncodingConfig,
//	appOpts servertypes.AppOptions,
//	baseAppOptions ...func(*baseapp.BaseApp),
//) cosmoscmd.App {
//	return app.New(logger, db, traceStore, loadLatest, skipUpgradeHeights, homePath, invCheckPeriod, encodingConfig, appOpts, baseAppOptions...)
//}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
