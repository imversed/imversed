package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	imversed "github.com/tharsis/ethermint/types"
)

// constants
const (
	ProposalTypeRegisterERC20    string = "RegisterERC20"
	ProposalTypeToggleTokenRelay string = "ToggleTokenRelay" // #nosec
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterERC20Proposal{}
	_ govtypes.Content = &ToggleTokenRelayProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterERC20)
	govtypes.RegisterProposalType(ProposalTypeToggleTokenRelay)
	govtypes.RegisterProposalTypeCodec(&RegisterERC20Proposal{}, "erc20/RegisterERC20Proposal")
	govtypes.RegisterProposalTypeCodec(&ToggleTokenRelayProposal{}, "erc20/ToggleTokenRelayProposal")
}

// CreateDenomDescription generates a string with the coin description
func CreateDenomDescription(address string) string {
	return fmt.Sprintf("Cosmos coin token representation of %s", address)
}

// CreateDenom generates a string the module name plus the address to avoid conflicts with names staring with a number
func CreateDenom(address string) string {
	return fmt.Sprintf("%s/%s", ModuleName, address)
}

// ValidateErc20Denom checks if a denom is a valid erc20/
// denomination
func ValidateErc20Denom(denom string) error {
	denomSplit := strings.SplitN(denom, "/", 2)

	if len(denomSplit) != 2 || denomSplit[0] != ModuleName {
		return fmt.Errorf("invalid denom. %s denomination should be prefixed with the format 'erc20/", denom)
	}

	return imversed.ValidateAddress(denomSplit[1])
}

// NewRegisterERC20Proposal returns new instance of RegisterERC20Proposal
func NewRegisterERC20Proposal(title, description, erc20Addr string) govtypes.Content {
	return &RegisterERC20Proposal{
		Title:        title,
		Description:  description,
		Erc20Address: erc20Addr,
	}
}

// ProposalRoute returns router key for this proposal
func (*RegisterERC20Proposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*RegisterERC20Proposal) ProposalType() string {
	return ProposalTypeRegisterERC20
}

// ValidateBasic performs a stateless check of the proposal fields
func (rtbp *RegisterERC20Proposal) ValidateBasic() error {
	if err := imversed.ValidateAddress(rtbp.Erc20Address); err != nil {
		return sdkerrors.Wrap(err, "ERC20 address")
	}
	return govtypes.ValidateAbstract(rtbp)
}

// NewToggleTokenRelayProposal returns new instance of ToggleTokenRelayProposal
func NewToggleTokenRelayProposal(title, description string, token string) govtypes.Content {
	return &ToggleTokenRelayProposal{
		Title:       title,
		Description: description,
		Token:       token,
	}
}

// ProposalRoute returns router key for this proposal
func (*ToggleTokenRelayProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*ToggleTokenRelayProposal) ProposalType() string {
	return ProposalTypeToggleTokenRelay
}

// ValidateBasic performs a stateless check of the proposal fields
func (etrp *ToggleTokenRelayProposal) ValidateBasic() error {
	// check if the token is a hex address, if not, check if it is a valid SDK
	// denom
	if err := imversed.ValidateAddress(etrp.Token); err != nil {
		if err := sdk.ValidateDenom(etrp.Token); err != nil {
			return err
		}
	}

	return govtypes.ValidateAbstract(etrp)
}
