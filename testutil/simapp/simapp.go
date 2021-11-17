package simapp
//
//import (
//	"github.com/cosmos/cosmos-sdk/baseapp"
//	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/server"
//	storetypes "github.com/cosmos/cosmos-sdk/store/types"
//	"github.com/cosmos/cosmos-sdk/testutil/testdata"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	"github.com/cosmos/cosmos-sdk/version"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
//	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
//	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
//	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
//	"github.com/cosmos/cosmos-sdk/x/authz"
//	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
//	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//	"github.com/cosmos/cosmos-sdk/x/capability"
//	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
//	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
//	"github.com/cosmos/cosmos-sdk/x/crisis"
//	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
//	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
//	distr "github.com/cosmos/cosmos-sdk/x/distribution"
//	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
//	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
//	"github.com/cosmos/cosmos-sdk/x/evidence"
//	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
//	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
//	"github.com/cosmos/cosmos-sdk/x/feegrant"
//	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
//	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
//	"github.com/cosmos/cosmos-sdk/x/genutil"
//	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
//	"github.com/cosmos/cosmos-sdk/x/gov"
//	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
//	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
//	"github.com/cosmos/cosmos-sdk/x/mint"
//	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
//	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
//	"github.com/cosmos/cosmos-sdk/x/params"
//	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
//	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
//	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
//	"github.com/cosmos/cosmos-sdk/x/slashing"
//	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
//	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
//	"github.com/cosmos/cosmos-sdk/x/staking"
//	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
//	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
//	"github.com/cosmos/cosmos-sdk/x/upgrade"
//	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
//	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
//	nftkeeper "github.com/fulldivevr/imversed/x/nft/keeper"
//	nfttypes "github.com/fulldivevr/imversed/x/nft/types"
//	"github.com/spf13/cast"
//	tmos "github.com/tendermint/tendermint/libs/os"
//	"time"
//
//	"github.com/cosmos/cosmos-sdk/simapp"
//	"github.com/tendermint/spm/cosmoscmd"
//	abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/libs/log"
//	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
//	tmtypes "github.com/tendermint/tendermint/types"
//	tmdb "github.com/tendermint/tm-db"
//
//	"github.com/fulldivevr/imversed/app"
//)
//
//type SimApp struct {
//	*baseapp.BaseApp
//	appCodec          codec.Codec
//	legacyAmino       *codec.LegacyAmino
//	//interfaceRegistry types.InterfaceRegistry
//	//msgSvcRouter      *authmiddleware.MsgServiceRouter
//	//legacyRouter      sdk.Router
//
//	invCheckPeriod uint
//
//	// keys to access the substores
//	keys    map[string]*storetypes.KVStoreKey
//	tkeys   map[string]*storetypes.TransientStoreKey
//	memKeys map[string]*storetypes.MemoryStoreKey
//
//	// keepers
//	AccountKeeper    authkeeper.AccountKeeper
//	BankKeeper       bankkeeper.Keeper
//	CapabilityKeeper *capabilitykeeper.Keeper
//	StakingKeeper    stakingkeeper.Keeper
//	SlashingKeeper   slashingkeeper.Keeper
//	MintKeeper       mintkeeper.Keeper
//	DistrKeeper      distrkeeper.Keeper
//	GovKeeper        govkeeper.Keeper
//	CrisisKeeper     crisiskeeper.Keeper
//	UpgradeKeeper    upgradekeeper.Keeper
//	ParamsKeeper     paramskeeper.Keeper
//	AuthzKeeper      authzkeeper.Keeper
//	EvidenceKeeper   evidencekeeper.Keeper
//	FeeGrantKeeper   feegrantkeeper.Keeper
//	NFTKeeper        nftkeeper.Keeper
//
//	// the module manager
//	mm *module.Manager
//
//	// simulation manager
//	sm *module.SimulationManager
//
//	// module configurator
//	configurator module.Configurator
//}
//
////func init() {
////	userHomeDir, err := os.UserHomeDir()
////	if err != nil {
////		panic(err)
////	}
////	//DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
////}
//
//// New creates application instance with in-memory database and disabled logging.
//func New(dir string) cosmoscmd.App {
//	db := tmdb.NewMemDB()
//	logger := log.NewNopLogger()
//
//	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
//
//	a := app.New(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
//		simapp.EmptyAppOptions{})
//
//	// InitChain updates deliverState which is required when app.NewContext is called
//	a.InitChain(abci.RequestInitChain{
//		ConsensusParams: defaultConsensusParams,
//		AppStateBytes:   []byte("{}"),
//	})
//
//	return a
//}
//
//NewSimApp returns a reference to an initialized SimApp.
//func NewSimApp(
//) *SimApp {
//
//	db := tmdb.NewMemDB()
//	logger := log.NewNopLogger()
//	encodingConfig := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
//
//	appCodec := encodingConfig.Marshaler
//	legacyAmino := encodingConfig.Amino
//	interfaceRegistry := encodingConfig.InterfaceRegistry
//
//	bApp := baseapp.NewBaseApp("imversed-simapp", logger, db, encodingConfig.TxConfig.TxDecoder())
//
//	bApp.SetVersion(version.Version)
//	bApp.SetInterfaceRegistry(interfaceRegistry)
//
//	keys := sdk.NewKVStoreKeys(
//		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
//		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
//		govtypes.StoreKey, paramstypes.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
//		evidencetypes.StoreKey, capabilitytypes.StoreKey,
//		authzkeeper.StoreKey, nfttypes.StoreKey,
//	)
//	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
//	// NOTE: The testingkey is just mounted for testing purposes. Actual applications should
//	// not include this key.
//	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey, "testingkey")
//
//	//// configure state listening capabilities using AppOptions
//	//// we are doing nothing with the returned streamingServices and waitGroup in this case
//	//if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, keys); err != nil {
//	//	tmos.Exit(err.Error())
//	//}
//
//	app := &SimApp{
//		BaseApp:           bApp,
//		appCodec:          appCodec,
//		//invCheckPeriod:    invCheckPeriod,
//		keys:              keys,
//		tkeys:             tkeys,
//		memKeys:           memKeys,
//	}
//
//	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
//
//	// set the BaseApp's parameter store
//	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))
//
//	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
//	// Applications that wish to enforce statically created ScopedKeepers should call `Seal` after creating
//	// their scoped modules in `NewApp` with `ScopeToModule`
//	app.CapabilityKeeper.Seal()
//
//	// add keepers
//	app.AccountKeeper = authkeeper.NewAccountKeeper(
//		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix,
//	)
//	app.BankKeeper = bankkeeper.NewBaseKeeper(
//		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
//	)
//	stakingKeeper := stakingkeeper.NewKeeper(
//		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
//	)
//	app.MintKeeper = mintkeeper.NewKeeper(
//		appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), &stakingKeeper,
//		app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
//	)
//	app.DistrKeeper = distrkeeper.NewKeeper(
//		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
//		&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
//	)
//	app.SlashingKeeper = slashingkeeper.NewKeeper(
//		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
//	)
//	app.CrisisKeeper = crisiskeeper.NewKeeper(
//		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
//	)
//
//	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
//	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)
//
//	// register the staking hooks
//	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
//	app.StakingKeeper = *stakingKeeper.SetHooks(
//		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
//	)
//
//	app.AuthzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], appCodec, app.msgSvcRouter)
//
//	// register the proposal types
//	govRouter := govtypes.NewRouter()
//	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
//		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
//		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
//		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper))
//	govKeeper := govkeeper.NewKeeper(
//		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
//		&stakingKeeper, govRouter,
//	)
//
//	app.GovKeeper = *govKeeper.SetHooks(
//		govtypes.NewMultiGovHooks(
//			// register the governance hooks
//		),
//	)
//
//	app.NFTKeeper = nftkeeper.NewKeeper(appCodec, keys[nfttypes.StoreKey])
//
//	// create evidence keeper with router
//	evidenceKeeper := evidencekeeper.NewKeeper(
//		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
//	)
//	// If evidence needs to be handled for the app, set routes in router here and seal
//	app.EvidenceKeeper = *evidenceKeeper
//
//	/****  Module Options ****/
//
//	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
//	// we prefer to be more strict in what arguments the modules expect.
//	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
//
//	// NOTE: Any module instantiated in the module manager that is later modified
//	// must be passed by reference here.
//	app.mm = module.NewManager(
//		genutil.NewAppModule(
//			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
//			encodingConfig.TxConfig,
//		),
//		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
//		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
//		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
//		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
//		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
//		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
//		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
//		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
//		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
//		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
//		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
//		upgrade.NewAppModule(app.UpgradeKeeper),
//		evidence.NewAppModule(app.EvidenceKeeper),
//		params.NewAppModule(app.ParamsKeeper),
//		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
//		nftmodule.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
//	)
//
//	// During begin block slashing happens after distr.BeginBlocker so that
//	// there is nothing left over in the validator fee pool, so as to keep the
//	// CanWithdrawInvariant invariant.
//	// NOTE: staking module is required if HistoricalEntries param > 0
//	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
//	app.mm.SetOrderBeginBlockers(
//		upgradetypes.ModuleName, capabilitytypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
//		evidencetypes.ModuleName, stakingtypes.ModuleName,
//	)
//	app.mm.SetOrderEndBlockers(crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName)
//
//	// NOTE: The genutils module must occur after staking so that pools are
//	// properly initialized with tokens from genesis accounts.
//	// NOTE: Capability module must occur first so that it can initialize any capabilities
//	// so that other modules that want to create or claim capabilities afterwards in InitChain
//	// can do so safely.
//	app.mm.SetOrderInitGenesis(
//		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName, distrtypes.ModuleName, stakingtypes.ModuleName,
//		slashingtypes.ModuleName, govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
//		genutiltypes.ModuleName, evidencetypes.ModuleName, authz.ModuleName,
//		feegrant.ModuleName, nft.ModuleName,
//	)
//
//	app.mm.RegisterInvariants(&app.CrisisKeeper)
//	app.mm.RegisterRoutes(app.legacyRouter, app.QueryRouter(), encodingConfig.Amino)
//	app.configurator = module.NewConfigurator(app.appCodec, app.msgSvcRouter, app.GRPCQueryRouter())
//	app.mm.RegisterServices(app.configurator)
//
//	// add test gRPC service for testing gRPC queries in isolation
//	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})
//
//	// create the simulation manager and define the order of the modules for deterministic simulations
//	//
//	// NOTE: this is not required apps that don't use the simulator for fuzz testing
//	// transactions
//	app.sm = module.NewSimulationManager(
//		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
//		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
//		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
//		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
//		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
//		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
//		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
//		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
//		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
//		params.NewAppModule(app.ParamsKeeper),
//		evidence.NewAppModule(app.EvidenceKeeper),
//		//authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
//	)
//
//	app.sm.RegisterStoreDecoders()
//
//	// initialize stores
//	app.MountKVStores(keys)
//	app.MountTransientStores(tkeys)
//	app.MountMemoryStores(memKeys)
//
//	// initialize BaseApp
//	app.SetInitChainer(app.InitChainer)
//	app.SetBeginBlocker(app.BeginBlocker)
//	app.SetEndBlocker(app.EndBlocker)
//	app.setTxHandler(encodingConfig.TxConfig, cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents)))
//
//	if loadLatest {
//		if err := app.LoadLatestVersion(); err != nil {
//			tmos.Exit(err.Error())
//		}
//	}
//
//	return app
//}
//
//// initParamsKeeper init params keeper and its subspaces
//func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
//	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)
//
//	paramsKeeper.Subspace(authtypes.ModuleName)
//	paramsKeeper.Subspace(banktypes.ModuleName)
//	paramsKeeper.Subspace(stakingtypes.ModuleName)
//	paramsKeeper.Subspace(minttypes.ModuleName)
//	paramsKeeper.Subspace(distrtypes.ModuleName)
//	paramsKeeper.Subspace(slashingtypes.ModuleName)
//	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
//	paramsKeeper.Subspace(crisistypes.ModuleName)
//
//	return paramsKeeper
//}
//
//var defaultConsensusParams = &abci.ConsensusParams{
//	Block: &abci.BlockParams{
//		MaxBytes: 200000,
//		MaxGas:   2000000,
//	},
//	Evidence: &tmproto.EvidenceParams{
//		MaxAgeNumBlocks: 302400,
//		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
//		MaxBytes:        10000,
//	},
//	Validator: &tmproto.ValidatorParams{
//		PubKeyTypes: []string{
//			tmtypes.ABCIPubKeyTypeEd25519,
//		},
//	},
//}
