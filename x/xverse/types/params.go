package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

var (
	KeyTxRenameVerseCost            = []byte("TxRenameVerseCost")
	DefaultTxRenameVerseCost uint64 = 1 * 1000000
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(
	txRenameVerseCost uint64,
) Params {
	return Params{
		TxRenameVerseCost: txRenameVerseCost,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultTxRenameVerseCost,
	)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTxRenameVerseCost, &p.TxRenameVerseCost, validateTxRenameVerseCost),
	}
}

func (p Params) Validate() error {
	if err := validateTxRenameVerseCost(p.TxRenameVerseCost); err != nil {
		return err
	}
	return nil
}

func validateTxRenameVerseCost(v interface{}) error {
	_, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}
