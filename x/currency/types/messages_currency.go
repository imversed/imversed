package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCurrency{}

func NewMsgCreateCurrency(
	owner string,
	denom string,

) *MsgCreateCurrency {
	return &MsgCreateCurrency{
		Owner: owner,
		Denom: denom,
	}
}

func (msg *MsgCreateCurrency) Route() string {
	return RouterKey
}

func (msg *MsgCreateCurrency) Type() string {
	return "CreateCurrency"
}

func (msg *MsgCreateCurrency) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateCurrency) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCurrency) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCurrency{}

func NewMsgUpdateCurrency(
	owner string,
	denom string,

) *MsgUpdateCurrency {
	return &MsgUpdateCurrency{
		Owner: owner,
		Denom: denom,
	}
}

func (msg *MsgUpdateCurrency) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCurrency) Type() string {
	return "UpdateCurrency"
}

func (msg *MsgUpdateCurrency) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateCurrency) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCurrency) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCurrency{}

func NewMsgDeleteCurrency(
	owner string,
	denom string,

) *MsgDeleteCurrency {
	return &MsgDeleteCurrency{
		Owner: owner,
		Denom: denom,
	}
}
func (msg *MsgDeleteCurrency) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCurrency) Type() string {
	return "DeleteCurrency"
}

func (msg *MsgDeleteCurrency) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgDeleteCurrency) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCurrency) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
