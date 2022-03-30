package keeper_test

import (
	gocontext "context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fulldivevr/imversed/x/currency/types"
)

func (suite *KeeperSuite) TestCurrencyAllQuery() {
	suite.NoError(suite.keeper.Issue(suite.ctx, denom, address, ""))
	suite.NoError(suite.keeper.Issue(suite.ctx, denom2, address2, ""))

	response, err := suite.queryClient.CurrencyAll(gocontext.Background(), &types.QueryAllCurrencyRequest{})
	suite.NoError(err)
	suite.Equal(uint64(2), response.Pagination.Total)
}

func (suite *KeeperSuite) TestCurrencyQuery() {
	suite.NoError(suite.keeper.Issue(suite.ctx, denom, address, ""))
	suite.NoError(suite.keeper.Issue(suite.ctx, denom2, address2, ""))

	response, err := suite.queryClient.Currency(gocontext.Background(), &types.QueryGetCurrencyRequest{
		Denom: denom,
	})
	suite.NoError(err)
	owner, err := sdk.AccAddressFromBech32(response.Currency.Owner)
	suite.NoError(err)
	suite.Equal(address, owner)
}

func (suite *KeeperSuite) TestParamsQuery() {
	response, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.NoError(err)
	suite.Equal(types.DefaultTxMintCurrencyCost, response.Params.TxMintCurrencyCost)
}
