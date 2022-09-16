package types

import (
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeChangeMinGasPrices string = "ChangeMinGasPrices"
)

var (
	_ govv1.Content = &ChangeMinGasPricesProposal{}
)

func init() {
	govv1.RegisterProposalType(ProposalTypeChangeMinGasPrices)
}

// NewChangeMinGasPricesProposal returns new instance of ChangeMinGasPricesProposal
func NewChangeMinGasPricesProposal(title, description string, minGasPrices string) govv1.Content {
	return &ChangeMinGasPricesProposal{
		Title:        title,
		Description:  description,
		MinGasPrices: minGasPrices,
	}
}

// ProposalRoute returns router key for this proposal
func (*ChangeMinGasPricesProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*ChangeMinGasPricesProposal) ProposalType() string {
	return ProposalTypeChangeMinGasPrices
}

// ValidateBasic performs a stateless check of the proposal fields
func (etrp *ChangeMinGasPricesProposal) ValidateBasic() error {
	//TODO: add validate
	return govv1.ValidateAbstract(etrp)
}
