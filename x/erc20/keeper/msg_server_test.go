package keeper_test

import (
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/imversed/imversed/x/erc20/types"
)

func (suite *KeeperTestSuite) TestConvertCoinNativeCoin() {
	testCases := []struct {
		name           string
		mint           int64
		burn           int64
		malleate       func(common.Address)
		expPass        bool
		selfdestructed bool
	}{
		{"ok - sufficient funds", 100, 10, func(common.Address) {}, true, false},
		{"ok - equal funds", 10, 10, func(common.Address) {}, true, false},
		{
			"ok - suicided contract",
			10,
			10,
			func(erc20 common.Address) {
				stateDb := suite.StateDB()
				ok := stateDb.Suicide(erc20)
				suite.Require().True(ok)
				suite.Require().NoError(stateDb.Commit())
			},
			true,
			true,
		},
		{"fail - insufficient funds", 0, 10, func(common.Address) {}, false, false},
		{
			"fail - minting disabled",
			100,
			10,
			func(common.Address) {
				params := types.DefaultParams()
				params.EnableErc20 = false
				suite.app.Erc20Keeper.SetParams(suite.ctx, params)
			},
			false,
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()
			metadata, pair := suite.setupRegisterCoin()
			suite.Require().NotNil(metadata)
			erc20 := pair.GetERC20Contract()
			tc.malleate(erc20)
			suite.Commit()

			ctx := sdk.WrapSDKContext(suite.ctx)
			coins := sdk.NewCoins(sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.mint)))
			sender := sdk.AccAddress(suite.address.Bytes())
			msg := types.NewMsgConvertCoin(
				sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.burn)),
				suite.address,
				sender,
			)

			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, sender, coins)

			res, err := suite.app.Erc20Keeper.ConvertCoin(ctx, msg)
			expRes := &types.MsgConvertCoinResponse{}
			suite.Commit()
			balance := suite.BalanceOf(common.HexToAddress(pair.Erc20Address), suite.address)
			cosmosBalance := suite.app.BankKeeper.GetBalance(suite.ctx, sender, metadata.Base)

			if tc.expPass {
				suite.Require().NoError(err, tc.name)

				acc := suite.app.EvmKeeper.GetAccountWithoutBalance(suite.ctx, erc20)
				if tc.selfdestructed {
					suite.Require().Nil(acc, "expected contract to be destroyed")
				} else {
					suite.Require().NotNil(acc)
				}

				if tc.selfdestructed || !acc.IsContract() {
					id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, erc20.String())
					_, found := suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
					suite.Require().False(found)
				} else {
					suite.Require().Equal(expRes, res)
					suite.Require().Equal(cosmosBalance.Amount.Int64(), sdk.NewInt(tc.mint-tc.burn).Int64())
					suite.Require().Equal(balance.(*big.Int).Int64(), big.NewInt(tc.burn).Int64())
				}
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
	suite.mintFeeCollector = false
}

func (suite *KeeperTestSuite) TestConvertERC20NativeCoin() {
	testCases := []struct {
		name      string
		mint      int64
		burn      int64
		reconvert int64
		expPass   bool
	}{
		{"ok - sufficient funds", 100, 10, 5, true},
		{"ok - equal funds", 10, 10, 10, true},
		{"fail - insufficient funds", 10, 1, 5, false},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()
			metadata, pair := suite.setupRegisterCoin()
			suite.Require().NotNil(metadata)
			suite.Require().NotNil(pair)

			// Precondition: Convert Coin to ERC20
			coins := sdk.NewCoins(sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.mint)))
			sender := sdk.AccAddress(suite.address.Bytes())
			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, sender, coins)
			msg := types.NewMsgConvertCoin(
				sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.burn)),
				suite.address,
				sender,
			)

			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.Erc20Keeper.ConvertCoin(ctx, msg)
			suite.Require().NoError(err, tc.name)
			suite.Commit()
			balance := suite.BalanceOf(common.HexToAddress(pair.Erc20Address), suite.address)
			cosmosBalance := suite.app.BankKeeper.GetBalance(suite.ctx, sender, metadata.Base)
			suite.Require().Equal(cosmosBalance.Amount.Int64(), sdk.NewInt(tc.mint-tc.burn).Int64())
			suite.Require().Equal(balance, big.NewInt(tc.burn))

			// Convert ERC20s back to Coins
			ctx = sdk.WrapSDKContext(suite.ctx)
			contractAddr := common.HexToAddress(pair.Erc20Address)
			msgConvertERC20 := types.NewMsgConvertERC20(
				sdk.NewInt(tc.reconvert),
				sender,
				contractAddr,
				suite.address,
			)

			res, err := suite.app.Erc20Keeper.ConvertERC20(ctx, msgConvertERC20)
			expRes := &types.MsgConvertERC20Response{}
			suite.Commit()
			balance = suite.BalanceOf(contractAddr, suite.address)
			cosmosBalance = suite.app.BankKeeper.GetBalance(suite.ctx, sender, pair.Denom)
			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().Equal(expRes, res)
				suite.Require().Equal(cosmosBalance.Amount.Int64(), sdk.NewInt(tc.mint-tc.burn+tc.reconvert).Int64())
				suite.Require().Equal(balance.(*big.Int).Int64(), big.NewInt(tc.burn-tc.reconvert).Int64())
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
	suite.mintFeeCollector = false
}

