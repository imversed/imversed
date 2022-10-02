package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"net/url"
)

var (
	_ sdk.Msg = &MsgCreateVerse{}
	_ sdk.Msg = &MsgAddAssetToVerse{}
)

const (
	TypeMsgCreateVerse     = "create_verse"
	TypeMsgAddAssetToVerse = "add_asset_to_verse"
	NameLength             = 50
	NameRegex              = "^[A-Za-z0-9]*$"
	DescriptionLength      = 2000
	ContractType           = "contract"
)

// NewMsgCreateVerse creates a new instance of MsgCreateVerse
func NewMsgCreateVerse(sender sdk.AccAddress, desc string, icon string) *MsgCreateVerse { // nolint: interfacer
	return &MsgCreateVerse{
		Sender:      sender.String(),
		Description: desc,
		Icon:        icon,
	}
}

// Route should return the name of the module
func (msg *MsgCreateVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgCreateVerse) Type() string { return TypeMsgCreateVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	if len(msg.Description) > DescriptionLength {
		return fmt.Errorf("verse's description can't be more than 2000 chars")
	}

	if msg.Icon == "" {
		return nil
	} else if _, err := url.Parse(msg.Icon); err != nil {
		return sdkerrors.Wrap(err, "invalid icon: icon must be valid url")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateVerse) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
}

// NewMsgAddAssetToVerse creates a new instance of MsgAddAssetToVerse
func NewMsgAddAssetToVerse(sender sdk.AccAddress, desc string, icon string) *MsgAddAssetToVerse { // nolint: interfacer
	return &MsgAddAssetToVerse{
		Sender: sender.String(),
	}
}

// Route should return the name of the module
func (msg *MsgAddAssetToVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgAddAssetToVerse) Type() string { return TypeMsgAddAssetToVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgAddAssetToVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	switch msg.AssetType {
	case ContractType:
		if !common.IsHexAddress(msg.AssetId) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract hex address '%s'", msg.AssetId)
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgAddAssetToVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgAddAssetToVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	verseCreator, _ := sdk.AccAddressFromBech32(msg.VerseCreator)

	if verseCreator.String() != addr.String() {
		signers = append(signers, verseCreator)
	}

	switch msg.AssetType {
	case ContractType:
		assetCreator, _ := sdk.AccAddressFromHexUnsafe(msg.AssetCreator[2:])
		if assetCreator.String() != addr.String() {
			signers = append(signers, assetCreator)
		}
	}

	return signers
}
