package keeper_test

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	contractMinterBurner = iota + 1
	contractDirectBalanceManipulation
	contractMaliciousDelayed
)

const (
	erc20Name       = "Coin Token"
	erc20Symbol     = "CTKN"
	erc20Decimals   = uint8(18)
	cosmosTokenBase = "acoin"
	//cosmosTokenDisplay = "coin"
	//cosmosDecimals     = uint8(6)
	//defaultExponent    = uint32(18)
	//zeroExponent       = uint32(0)
)

func (suite *KeeperTestSuite) setupRegisterERC20Pair(contractType int) common.Address {
	suite.SetupTest()

	var contractAddr common.Address
	// Deploy contract
	switch contractType {
	case contractDirectBalanceManipulation:
		contractAddr = suite.DeployContractDirectBalanceManipulation(erc20Name, erc20Symbol)
	case contractMaliciousDelayed:
		contractAddr = suite.DeployContractMaliciousDelayed(erc20Name, erc20Symbol)
	default:
		contractAddr, _ = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
	}
	suite.Commit()

	_, err := suite.app.Erc20Keeper.RegisterERC20(suite.ctx, contractAddr)
	suite.Require().NoError(err)
	return contractAddr
}

//func (suite *KeeperTestSuite) setupRegisterCoin() (banktypes.Metadata, *types.TokenPair) {
//	suite.SetupTest()
//	validMetadata := banktypes.Metadata{
//		Description: "description of the token",
//		Base:        cosmosTokenBase,
//		// NOTE: Denom units MUST be increasing
//		DenomUnits: []*banktypes.DenomUnit{
//			{
//				Denom:    cosmosTokenBase,
//				Exponent: 0,
//			},
//			{
//				Denom:    cosmosTokenBase[1:],
//				Exponent: uint32(18),
//			},
//		},
//		Name:    cosmosTokenBase,
//		Symbol:  erc20Symbol,
//		Display: cosmosTokenBase,
//	}
//
//	err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(validMetadata.Base, 1)})
//	suite.Require().NoError(err)
//
//	//sender := sdk.AccAddress(suite.address.Bytes())
//	//pair := types.NewTokenPair(contractAddr, cosmosTokenBase, true, types.OWNER_MODULE,sender.String() )
//	msg := types.NewMsgRegisterCoin(
//		contractAddr,
//		newContractAddr,
//		tc.sender,
//	)
//	ctx := sdk.WrapSDKContext(suite.ctx)
//	suite.app.Erc20Keeper.RegisterCoin(ctx, msg)
//
//	createdPairId := suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, cosmosTokenBase)
//	createdPair, _ := suite.app.Erc20Keeper.GetTokenPair(suite.ctx, createdPairId)
//	suite.Require().NoError(err)
//	suite.Commit()
//	return validMetadata, &createdPair
//}

