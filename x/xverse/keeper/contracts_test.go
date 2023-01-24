package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/encoding"
	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/x/xverse/keeper"
	"github.com/imversed/imversed/x/xverse/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

type KeeperSuite struct {
	suite.Suite

	app    app.ImversedApp
	ctx    sdk.Context
	keeper keeper.Keeper

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

	encCfg := encoding.MakeConfig(app.ModuleBasics)
	imvApp := *app.NewImversedApp(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encCfg, EmptyAppOptions{})

	suite.app = imvApp
	suite.ctx = imvApp.BaseApp.NewUncachedContext(false, tmproto.Header{})

	suite.keeper = imvApp.VerseKeeper

	suite.keeper.SetParams(suite.ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, imvApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, imvApp.VerseKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *KeeperSuite) TestGetAllContract() {
	for i := 0; i < 10; i++ {
		verse := types.Verse{
			Name: strconv.Itoa(i),
		}
		suite.NoError(suite.keeper.SetVerse(suite.ctx, verse))
	}

	for i := 0; i < 15; i++ {
		contract := types.Contract{
			Hash: strconv.Itoa(i),
		}
		suite.NoError(suite.keeper.SetContract(suite.ctx, contract))
	}

	verses := suite.keeper.GetAllVerses(suite.ctx)

	contracts := suite.keeper.GetAllContracts(suite.ctx)

	suite.True(len(verses) == 10)

	suite.True(len(contracts) == 15)
}
