package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	TypeMsgIssueDenom    = "issue_denom"
	TypeMsgUpdateDenom   = "update_denom"
	TypeMsgTransferNFT   = "transfer_nft"
	TypeMsgEditNFT       = "edit_nft"
	TypeMsgMintNFT       = "mint_nft"
	TypeMsgBurnNFT       = "burn_nft"
	TypeMsgTransferDenom = "transfer_denom"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
	_ sdk.Msg = &MsgUpdateDenom{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
	_ sdk.Msg = &MsgTransferDenom{}
)

// NewMsgIssueDenom is a constructor function for MsgIssueDenom
func NewMsgIssueDenom(denomID, denomName, schema, sender, symbol string, mintRestricted, updateRestricted bool, oracleUrl string) *MsgIssueDenom {
	return &MsgIssueDenom{
		Sender:           sender,
		Id:               denomID,
		Name:             denomName,
		Schema:           schema,
		Symbol:           symbol,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		OracleUrl: 		  oracleUrl,
	}
}
// NewMsgUpdateDenom is a constructor function for MsgUpdateDenom
func NewMsgUpdateDenom(denomID, denomName, schema, sender string, mintRestricted, updateRestricted bool, oracleUrl string) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Id:               denomID,
		Name:             denomName,
		Schema:           schema,
		Sender:           sender,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		OracleUrl: 		  oracleUrl,
	}
}

// Route Implements Msg
func (msg MsgIssueDenom) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgIssueDenom) Type() string { return TypeMsgIssueDenom }

// ValidateBasic Implements Msg.
func (msg MsgIssueDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.Id)
}

// ValidateBasic Implements Msg.
func (msg MsgUpdateDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSigners Implements Msg.
func (msg MsgUpdateDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(
	tokenID, denomID, tokenName, tokenURI, tokenData, sender, recipient string,
) *MsgTransferNFT {
	return &MsgTransferNFT{
		Id:        tokenID,
		DenomId:   denomID,
		Name:      tokenName,
		URI:       tokenURI,
		Data:      tokenData,
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgTransferNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return TypeMsgTransferNFT }

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgEditNFT is a constructor function for MsgSetName
func NewMsgEditNFT(
	tokenID, denomID, tokenName, tokenURI, tokenData, sender string,
) *MsgEditNFT {
	return &MsgEditNFT{
		Id:      tokenID,
		DenomId: denomID,
		Name:    tokenName,
		URI:     tokenURI,
		Data:    tokenData,
		Sender:  sender,
	}
}

// Route Implements Msg
func (msg MsgEditNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFT) Type() string { return TypeMsgEditNFT }

// ValidateBasic Implements Msg.
func (msg MsgEditNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if err := ValidateTokenURI(msg.URI); err != nil {
		return err
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(
	tokenID, denomID, tokenName, tokenURI, tokenData, sender, recipient string,
) *MsgMintNFT {
	return &MsgMintNFT{
		Id:        tokenID,
		DenomId:   denomID,
		Name:      tokenName,
		URI:       tokenURI,
		Data:      tokenData,
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return TypeMsgMintNFT }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receipt address (%s)", err)
	}
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}
	if err := ValidateKeywords(msg.DenomId); err != nil {
		return err
	}
	if err := ValidateTokenURI(msg.URI); err != nil {
		return err
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(sender, tokenID, denomID string) *MsgBurnNFT {
	return &MsgBurnNFT{
		Sender:  sender,
		Id:      tokenID,
		DenomId: denomID,
	}
}

// Route Implements Msg
func (msg MsgBurnNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return TypeMsgBurnNFT }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}
	return ValidateTokenID(msg.Id)
}

// GetSignBytes Implements Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgTransferDenom is a constructor function for msgTransferDenom
func NewMsgTransferDenom(denomId, sender, recipient string) *MsgTransferDenom {
	return &MsgTransferDenom{
		Id:        denomId,
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgTransferDenom) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferDenom) Type() string { return TypeMsgTransferDenom }

// ValidateBasic Implements Msg.
func (msg MsgTransferDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgTransferDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
