package v3_test

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/encoding"
	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/x/xverse/keeper"
	v2 "github.com/imversed/imversed/x/xverse/migrations/v2/types"
	v3 "github.com/imversed/imversed/x/xverse/migrations/v3"
	"github.com/imversed/imversed/x/xverse/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func (suite *KeeperSuite) TestMigration() {
	// random contracts
	contracts := []string{
		"0x3230554546f7c1090C82c7af35107A664cde431C",
		"0xD6cC98edDA7d4829533A5F3702e4a752EeAd4316",
		"0xeC2Da9E7c24352402a16f3eE38170f597Ec3aa95",
		"0x1B3e945f6E1bF9f89C65fd84dDaB13DF30533b25",
		"0xca290375A288e2F6a41FDa24df397Da13145cede",
		"0xfC4c66a8E14e21B43612F4781A2D554fe672131a",
		"0x8c5839b5eF19c5475150000392f5A6C885D194c2",
		"0x2e9735973d0c26bbD28b47C25A41624497Cb5F58",
		"0x9Fe4DBF1f846dAF20124c49cC4879C102e77A21c",
		"0x3CF178364E3d58cD6a257081c451bf9Cd366A83D",
	}

	skey := suite.app.GetKey(types.ModuleName)
	cdc := suite.app.AppCodec()
	store := prefix.NewStore(suite.ctx.KVStore(skey), types.KeyPrefixVerse)

	// Phase 1: fill store old verses
	for i := 0; i < 10; i++ {
		verse := v2.Verse{
			Name:           strconv.Itoa(i),
			SmartContracts: []string{contracts[i]},
		}
		b := cdc.MustMarshal(&verse)
		store.Set(types.VerseKey(verse.Name), b)
	}

	verses := suite.keeper.GetAllVerses(suite.ctx)

	suite.True(len(verses) == 10)

	// Phase 2: migration
	suite.NoError(v3.MigrateStore(suite.ctx, skey, cdc))

	// Phase 3: check store

	store = prefix.NewStore(suite.ctx.KVStore(skey), types.KeyPrefixContract)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		var contract types.Contract
		cdc.MustUnmarshal(iterator.Value(), &contract)

		index, err := strconv.Atoi(contract.Verse)
		suite.NoError(err)

		// key has '/' in the end, so contract hash checked by its length, and no check for last '/' symbol
		for i, v := range []byte(strings.ToLower(contracts[index])) {
			suite.True(v == key[i])
		}
	}
}

func (suite *KeeperSuite) TestOwnerVersesMappingMigration() {
	skey := suite.app.GetKey(types.ModuleName)
	cdc := suite.app.AppCodec()
	store := prefix.NewStore(suite.ctx.KVStore(skey), types.KeyPrefixVerse)
	// Phase 1: fill store old verses
	for i := 0; i < 9; i++ {
		verse := v2.Verse{
			Name: strconv.Itoa(i),
			// mock 3 only owners
			Owner:          strconv.Itoa(i % 3),
			SmartContracts: nil,
		}
		b := cdc.MustMarshal(&verse)
		store.Set(types.VerseKey(verse.Name), b)
	}

	suite.NoError(v3.MigrateStore(suite.ctx, skey, cdc))

	for i := 0; i < 3; i++ {
		mappingStore := prefix.NewStore(suite.ctx.KVStore(skey), types.KeyPrefixCreatorToVerse(strconv.Itoa(i)))
		iterator := sdk.KVStorePrefixIterator(mappingStore, []byte{})

		j := 0
		for ; iterator.Valid(); iterator.Next() {
			verse := string(iterator.Key())
			name, err := strconv.Atoi(string(verse))
			suite.NoError(err)
			suite.True(name < 9)
			j++
		}
		suite.True(j == 3)
		iterator.Close()
	}
}
