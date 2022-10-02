package baseAppHelper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type baseAppHelper struct {
	bApp *baseapp.BaseApp
}

func Create(bApp *baseapp.BaseApp) {
	init := baseAppHelper{}
	init.bApp = bApp
	Helper = init
}

func (g baseAppHelper) GetHeightedContext(height int64) sdk.Context {
	ms := g.bApp.CommitMultiStore()
	cms, _ := ms.CacheMultiStoreWithVersion(height)
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithBlockHeight(height)
}

var (
	Helper baseAppHelper
)
