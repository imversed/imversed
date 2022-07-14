package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"net/url"
	"regexp"
)

var (
	_ sdk.Msg = &MsgCreateVerse{}
)

const (
	TypeMsgCreateVerse = "create_verse"
	NameLength         = 50
	NameRegex          = "^[A-Za-z0-9]*$"
	DescriptionLength  = 2000
)

// NewMsgCreateVerse creates a new instance of MsgCreateVerse
func NewMsgCreateVerse(sender sdk.AccAddress, name string, desc string, icon string) *MsgCreateVerse { // nolint: interfacer
	return &MsgCreateVerse{
		Sender:      sender.String(),
		Name:        name,
		Description: desc,
		Icon:        icon,
	}
}

// Route should return the name of the module
func (msg MsgCreateVerse) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateVerse) Type() string { return TypeMsgCreateVerse }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	if len(msg.Name) > NameLength {
		return fmt.Errorf("verse's name can't be more than 50 chars")
	} else if re := regexp.MustCompile(NameRegex); re.FindString(msg.Icon) == "" {
		return fmt.Errorf("characters in the verse's name should only [A-Z][a-z][0-9]")
	}

	if len(msg.Description) > DescriptionLength {
		return fmt.Errorf("verse's description can't be more than 2000 chars")
	}

	if msg.Icon == "" {
		return nil
	} else if _, err := url.Parse(msg.Icon); err != nil {
		return sdkerrors.Wrap(err, "invalid icon: icon must be valid url or should not be specified")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateVerse) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
}
