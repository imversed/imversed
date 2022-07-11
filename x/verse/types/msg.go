package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateVerse{}
)

const (
	TypeMsgCreateVerse = "create_verse"
)

// NewMsgCreateVerse creates a new instance of MsgCreateVerse
func NewMsgCreateVerse(name string, sender sdk.AccAddress) *MsgCreateVerse { // nolint: interfacer
	return &MsgCreateVerse{
		Sender: sender.String(),
		Name:   name,
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
