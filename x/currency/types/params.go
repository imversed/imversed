package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyTxMintCurrencyCost = []byte("TxMintCurrencyCost")
	// TODO: Determine the default value
	DefaultTxMintCurrencyCost uint64 = 1000000
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	txMintCurrencyCost uint64,
) Params {
	return Params{
		TxMintCurrencyCost: txMintCurrencyCost,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultTxMintCurrencyCost,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTxMintCurrencyCost, &p.TxMintCurrencyCost, validateTxMintCurrencyCost),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateTxMintCurrencyCost(p.TxMintCurrencyCost); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateTxMintCurrencyCost validates the TxMintCurrencyCost param
func validateTxMintCurrencyCost(v interface{}) error {
	txMintCurrencyCost, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = txMintCurrencyCost

	return nil
}
