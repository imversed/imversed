package types

import (
	"fmt"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CurrencyList: []Currency{},
		// this line is used by starport scaffolding # genesis/types/default
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in currency
	currencyIndexMap := make(map[string]struct{})

	for _, elem := range gs.CurrencyList {
		index := string(CurrencyKey(elem.Denom))
		if _, ok := currencyIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for currency")
		}
		currencyIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
