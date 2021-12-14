package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCurrencyMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CurrencyKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	owner := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCurrency{Owner: owner,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateCurrency(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCurrency(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Owner, rst.Owner)
	}
}

func TestCurrencyMsgServerUpdate(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCurrency
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateCurrency{Owner: owner,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCurrency{Owner: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCurrency{Owner: owner,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CurrencyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCurrency{Owner: owner,
				Denom: strconv.Itoa(0),
			}
			_, err := srv.CreateCurrency(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCurrency(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCurrency(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Owner, rst.Owner)
			}
		})
	}
}

func TestCurrencyMsgServerDelete(t *testing.T) {
	owner := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCurrency
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteCurrency{Owner: owner,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCurrency{Owner: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCurrency{Owner: owner,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CurrencyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCurrency(wctx, &types.MsgCreateCurrency{Owner: owner,
				Denom: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteCurrency(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCurrency(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}
