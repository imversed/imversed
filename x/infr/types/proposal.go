package types

import govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

const (
	ProposalTypeChangeMinGasPrices string = "ChangeMinGasPrices"
)

var (
	_ govtypes.Content = &ChangeMinGasPricesProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeChangeMinGasPrices)
	govtypes.RegisterProposalTypeCodec(&ChangeMinGasPricesProposal{}, "infr/ChangeMinGasPricesProposal")
}

// NewChangeMinGasPricesProposal returns new instance of ChangeMinGasPricesProposal
func NewChangeMinGasPricesProposal(title, description string, minGasPrices string) govtypes.Content {
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
	return govtypes.ValidateAbstract(etrp)
}