//func (suite KeeperTestSuite) TestRegisterCoin() {
//	metadata := banktypes.Metadata{
//		Description: "description",
//		Base:        cosmosTokenBase,
//		// NOTE: Denom units MUST be increasing
//		DenomUnits: []*banktypes.DenomUnit{
//			{
//				Denom:    cosmosTokenBase,
//				Exponent: 0,
//			},
//			{
//				Denom:    cosmosTokenDisplay,
//				Exponent: defaultExponent,
//			},
//		},
//		Name:    cosmosTokenBase,
//		Symbol:  erc20Symbol,
//		Display: cosmosTokenDisplay,
//	}
//
//	testCases := []struct {
//		name     string
//		malleate func()
//		expPass  bool
//	}{
//		{
//			"intrarelaying is disabled globally",
//			func() {
//				params := types.DefaultParams()
//				params.EnableErc20 = false
//				suite.app.Erc20Keeper.SetParams(suite.ctx, params)
//			},
//			false,
//		},
//		{
//			"denom already registered",
//			func() {
//				sender := sdk.AccAddress(suite.address.Bytes())
//
//				regPair := types.NewTokenPair(tests.GenerateAddress(), metadata.Base, true, types.OWNER_MODULE, sender.String())
//				suite.app.Erc20Keeper.SetDenomMap(suite.ctx, regPair.Denom, regPair.GetID())
//				suite.Commit()
//			},
//			false,
//		},
//		{
//			"token doesn't have supply",
//			func() {
//			},
//			false,
//		},
//		{
//			"metadata different that stored",
//			func() {
//				metadata.Base = cosmosTokenBase
//				validMetadata := banktypes.Metadata{
//					Description: "description",
//					Base:        cosmosTokenBase,
//					// NOTE: Denom units MUST be increasing
//					DenomUnits: []*banktypes.DenomUnit{
//						{
//							Denom:    cosmosTokenBase,
//							Exponent: 0,
//						},
//						{
//							Denom:    cosmosTokenDisplay,
//							Exponent: uint32(18),
//						},
//					},
//					Name:    erc20Name,
//					Symbol:  erc20Symbol,
//					Display: cosmosTokenDisplay,
//				}
//
//				err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(validMetadata.Base, 1)})
//				suite.Require().NoError(err)
//				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, validMetadata)
//			},
//			false,
//		}, /*
//			{
//				"evm denom registration - evm",
//				func() {
//					metadata.Base = "evm"
//					err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
//					suite.Require().NoError(err)
//				},
//				false,
//			},
//			{
//				"evm denom registration - evmos",
//				func() {
//					metadata.Base = "evmos"
//					err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
//					suite.Require().NoError(err)
//				},
//				false,
//			},
//			{
//				"evm denom registration - aevmos",
//				func() {
//					metadata.Base = "aevmos"
//					err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
//					suite.Require().NoError(err)
//				},
//				false,
//			},
//			{
//				"evm denom registration - wevmos",
//				func() {
//					metadata.Base = "wevmos"
//					err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
//					suite.Require().NoError(err)
//				},
//				false,
//			},*/
//		{
//			"ok",
//			func() {
//				metadata.Base = cosmosTokenBase
//				err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
//				suite.Require().NoError(err)
//			},
//			true,
//		},
//	}
//	for _, tc := range testCases {
//		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
//			suite.SetupTest() // reset
//
//			tc.malleate()
//
//			pair, err := suite.app.Erc20Keeper.RegisterCoin(suite.ctx, metadata)
//			suite.Commit()
//
//			expPair := &types.TokenPair{
//				Erc20Address:  "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
//				Denom:         "acoin",
//				Enabled:       true,
//				ContractOwner: 1,
//			}
//
//			if tc.expPass {
//				suite.Require().NoError(err, tc.name)
//				suite.Require().Equal(pair, expPair)
//			} else {
//				suite.Require().Error(err, tc.name)
//			}
//		})
//	}
//}
//
//func (suite KeeperTestSuite) TestRegisterERC20() {
//	var (
//		contractAddr common.Address
//		pair         types.TokenPair
//	)
//	testCases := []struct {
//		name     string
//		malleate func()
//		expPass  bool
//	}{
//		{
//			"intrarelaying is disabled globally",
//			func() {
//				params := types.DefaultParams()
//				params.EnableErc20 = false
//				suite.app.Erc20Keeper.SetParams(suite.ctx, params)
//			},
//			false,
//		},
//		{
//			"token ERC20 already registered",
//			func() {
//				suite.app.Erc20Keeper.SetERC20Map(suite.ctx, pair.GetERC20Contract(), pair.GetID())
//			},
//			false,
//		},
//		{
//			"denom already registered",
//			func() {
//				suite.app.Erc20Keeper.SetDenomMap(suite.ctx, pair.Denom, pair.GetID())
//			},
//			false,
//		},
//		{
//			"meta data already stored",
//			func() {
//				suite.app.Erc20Keeper.CreateCoinMetadata(suite.ctx, contractAddr)
//			},
//			false,
//		},
//		{
//			"ok",
//			func() {},
//			true,
//		},
//	}
//	for _, tc := range testCases {
//		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
//			suite.SetupTest() // reset
//
//			var err error
//			contractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, cosmosDecimals)
//			suite.Require().NoError(err)
//			suite.Commit()
//			coinName := types.CreateDenom(contractAddr.String())
//			sender := sdk.AccAddress(suite.address.Bytes())
//
//			pair = types.NewTokenPair(contractAddr, coinName, true, types.OWNER_EXTERNAL, sender.String())
//
//			tc.malleate()
//
//			_, err = suite.app.Erc20Keeper.RegisterERC20(suite.ctx, contractAddr)
//			metadata, found := suite.app.BankKeeper.GetDenomMetaData(suite.ctx, coinName)
//			if tc.expPass {
//				suite.Require().NoError(err, tc.name)
//				// Metadata variables
//				suite.Require().True(found)
//				suite.Require().Equal(coinName, metadata.Base)
//				suite.Require().Equal(coinName, metadata.Name)
//				suite.Require().Equal(types.SanitizeERC20Name(erc20Name), metadata.Display)
//				suite.Require().Equal(erc20Symbol, metadata.Symbol)
//				// Denom units
//				suite.Require().Equal(len(metadata.DenomUnits), 2)
//				suite.Require().Equal(coinName, metadata.DenomUnits[0].Denom)
//				suite.Require().Equal(uint32(zeroExponent), metadata.DenomUnits[0].Exponent)
//				suite.Require().Equal(types.SanitizeERC20Name(erc20Name), metadata.DenomUnits[1].Denom)
//				// Custom exponent at contract creation matches coin with token
//				suite.Require().Equal(metadata.DenomUnits[1].Exponent, uint32(cosmosDecimals))
//			} else {
//				suite.Require().Error(err, tc.name)
//			}
//		})
//	}
//}