func (suite *KeeperTestSuite) TestConvertERC20NativeERC20() {
	var contractAddr common.Address

	testCases := []struct {
		name           string
		mint           int64
		transfer       int64
		malleate       func(common.Address)
		contractType   int
		expPass        bool
		selfdestructed bool
	}{
		{
			"ok - sufficient funds",
			100,
			10,
			func(common.Address) {},
			contractMinterBurner,
			true,
			false,
		},
		{
			"ok - equal funds",
			10,
			10,
			func(common.Address) {},
			contractMinterBurner,
			true,
			false,
		},
		{
			"ok - equal funds",
			10,
			10,
			func(common.Address) {},
			contractMinterBurner,
			true,
			false,
		},
		{
			"ok - suicided contract",
			10,
			10,
			func(erc20 common.Address) {
				stateDb := suite.StateDB()
				ok := stateDb.Suicide(erc20)
				suite.Require().True(ok)
				suite.Require().NoError(stateDb.Commit())
			},
			contractMinterBurner,
			true,
			true,
		},
		{
			"fail - insufficient funds - callEVM",
			0,
			10,
			func(common.Address) {},
			contractMinterBurner,
			false,
			false,
		},
		{
			"fail - minting disabled",
			100,
			10,
			func(common.Address) {
				params := types.DefaultParams()
				params.EnableErc20 = false
				suite.app.Erc20Keeper.SetParams(suite.ctx, params)
			},
			contractMinterBurner,
			false,
			false,
		},
		{
			"fail - direct balance manipulation contract",
			100,
			10,
			func(common.Address) {},
			contractDirectBalanceManipulation,
			false,
			false,
		},
		{
			"fail - delayed malicious contract",
			10,
			10,
			func(common.Address) {},
			contractMaliciousDelayed,
			false,
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()

			contractAddr = suite.setupRegisterERC20Pair(tc.contractType)

			tc.malleate(contractAddr)
			suite.Require().NotNil(contractAddr)
			suite.Commit()

			coinName := types.CreateDenom(contractAddr.String())
			sender := sdk.AccAddress(suite.address.Bytes())
			msg := types.NewMsgConvertERC20(
				sdk.NewInt(tc.transfer),
				sender,
				contractAddr,
				suite.address,
			)

			suite.MintERC20Token(contractAddr, suite.address, suite.address, big.NewInt(tc.mint))
			suite.Commit()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.app.Erc20Keeper.ConvertERC20(ctx, msg)

			expRes := &types.MsgConvertERC20Response{}
			suite.Commit()
			balance := suite.BalanceOf(contractAddr, suite.address)
			cosmosBalance := suite.app.BankKeeper.GetBalance(suite.ctx, sender, coinName)
			if tc.expPass {
				suite.Require().NoError(err, tc.name)

				acc := suite.app.EvmKeeper.GetAccountWithoutBalance(suite.ctx, contractAddr)
				if tc.selfdestructed {
					suite.Require().Nil(acc, "expected contract to be destroyed")
				} else {
					suite.Require().NotNil(acc)
				}

				if tc.selfdestructed || !acc.IsContract() {
					id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
					_, found := suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
					suite.Require().False(found)
				} else {
					suite.Require().Equal(expRes, res)
					suite.Require().Equal(cosmosBalance.Amount, sdk.NewInt(tc.transfer))
					suite.Require().Equal(balance.(*big.Int).Int64(), big.NewInt(tc.mint-tc.transfer).Int64())
				}
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
	suite.mintFeeCollector = false
}

func (suite *KeeperTestSuite) TestConvertCoinNativeERC20() {
	var contractAddr common.Address

	testCases := []struct {
		name         string
		mint         int64
		convert      int64
		malleate     func(common.Address)
		contractType int
		expPass      bool
	}{
		{
			"ok - sufficient funds",
			100,
			10,
			func(common.Address) {},
			contractMinterBurner,
			true,
		},
		{
			"ok - equal funds",
			100,
			100,
			func(common.Address) {},
			contractMinterBurner,
			true,
		},
		{
			"fail - insufficient funds",
			100,
			200,
			func(common.Address) {},
			contractMinterBurner,
			false,
		},
		{
			"fail - direct balance manipulation contract",
			100,
			10,
			func(common.Address) {},
			contractDirectBalanceManipulation,
			false,
		},
		{
			"fail - malicious delayed contract",
			100,
			10,
			func(common.Address) {},
			contractMaliciousDelayed,
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()
			contractAddr = suite.setupRegisterERC20Pair(tc.contractType)
			suite.Require().NotNil(contractAddr)

			id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
			pair, _ := suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
			coins := sdk.NewCoins(sdk.NewCoin(pair.Denom, sdk.NewInt(tc.mint)))
			coinName := types.CreateDenom(contractAddr.String())
			sender := sdk.AccAddress(suite.address.Bytes())

			// Precondition: Mint Coins to convert on sender account
			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, sender, coins)
			cosmosBalance := suite.app.BankKeeper.GetBalance(suite.ctx, sender, coinName)
			suite.Require().Equal(sdk.NewInt(tc.mint), cosmosBalance.Amount)

			// Precondition: Mint escrow tokens on module account
			suite.GrantERC20Token(contractAddr, suite.address, types.ModuleAddress, "MINTER_ROLE")
			suite.MintERC20Token(contractAddr, types.ModuleAddress, types.ModuleAddress, big.NewInt(tc.mint))
			tokenBalance := suite.BalanceOf(contractAddr, types.ModuleAddress)
			suite.Require().Equal(big.NewInt(tc.mint), tokenBalance)

			tc.malleate(contractAddr)
			suite.Commit()

			// Convert Coins back to ERC20s
			receiver := suite.address
			ctx := sdk.WrapSDKContext(suite.ctx)
			msg := types.NewMsgConvertCoin(
				sdk.NewCoin(coinName, sdk.NewInt(tc.convert)),
				receiver,
				sender,
			)
			res, err := suite.app.Erc20Keeper.ConvertCoin(ctx, msg)

			expRes := &types.MsgConvertCoinResponse{}
			suite.Commit()
			tokenBalance = suite.BalanceOf(contractAddr, suite.address)
			cosmosBalance = suite.app.BankKeeper.GetBalance(suite.ctx, sender, coinName)
			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().Equal(expRes, res)
				suite.Require().Equal(sdk.NewInt(tc.mint-tc.convert), cosmosBalance.Amount)
				suite.Require().Equal(big.NewInt(tc.convert), tokenBalance.(*big.Int))
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
	suite.mintFeeCollector = false
}

/*
func (suite *KeeperTestSuite) TestConvertNativeIBC() {
	suite.SetupTest()
	base := "ibc/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2"

	validMetadata := banktypes.Metadata{
		Description: "ATOM IBC voucher (channel 14)",
		Base:        base,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    base,
				Exponent: 0,
			},
		},
		Name:    "ATOM channel-14",
		Symbol:  "ibcATOM-14",
		Display: base,
	}

	err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(base, 1)})
	suite.Require().NoError(err)

	_, err = suite.app.Erc20Keeper.RegisterCoin(suite.ctx, validMetadata)
	suite.Require().NoError(err)
	suite.Commit()
}
*/
func (suite *KeeperTestSuite) TestWrongPairOwnerERC20NativeCoin() {
	testCases := []struct {
		name      string
		mint      int64
		burn      int64
		reconvert int64
		expPass   bool
	}{
		{"ok - sufficient funds", 100, 10, 5, true},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.mintFeeCollector = true
			suite.SetupTest()
			metadata, pair := suite.setupRegisterCoin()
			suite.Require().NotNil(metadata)
			suite.Require().NotNil(pair)

			// Precondition: Convert Coin to ERC20
			coins := sdk.NewCoins(sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.mint)))
			sender := sdk.AccAddress(suite.address.Bytes())
			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, sender, coins)
			msg := types.NewMsgConvertCoin(
				sdk.NewCoin(cosmosTokenBase, sdk.NewInt(tc.burn)),
				suite.address,
				sender,
			)

			pair.ContractOwner = types.OWNER_UNSPECIFIED
			suite.app.Erc20Keeper.SetTokenPair(suite.ctx, *pair)

			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.Erc20Keeper.ConvertCoin(ctx, msg)
			suite.Require().Error(err, tc.name)

			// Convert ERC20s back to Coins
			ctx = sdk.WrapSDKContext(suite.ctx)
			contractAddr := common.HexToAddress(pair.Erc20Address)
			msgConvertERC20 := types.NewMsgConvertERC20(
				sdk.NewInt(tc.reconvert),
				sender,
				contractAddr,
				suite.address,
			)

			_, err = suite.app.Erc20Keeper.ConvertERC20(ctx, msgConvertERC20)
			suite.Require().Error(err, tc.name)
		})
	}
}

