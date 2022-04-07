package nft_test

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/fulldivevr/imversed/app"
	"github.com/fulldivevr/imversed/x/nft"
	"github.com/fulldivevr/imversed/x/nft/keeper"
	"github.com/fulldivevr/imversed/x/nft/simulation"
	"github.com/fulldivevr/imversed/x/nft/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spm/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
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

func TestNftInitSuite(t *testing.T) {
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

	suite.keeper = app.NFTKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.NFTKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *InitGenesisTestSuite) TestInitGenesis() {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	collections := types.NewCollections(
		types.NewCollection(
			types.Denom{
				Id:      "doggos",
				Name:    "doggos",
				Schema:  "",
				Creator: "",
				Symbol:  "dog",
			},
			types.NFTs{},
		),
		types.NewCollection(
			types.Denom{
				Id:      "kitties",
				Name:    "kitties",
				Schema:  "",
				Creator: "",
				Symbol:  "kit",
			},
			types.NFTs{}),
	)
	for _, acc := range simState.Accounts {
		// 10% of accounts own an NFT
		if simState.Rand.Intn(100) < 10 {
			baseNFT := types.NewBaseNFT(
				RandnNFTID(simState.Rand, types.MinDenomLen, types.MaxDenomLen), // id
				simtypes.RandStringOfLength(simState.Rand, 10),
				acc.Address,
				simtypes.RandStringOfLength(simState.Rand, 45), // tokenURI
				simtypes.RandStringOfLength(simState.Rand, 10),
			)

			// 50% doggos and 50% kitties
			if simState.Rand.Intn(100) < 50 {
				collections[0].Denom.Creator = baseNFT.Owner
				collections[0] = collections[0].AddNFT(baseNFT)
			} else {
				collections[1].Denom.Creator = baseNFT.Owner
				collections[1] = collections[1].AddNFT(baseNFT)
			}
		}
	}

	st := types.NewGenesisState(collections)

	nft.InitGenesis(suite.ctx, suite.keeper, *st)
	exp := nft.ExportGenesis(suite.ctx, suite.keeper)

	suite.Require().Equal(*st, *exp)
}

func RandnNFTID(r *rand.Rand, min, max int) string {
	n := simtypes.RandIntBetween(r, min, max)
	id := simtypes.RandStringOfLength(r, n)
	return strings.ToLower(id)
}
