package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

var (
	ParamStoreKeyMinGasPrices = []byte("MinGasPrices")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		MinGasPrices: "",
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyMinGasPrices, &p.MinGasPrices, func(i interface{}) error { return nil }),
	}
}

func (p Params) Validate() error { return nil }
