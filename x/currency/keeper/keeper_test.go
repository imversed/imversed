package keeper_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/evmos/ethermint/encoding"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/testutil"
	"github.com/imversed/imversed/x/currency/keeper"
	"github.com/imversed/imversed/x/currency/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	denom    = "testcoin"
	denom2   = "testcoin2"
	address  = testutil.CreateTestAddrs(1)[0]
	address2 = testutil.CreateTestAddrs(2)[1]
)

type KeeperSuite struct {
	suite.Suite

	app         app.ImversedApp
	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	keeper      keeper.Keeper

	queryClient types.QueryClient
}

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}

func (suite *KeeperSuite) SetupTest() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	homePath := filepath.Join(userHomeDir, ".simapp")

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	db := dbm.NewMemDB()
	//encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	app := *app.NewImversedApp(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encCfg, EmptyAppOptions{})

	suite.app = app
	suite.ctx = app.BaseApp.NewUncachedContext(false, tmproto.Header{})
	suite.legacyAmino = app.LegacyAmino()

	suite.keeper = app.CurrencyKeeper

	suite.keeper.SetParams(suite.ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CurrencyKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *KeeperSuite) TestIssue() {
	// should bind coin denom to address
	suite.NoError(suite.keeper.Issue(suite.ctx, denom, address, ""))

	// error if denom exists
	suite.Error(suite.keeper.Issue(suite.ctx, denom, address2, ""))
}

func (suite *KeeperSuite) TestMint() {
	suite.NoError(suite.keeper.Issue(suite.ctx, denom, address, ""))

	// should mint user owned coin
	suite.NoError(suite.keeper.Mint(suite.ctx, sdk.NewCoin(denom, sdk.NewInt(1000000000)), address))

	// should fail if user don't own coin's denom
	suite.Error(suite.keeper.Mint(suite.ctx, sdk.NewCoin(denom, sdk.NewInt(1000000000)), address2))
}
