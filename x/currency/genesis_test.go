package currency_test

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/app"
	"github.com/fulldivevr/imversed/x/currency"
	"github.com/fulldivevr/imversed/x/currency/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

type InitGenesisTestSuite struct {
	suite.Suite

	app         app.ImversedApp
	ctx         sdk.Context
	queryClient types.QueryClient
	legacyAmino *codec.LegacyAmino
	keeper      keeper.Keeper
}

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(InitGenesisTestSuite))
}

func (suite *InitGenesisTestSuite) SetupTest() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	homePath := filepath.Join(userHomeDir, ".simapp")

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	db := dbm.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	app := *app.New(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encoding, EmptyAppOptions{})

	suite.app = app
	suite.ctx = app.BaseApp.NewUncachedContext(false, tmproto.Header{})
	suite.legacyAmino = app.LegacyAmino()

	suite.keeper = app.CurrencyKeeper

	suite.keeper.SetParams(suite.ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CurrencyKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *InitGenesisTestSuite) TestInitGenesis() {
	g := types.DefaultGenesis()
	for i := 0; i < 2; i++ {
		g.CurrencyList = append(g.CurrencyList, types.Currency{
			Denom: strconv.Itoa(i),
		})
	}

	currency.InitGenesis(suite.ctx, suite.keeper, *g)
	exp := currency.ExportGenesis(suite.ctx, suite.keeper)

	suite.Require().Equal(*g, *exp)
}
