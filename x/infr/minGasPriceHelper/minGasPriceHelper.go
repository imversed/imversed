package minGasPriceHelper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
)

type minGasPriceHelper struct {
	bApp     *baseapp.BaseApp
	setter   func(gasPricesStr string) func(app *baseapp.BaseApp)
	GasPrice string
}

func (g minGasPriceHelper) Set(minGasPrices string) {
	fmt.Printf("\nsetting mingas %s", minGasPrices)
	bAppSetter := g.setter(minGasPrices)
	bAppSetter(g.bApp)
	g.GasPrice = minGasPrices
}

func (g minGasPriceHelper) Get() string {
	return g.GasPrice
}

func Create(f func(gasPricesStr string) func(app *baseapp.BaseApp), initGasPrice string) {
	init := minGasPriceHelper{}
	init.setter = f
	init.GasPrice = initGasPrice
	Helper = init
}

func SetBaseApp(bApp *baseapp.BaseApp) {
	Helper.bApp = bApp
}

var (
	Helper minGasPriceHelper
)
