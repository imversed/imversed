package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"net/url"
	"regexp"
)

var (
	_ sdk.Msg = &MsgCreateVerse{}
	_ sdk.Msg = &MsgAddAssetToVerse{}
)

const (
	TypeMsgCreateVerse          = "create_verse"
	TypeMsgAddAssetToVerse      = "add_asset_to_verse"
	TypeMsgRenameVerse          = "rename_verse"
	TypeMsgRemoveAssetFromVerse = "remove_asset_from_verse"
	TypeMsgAddOracleToVerse     = "add_oracle_to_verse"
	TypeAuthorizeKeyToVerse     = "authorize_key_to_verse"
	TypeDeauthorizeKeyToVerse   = "deauthorize_key_to_verse"
	TypeUpdateIcon              = "update_icon"
	TypeUpdateDescription       = "update_description"
	NameLength                  = 50
	NameRegex                   = "^[A-Za-z0-9]*$"
	DescriptionLength           = 2000
	ContractType                = "contract"
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
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid asset type: '%s'", msg.AssetType)
	}

	if msg.Sender != msg.VerseCreator {
		_, err := sdk.AccAddressFromBech32(msg.VerseCreator)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid verseCreator address")
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

	verseCreator, err := sdk.AccAddressFromBech32(msg.VerseCreator)
	if err != nil {
		return nil
	}

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

// Route should return the name of the module
func (msg *MsgRenameVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgRenameVerse) Type() string { return TypeMsgRenameVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgRenameVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	if len(msg.VerseNewName) > NameLength {
		return fmt.Errorf("verse's name can't be more than 50 chars")
	} else if re := regexp.MustCompile(NameRegex); re.FindString(msg.VerseNewName) == "" {
		return fmt.Errorf("characters in the verse's name should only [A-Z][a-z][0-9]")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRenameVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgRenameVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	verseCreator, err := sdk.AccAddressFromBech32(msg.VerseCreator)
	if err != nil {
		return nil
	}

	if verseCreator.String() != addr.String() {
		signers = append(signers, verseCreator)
	}

	return signers
}

// Route should return the name of the module
func (msg *MsgRemoveAssetFromVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgRemoveAssetFromVerse) Type() string { return TypeMsgRemoveAssetFromVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgRemoveAssetFromVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	switch msg.AssetType {
	case ContractType:
		if !common.IsHexAddress(msg.AssetId) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract hex address '%s'", msg.AssetId)
		}
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid asset type: '%s'", msg.AssetType)
	}

	if msg.Sender != msg.VerseCreator {
		_, err := sdk.AccAddressFromBech32(msg.VerseCreator)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid verseCreator address")
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRemoveAssetFromVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgRemoveAssetFromVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	verseCreator, err := sdk.AccAddressFromBech32(msg.VerseCreator)
	if err != nil {
		return nil
	}

	if verseCreator.String() != addr.String() {
		signers = append(signers, verseCreator)
	}

	return signers
}

// Route should return the name of the module
func (msg *MsgAddOracleToVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgAddOracleToVerse) Type() string { return TypeMsgAddOracleToVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgAddOracleToVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Oracle)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid oracle address")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgAddOracleToVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgAddOracleToVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	return signers
}

// Route should return the name of the module
func (msg *MsgAuthorizeKeyToVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgAuthorizeKeyToVerse) Type() string { return TypeAuthorizeKeyToVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgAuthorizeKeyToVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid key address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgAuthorizeKeyToVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgAuthorizeKeyToVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	return signers
}

// Route should return the name of the module
func (msg *MsgDeauthorizeKeyToVerse) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgDeauthorizeKeyToVerse) Type() string { return TypeDeauthorizeKeyToVerse }

// ValidateBasic runs stateless checks on the message
func (msg *MsgDeauthorizeKeyToVerse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid key address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgDeauthorizeKeyToVerse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgDeauthorizeKeyToVerse) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	return signers
}

// Route should return the name of the module
func (msg *MsgUpdateVerseIcon) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgUpdateVerseIcon) Type() string { return TypeUpdateIcon }

// ValidateBasic runs stateless checks on the message
func (msg *MsgUpdateVerseIcon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUpdateVerseIcon) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgUpdateVerseIcon) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	return signers
}

// Route should return the name of the module
func (msg *MsgUpdateVerseDescription) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgUpdateVerseDescription) Type() string { return TypeUpdateDescription }

// ValidateBasic runs stateless checks on the message
func (msg *MsgUpdateVerseDescription) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUpdateVerseDescription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgUpdateVerseDescription) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	signers = append(signers, addr)

	return signers
}
