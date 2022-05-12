package keeper_test

import (
	"github.com/tharsis/ethermint/encoding"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"

	imversedapp "github.com/imversed/imversed/app"
	"github.com/imversed/imversed/testutil"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"os"
	"path/filepath"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftkeeper "github.com/imversed/imversed/x/nft/keeper"
	nft "github.com/imversed/imversed/x/nft/types"
)

var (
	denomID      = "denomid"
	denomNm      = "denomnm"
	denomSymbol  = "denomSymbol"
	schema       = "{a:a,b:b}"
	denomID2     = "denomid2"
	denomNm2     = "denom2nm"
	denomSymbol2 = "denomSymbol2"

	tokenID  = "tokenid"
	tokenID2 = "tokenid2"
	tokenID3 = "tokenid3"

	tokenNm  = "tokennm"
	tokenNm2 = "tokennm2"
	tokenNm3 = "tokennm3"

	denomID3     = "denomid3"
	denomNm3     = "denom3nm"
	denomSymbol3 = "denomSymbol3"

	address   = testutil.CreateTestAddrs(1)[0]
	address2  = testutil.CreateTestAddrs(2)[1]
	address3  = testutil.CreateTestAddrs(3)[2]
	tokenURI  = "https://google.com/token-1.json"
	tokenURI2 = "https://google.com/token-2.json"
	tokenData = "{a:a,b:b}"

	oracleUrl = "https://www.yArtViq.ru/pnSwlld"
)

type KeeperSuite struct {
	suite.Suite

	app         imversedapp.ImversedApp
	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	keeper      nftkeeper.Keeper

	queryClient nft.QueryClient
}

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func (suite *KeeperSuite) SetupTest() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	homePath := filepath.Join(userHomeDir, ".simapp")

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	db := dbm.NewMemDB()

	app := *imversedapp.NewImversedApp(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encoding.MakeConfig(imversedapp.ModuleBasics), EmptyAppOptions{})

	suite.app = app
	suite.ctx = app.BaseApp.NewUncachedContext(false, tmproto.Header{})
	suite.legacyAmino = app.LegacyAmino()

	suite.keeper = app.NFTKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	nft.RegisterQueryServer(queryHelper, app.NFTKeeper)
	suite.queryClient = nft.NewQueryClient(queryHelper)

	err = suite.keeper.IssueDenom(suite.ctx, denomID, denomNm, schema, denomSymbol, address, false, false, oracleUrl)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.IssueDenom(suite.ctx, denomID2, denomNm2, schema, denomSymbol2, address, false, false, oracleUrl)
	suite.NoError(err)
	//
	err = suite.keeper.IssueDenom(suite.ctx, denomID3, denomNm3, schema, denomSymbol3, address3, true, true, oracleUrl)
	suite.NoError(err)

	// collections should equal 3
	collections := suite.keeper.GetCollections(suite.ctx)
	suite.NotEmpty(collections)
	suite.Equal(len(collections), 3)
}

func (suite *KeeperSuite) TestUpdateDenom() {
	denomE := nft.NewDenom(denomID, denomNm, schema, denomSymbol, address, false, false, "")
	err := suite.keeper.UpdateDenom(suite.ctx, denomE)
	suite.NoError(err)
	denom, ok := suite.keeper.GetDenom(suite.ctx, denomID)
	suite.Equal(ok, true)
	suite.Equal(denom, denomE)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}

func (suite *KeeperSuite) TestMintNFT() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenData, address)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection exists
	err = suite.keeper.MintNFT(suite.ctx, denomID, tokenID2, tokenNm2, tokenURI, tokenData, address)
	suite.NoError(err)

	// MintNFT should fail when owner not equal to denom owner
	err = suite.keeper.MintNFT(suite.ctx, denomID3, tokenID3, tokenNm3, tokenURI, tokenData, address)
	suite.Error(err)

	// MintNFT shouldn't fail when owner equal to denom owner
	err = suite.keeper.MintNFT(suite.ctx, denomID3, tokenID3, tokenNm3, tokenURI, tokenData, address3)
	suite.NoError(err)

}

func (suite *KeeperSuite) TestUpdateNFT() {
	// EditNFT should fail when NFT doesn't exists
	err := suite.keeper.EditNFT(suite.ctx, denomID, tokenID, tokenNm3, tokenURI, tokenData, address)
	suite.Error(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenData, address)
	suite.NoError(err)

	// EditNFT should fail when NFT doesn't exists
	err = suite.keeper.EditNFT(suite.ctx, denomID, tokenID2, tokenNm2, tokenURI, tokenData, address)
	suite.Error(err)

	// EditNFT shouldn't fail when NFT exists
	err = suite.keeper.EditNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI2, tokenData, address)
	suite.NoError(err)

	// EditNFT should fail when NFT failed to authorize
	err = suite.keeper.EditNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI2, tokenData, address2)
	suite.Error(err)

	// GetNFT should get the NFT with new tokenURI
	receivedNFT, err := suite.keeper.GetNFT(suite.ctx, denomID, tokenID)
	suite.NoError(err)
	suite.Equal(receivedNFT.GetURI(), tokenURI2)

	// EditNFT shouldn't fail when NFT exists
	err = suite.keeper.EditNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI2, tokenData, address2)
	suite.Error(err)

	err = suite.keeper.MintNFT(suite.ctx, denomID3, denomID3, tokenID3, tokenURI, tokenData, address3)
	suite.NoError(err)

	// EditNFT should fail if updateRestricted equal to true, nobody can update the NFT under this denom
	err = suite.keeper.EditNFT(suite.ctx, denomID3, denomID3, tokenID3, tokenURI, tokenData, address3)
	suite.Error(err)
}

func (suite *KeeperSuite) TestTransferOwner() {

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenData, address)
	suite.NoError(err)

	// invalid owner
	err = suite.keeper.TransferOwner(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenData, address2, address3)
	suite.Error(err)

	// right
	err = suite.keeper.TransferOwner(suite.ctx, denomID, tokenID, tokenNm2, tokenURI2, tokenData, address, address2)
	suite.NoError(err)

	nft, err := suite.keeper.GetNFT(suite.ctx, denomID, tokenID)
	suite.NoError(err)
	suite.Equal(tokenURI2, nft.GetURI())
}

func (suite *KeeperSuite) TestTransferDenom() {

	// invalid owner
	err := suite.keeper.TransferDenomOwner(suite.ctx, denomID, address3, address)
	suite.Error(err)

	// right
	err = suite.keeper.TransferDenomOwner(suite.ctx, denomID, address, address3)
	suite.NoError(err)

	denom, _ := suite.keeper.GetDenom(suite.ctx, denomID)

	// denom.Creator should equal to address3 after transfer
	suite.Equal(denom.Creator, address3.String())
}

func (suite *KeeperSuite) TestBurnNFT() {
	// MintNFT should not fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenData, address)
	suite.NoError(err)

	// BurnNFT should fail when NFT doesn't exist but collection does exist
	err = suite.keeper.BurnNFT(suite.ctx, denomID, tokenID, address)
	suite.NoError(err)

	// NFT should no longer exist
	isNFT := suite.keeper.HasNFT(suite.ctx, denomID, tokenID)
	suite.False(isNFT)

	//msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	//suite.False(fail, msg)
}
