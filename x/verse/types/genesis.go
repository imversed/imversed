package types

import "fmt"

// DefaultGenesisState sets default genesis state and
// default params and chain config values.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]Verse{}, DefaultParams())
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(verseList []Verse, params Params) *GenesisState {
	return &GenesisState{
		VerseList: verseList,
		Params:    params,
	}
}

func (gs GenesisState) Validate() error {
	verseIndexMap := make(map[string]struct{})

	for _, elem := range gs.VerseList {
		index := string(VerseKey(elem.Name))
		if _, ok := verseIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for verse")
		}
		verseIndexMap[index] = struct{}{}
	}
	return nil
}
