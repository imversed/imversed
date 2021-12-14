package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fulldivevr/imversed/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateCurrency_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCurrency
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCurrency{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCurrency{
				Owner: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateCurrency_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCurrency
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCurrency{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCurrency{
				Owner: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteCurrency_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCurrency
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCurrency{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCurrency{
				Owner: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
