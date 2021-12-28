package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fulldivevr/imversed/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMint_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMint{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMint{
				Sender: sample.AccAddress(),
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

func TestMsgIssue_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgIssue
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgIssue{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgIssue{
				Sender: sample.AccAddress(),
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
