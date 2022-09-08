package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
	imversed "github.com/evmos/ethermint/types"
	"strings"
)

var (
	_ sdk.Msg = &MsgConvertCoin{}
	_ sdk.Msg = &MsgConvertERC20{}
	_ sdk.Msg = &MsgUpdateTokenPairERC20{}
	_ sdk.Msg = &MsgRegisterCoin{}
	_ sdk.Msg = &MsgToggleTokenRelay{}
	_ sdk.Msg = &MsgRegisterERC20{}
)

const (
	TypeMsgConvertCoin          = "convert_coin"
	TypeMsgConvertERC20         = "convert_ERC20"
	TypeMsgRegisterCoin         = "register_coin"
	TypeMsgRegisterERC20        = "register_erc20"
	TypeMsgUpdateTokenPairERC20 = "update_token_pair_erc20"
	TypeMsgToggleTokenRelay     = "toggle_token_relay"
)

// NewMsgConvertCoin creates a new instance of MsgConvertCoin
func NewMsgConvertCoin(coin sdk.Coin, receiver common.Address, sender sdk.AccAddress) *MsgConvertCoin { // nolint: interfacer
	return &MsgConvertCoin{
		Coin:     coin,
		Receiver: receiver.Hex(),
		Sender:   sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertCoin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertCoin) Type() string { return TypeMsgConvertCoin }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertCoin) ValidateBasic() error {
	if err := ValidateErc20Denom(msg.Coin.Denom); err != nil {
		if err := ibctransfertypes.ValidateIBCDenom(msg.Coin.Denom); err != nil {
			return err
		}
	}

	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(msg.Receiver) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver hex address %s", msg.Receiver)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgConvertCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
}

// NewMsgConvertERC20 creates a new instance of MsgConvertERC20
func NewMsgConvertERC20(amount sdk.Int, receiver sdk.AccAddress, contract, sender common.Address) *MsgConvertERC20 { // nolint: interfacer
	return &MsgConvertERC20{
		ContractAddress: contract.String(),
		Amount:          amount,
		Receiver:        receiver.String(),
		Sender:          sender.Hex(),
	}
}

// Route should return the name of the module
func (msg MsgConvertERC20) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertERC20) Type() string { return TypeMsgConvertERC20 }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertERC20) ValidateBasic() error {
	if !common.IsHexAddress(msg.ContractAddress) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract hex address '%s'", msg.ContractAddress)
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid reciver address")
	}
	if !common.IsHexAddress(msg.Sender) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender hex address %s", msg.Sender)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgConvertERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertERC20) GetSigners() []sdk.AccAddress {
	addr := common.HexToAddress(msg.Sender)
	return []sdk.AccAddress{addr.Bytes()}
}

// TODO: reimplement in tests needs
// NewMsgRegisterCoin creates a new instance of MsgRegisterCoin
func NewMsgRegisterCoin() *MsgRegisterCoin { // nolint: interfacer

	//return &MsgRegisterCoin{
	//	Metadata:
	//}
	return nil
}

// Route should return the name of the module
func (msg MsgRegisterCoin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterCoin) Type() string { return TypeMsgRegisterCoin }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr.Bytes()}
}

// NewMsgUpdateTokenPairERC20 creates a new instance of MsgUpdateTokenPairERC20
func NewMsgUpdateTokenPairERC20(erc20Address common.Address, newErc20Address common.Address, sender sdk.AccAddress) *MsgUpdateTokenPairERC20 { // nolint: interfacer
	return &MsgUpdateTokenPairERC20{
		Erc20Address:    erc20Address.String(),
		NewErc20Address: newErc20Address.String(),
		Sender:          sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgUpdateTokenPairERC20) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateTokenPairERC20) Type() string { return TypeMsgUpdateTokenPairERC20 }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateTokenPairERC20) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	if !common.IsHexAddress(msg.Erc20Address) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract hex address '%s'", msg.Erc20Address)
	}

	if !common.IsHexAddress(msg.NewErc20Address) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract hex address '%s'", msg.NewErc20Address)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUpdateTokenPairERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateTokenPairERC20) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
}

// TODO: reimplement in tests needs
// NewMsgRegisterERC20 creates a new instance of MsgRegisterERC20
func NewMsgRegisterERC20(address string, sender string) *MsgRegisterERC20 { // nolint: interfacer

	return &MsgRegisterERC20{
		Erc20Address: address,
		Sender:       sender,
	}
}

// Route should return the name of the module
func (msg MsgRegisterERC20) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterERC20) Type() string { return TypeMsgRegisterERC20 }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterERC20) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterERC20) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr.Bytes()}
}

// NewMsgToggleTokenRelay creates a new instance of MsgToggleTokenRelay
func NewMsgToggleTokenRelay(token string, sender sdk.AccAddress) *MsgToggleTokenRelay { // nolint: interfacer
	return &MsgToggleTokenRelay{
		Token:  token,
		Sender: sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgToggleTokenRelay) Route() string { return RouterKey }

// Type should return the action
func (msg MsgToggleTokenRelay) Type() string { return TypeMsgToggleTokenRelay }

// ValidateBasic runs stateless checks on the message
func (msg MsgToggleTokenRelay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgToggleTokenRelay) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgToggleTokenRelay) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
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