func (suite KeeperTestSuite) TestUpdateTokenPairERC20() {
	var (
		contractAddr    common.Address
		pair            types.TokenPair
		metadata        banktypes.Metadata
		newContractAddr common.Address
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"token not registered",
			func() {
				contractAddr, err := suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
				suite.Commit()
				pair = types.NewTokenPair(contractAddr, cosmosTokenBase, true, types.OWNER_MODULE)
			},
			false,
		},
		{
			"token not registered - pair not found",
			func() {
				contractAddr, err := suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
				suite.Commit()
				pair = types.NewTokenPair(contractAddr, cosmosTokenBase, true, types.OWNER_MODULE)

				suite.app.Erc20Keeper.SetERC20Map(suite.ctx, common.HexToAddress(pair.Erc20Address), pair.GetID())
			},
			false,
		},
		{
			"token not registered - Metadata not found",
			func() {
				contractAddr, err := suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
				suite.Commit()
				pair = types.NewTokenPair(contractAddr, cosmosTokenBase, true, types.OWNER_MODULE)

				suite.app.Erc20Keeper.SetTokenPair(suite.ctx, pair)
				suite.app.Erc20Keeper.SetDenomMap(suite.ctx, pair.Denom, pair.GetID())
				suite.app.Erc20Keeper.SetERC20Map(suite.ctx, common.HexToAddress(pair.Erc20Address), pair.GetID())
			},
			false,
		},
		{
			"newErc20 not found",
			func() {
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				newContractAddr = common.Address{}
			},
			false,
		},
		{
			"empty denom units",
			func() {
				var found bool
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, found = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				suite.Require().True(found)
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, banktypes.Metadata{Base: pair.Denom})
				suite.Commit()

				// Deploy a new contract with the same values
				var err error
				newContractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"metadata ERC20 details mismatch",
			func() {
				var found bool
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, found = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				suite.Require().True(found)
				metadata := banktypes.Metadata{Base: pair.Denom, DenomUnits: []*banktypes.DenomUnit{{}}}
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, metadata)
				suite.Commit()

				// Deploy a new contract with the same values
				var err error
				newContractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"no denom unit with ERC20 name",
			func() {
				var found bool
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, found = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				suite.Require().True(found)
				metadata := banktypes.Metadata{Base: pair.Denom, Display: erc20Name, Description: types.CreateDenomDescription(contractAddr.String()), Symbol: erc20Symbol, DenomUnits: []*banktypes.DenomUnit{{}}}
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, metadata)
				suite.Commit()

				// Deploy a new contract with the same values
				var err error
				newContractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"denom unit and ERC20 decimals mismatch",
			func() {
				var found bool
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, found = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				suite.Require().True(found)
				metadata := banktypes.Metadata{Base: pair.Denom, Display: erc20Name, Description: types.CreateDenomDescription(contractAddr.String()), Symbol: erc20Symbol, DenomUnits: []*banktypes.DenomUnit{{Denom: erc20Name}}}
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, metadata)
				suite.Commit()

				// Deploy a new contract with the same values
				var err error
				newContractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
			},
			false,
		},
		{
			"ok",
			func() {
				var found bool
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, found = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				suite.Require().True(found)
				metadata := banktypes.Metadata{Base: pair.Denom, Display: erc20Name, Description: types.CreateDenomDescription(contractAddr.String()), Symbol: erc20Symbol, DenomUnits: []*banktypes.DenomUnit{{Denom: erc20Name, Exponent: 18}}}
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, metadata)
				suite.Commit()

				// Deploy a new contract with the same values
				var err error
				newContractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			tc.malleate()
			sender := sdk.AccAddress(suite.address.Bytes())
			msg := types.NewMsgUpdateTokenPairERC20(
				contractAddr,
				newContractAddr,
				sender,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)

			_, err := suite.app.Erc20Keeper.UpdateTokenPairERC20(ctx, msg)

			metadata, _ = suite.app.BankKeeper.GetDenomMetaData(suite.ctx, types.CreateDenom(contractAddr.String()))

			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().Equal(types.CreateDenomDescription(newContractAddr.String()), metadata.Description)
			} else {
				suite.Require().Error(err, tc.name)

			}
		})
	}
}
