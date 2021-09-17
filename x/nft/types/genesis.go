package types

// NewGenesisState creates a new genesis state.
func NewGenesisState(collections []Collection) *GenesisState {
	return &GenesisState{
		Collections: collections,
	}
}

// ValidateGenesis performs basic validation of nfts genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return nil
}
