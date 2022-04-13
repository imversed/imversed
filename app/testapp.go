package app

import (
	"encoding/json"
	"github.com/imversed/imversed/encoding"
	"io"
	"os"
	"path/filepath"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"
)

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func CreateTestApp() *ImversedApp {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	homePath := filepath.Join(userHomeDir, ".simapp")

	logger := log.NewTMLogger(log.NewSyncWriter(io.Discard))
	db := dbm.NewMemDB()
	//encoding := cosmoscmd.MakeEncodingConfig(ModuleBasics)

	testapp := New(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encoding.MakeConfig(ModuleBasics), EmptyAppOptions{})

	genesisState := NewDefaultGenesisState(testapp.AppCodec())
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

	return testapp
}
