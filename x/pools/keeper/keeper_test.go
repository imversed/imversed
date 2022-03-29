package keeper_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/fulldivevr/imversed/app"

	"github.com/fulldivevr/imversed/x/pools/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

type KeeperTestSuite struct {
	suite.Suite

	app         app.ImversedApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	homePath := filepath.Join(userHomeDir, ".simapp")

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	db := dbm.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	testapp := *app.New(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encoding, EmptyAppOptions{})

	genesisState := app.NewDefaultGenesisState(testapp.AppCodec())
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	testapp.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	suite.app = testapp
	suite.ctx = testapp.BaseApp.NewUncachedContext(false, tmproto.Header{})

	suite.app.PoolsKeeper.SetParams(suite.ctx, types.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.PoolsKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

var (
	acc1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

func (suite *KeeperTestSuite) preparePoolWithPoolParams(poolParams types.PoolParams) uint64 {
	// Mint some assets to the accounts.
	for _, acc := range []sdk.AccAddress{acc1, acc2, acc3} {
		err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, acc, sdk.NewCoins(
			sdk.NewCoin("nimv", sdk.NewInt(10000000000)),
			sdk.NewCoin("foo", sdk.NewInt(10000000)),
			sdk.NewCoin("bar", sdk.NewInt(10000000)),
			sdk.NewCoin("baz", sdk.NewInt(10000000)),
		))
		if err != nil {
			panic(err)
		}
	}

	poolId, err := suite.app.PoolsKeeper.CreatePool(suite.ctx, acc1, poolParams, []types.PoolAsset{
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("foo", sdk.NewInt(5000000)),
		},
		{
			Weight: sdk.NewInt(200),
			Token:  sdk.NewCoin("bar", sdk.NewInt(5000000)),
		},
		{
			Weight: sdk.NewInt(300),
			Token:  sdk.NewCoin("baz", sdk.NewInt(5000000)),
		},
	})
	suite.NoError(err)
	return poolId
}

func (suite *KeeperTestSuite) preparePool() uint64 {
	poolId := suite.preparePoolWithPoolParams(types.PoolParams{
		SwapFee: sdk.NewDec(0),
		ExitFee: sdk.NewDec(0),
	})

	spotPrice, err := suite.app.PoolsKeeper.CalculateSpotPriceWithSwapFee(suite.ctx, poolId, "foo", "bar")
	suite.NoError(err)
	suite.Equal(sdk.NewDec(2).String(), spotPrice.String())
	spotPrice, err = suite.app.PoolsKeeper.CalculateSpotPriceWithSwapFee(suite.ctx, poolId, "bar", "baz")
	suite.NoError(err)
	suite.Equal(sdk.NewDecWithPrec(15, 1).String(), spotPrice.String())
	spotPrice, err = suite.app.PoolsKeeper.CalculateSpotPriceWithSwapFee(suite.ctx, poolId, "baz", "foo")
	suite.NoError(err)
	suite.Equal(sdk.NewDec(1).Quo(sdk.NewDec(3)).String(), spotPrice.String())

	return poolId
}
