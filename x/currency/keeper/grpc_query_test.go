package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCurrencyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCurrency(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCurrencyRequest
		response *types.QueryGetCurrencyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCurrencyRequest{
				Denom: msgs[0].Denom,
			},
			response: &types.QueryGetCurrencyResponse{Currency: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCurrencyRequest{
				Denom: msgs[1].Denom,
			},
			response: &types.QueryGetCurrencyResponse{Currency: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCurrencyRequest{
				Denom: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Currency(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCurrencyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCurrency(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCurrencyRequest {
		return &types.QueryAllCurrencyRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CurrencyAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Currency), step)
			require.Subset(t, msgs, resp.Currency)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CurrencyAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Currency), step)
			require.Subset(t, msgs, resp.Currency)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CurrencyAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CurrencyAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.CurrencyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
