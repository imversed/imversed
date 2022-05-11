package app

import (
	"encoding/json"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tharsis/ethermint/encoding"
	"os"

	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"
)

type EmptyAppOptions struct{}

func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

func CreateTestApp() *ImversedApp {

	homePath := DefaultNodeHome

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	db := dbm.NewMemDB()

	testapp := NewImversedApp(logger, db, nil, true,
		map[int64]bool{}, homePath, 0, encoding.MakeConfig(ModuleBasics), EmptyAppOptions{})

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		panic(err)
	}

	testapp.InitChain(
		abci.RequestInitChain{
			ChainId:       "imversed_1234-1",
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	return testapp
}
