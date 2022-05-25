package cli_test

import (
	"fmt"
	ethermint "github.com/tharsis/ethermint/types"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/imversed/imversed/app"
	//"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/imversed/imversed/testutil/network"
	"github.com/imversed/imversed/x/pools/client/cli"
	poolstestutil "github.com/imversed/imversed/x/pools/client/testutil"
	"github.com/imversed/imversed/x/pools/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tharsis/ethermint/crypto/hd"

	testnet "github.com/imversed/imversed/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.cfg = testnet.DefaultConfig()
	s.cfg.NumValidators = 1

	// modification to pay fee with test bond denom "stake"
	genesisState := app.ModuleBasics.DefaultGenesis(s.cfg.Codec)
	poolsGen := types.DefaultGenesis()
	poolsGen.Params.PoolCreationFee = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 1000000)}
	poolsGenJson := s.cfg.Codec.MustMarshalJSON(poolsGen)
	genesisState[types.ModuleName] = poolsGenJson

	s.cfg.GenesisState = genesisState

	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]

	// create a new pool
	_, err = poolstestutil.MsgCreatePool(s.T(), val.ClientCtx, val.Address, "5stake,5node0token", "100stake,100node0token", "0.01", "0.01")
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestNewCreatePoolCmd() {
	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	var buf testutil.BufferWriter
	buf, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewInt64Coin(s.cfg.BondDenom, 200000000), sdk.NewInt64Coin("node0token", 20000)), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	_ = buf
	s.Require().NoError(err)

	testCases := []struct {
		name         string
		json         string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"one token pair pool",
			fmt.Sprintf(`
			{
			  "%s": "1node0token",
			  "%s": "100node0token",
			  "%s": "0.001",
			  "%s": "0.001"
			}
			`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee),
			true, &sdk.TxResponse{}, 4,
		},
		{
			"two tokens pair pool",
			fmt.Sprintf(`
				{
				  "%s": "1node0token,3stake",
				  "%s": "100node0token,100stake",
				  "%s": "0.001",
				  "%s": "0.001"
				}
				`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee),
			false, &sdk.TxResponse{}, 0,
		},
		{
			"change order of json fields",
			fmt.Sprintf(`
				{
				  "%s": "100node0token,100stake",
				  "%s": "0.001",
				  "%s": "1node0token,3stake",
				  "%s": "0.001"
				}
				`, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileWeights, cli.PoolFileExitFee),
			false, &sdk.TxResponse{}, 0,
		},
		{ // --record-tokens=100.0stake2 --record-tokens=100.0stake --record-tokens-weight=5 --record-tokens-weight=5 --swap-fee=0.01 --exit-fee=0.01 --from=validator --keyring-backend=test --chain-id=testing --yes
			"three tokens pair pool - insufficient balance check",
			fmt.Sprintf(`
				{
				  "%s": "1node0token,1stake,2btc",
				  "%s": "100node0token,100stake,100btc",
				  "%s": "0.001",
				  "%s": "0.001"
				}
				`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee),
			false, &sdk.TxResponse{}, 5,
		},
		{
			"not valid json",
			"bad json",
			true, &sdk.TxResponse{}, 0,
		},
		{
			"bad pool json - missing quotes around exit fee",
			fmt.Sprintf(`
			{
			  "%s": "1node0token,3stake",
			  "%s": "100node0token,100stake",
			  "%s": "0.001",
			  "%s": 0.001
			}
	`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee),
			true, &sdk.TxResponse{}, 0,
		},
		{
			"empty pool json",
			"", true, &sdk.TxResponse{}, 0,
		},
		{
			"smooth change params",
			fmt.Sprintf(`
					{
						"%s": "1node0token,3stake",
						"%s": "100node0token,100stake",
						"%s": "0.001",
						"%s": "0.001",
						"%s": {
							"%s": "864h",
							"%s": "2node0token,1stake",
							"%s": "2006-01-02T15:04:05Z"
						}
					}
					`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee,
				cli.PoolFileSmoothWeightChangeParams, cli.PoolFileDuration, cli.PoolFileTargetPoolWeights, cli.PoolFileStartTime,
			),
			false, &sdk.TxResponse{}, 0,
		},
		{
			"smooth change params - no start time",
			fmt.Sprintf(`
					{
						"%s": "1node0token,3stake",
						"%s": "100node0token,100stake",
						"%s": "0.001",
						"%s": "0.001",
						"%s": {
							"%s": "864h",
							"%s": "2node0token,1stake"
						}
					}
					`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee,
				cli.PoolFileSmoothWeightChangeParams, cli.PoolFileDuration, cli.PoolFileTargetPoolWeights,
			),
			false, &sdk.TxResponse{}, 0,
		},
		{
			"empty smooth change params",
			fmt.Sprintf(`
					{
						"%s": "1node0token,3stake",
						"%s": "100node0token,100stake",
						"%s": "0.001",
						"%s": "0.001",
						"%s": {}
					}
					`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee,
				cli.PoolFileSmoothWeightChangeParams,
			),
			false, &sdk.TxResponse{}, 0,
		},
		{
			"smooth change params wrong type",
			fmt.Sprintf(`
				{
					"%s": "1node0token,3stake",
					"%s": "100node0token,100stake",
					"%s": "0.001",
					"%s": "0.001",
					"%s": "invalid string"
				}
				`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee,
				cli.PoolFileSmoothWeightChangeParams,
			),
			true, &sdk.TxResponse{}, 0,
		},
		{
			"smooth change params missing duration",
			fmt.Sprintf(`
				{
					"%s": "1node0token,3stake",
					"%s": "100node0token,100stake",
					"%s": "0.001",
					"%s": "0.001",
					"%s": {
						"%s": "2node0token,1stake"
					}
				}
				`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee,
				cli.PoolFileSmoothWeightChangeParams, cli.PoolFileTargetPoolWeights,
			),
			true, &sdk.TxResponse{}, 0,
		},
		{
			"unknown fields in json",
			fmt.Sprintf(`
			{
			  "%s": "1node0token",
			  "%s": "100node0token",
			  "%s": "0.001",
			  "%s": "0.001"
			  "unknown": true,
			}
			`, cli.PoolFileWeights, cli.PoolFileInitialDeposit, cli.PoolFileSwapFee, cli.PoolFileExitFee),
			true, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewCreatePoolCmd()
			clientCtx := val.ClientCtx

			jsonFile := testutil.WriteToNewTempFile(s.T(), tc.json)

			args := []string{
				fmt.Sprintf("--%s=%s", cli.FlagPoolFile, jsonFile.Name()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
				fmt.Sprintf("--%s=%s", flags.FlagGas, fmt.Sprint(300000)),
			}

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				err = clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType)
				s.Require().NoError(err, out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewJoinPoolCmd() {

	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewJoinPoolAddr", keyring.English, ethermint.BIP44HDPath, "", hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewInt64Coin(s.cfg.BondDenom, 20000), sdk.NewInt64Coin("node0token", 20000)), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"join pool with insufficient balance",
			[]string{
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", cli.FlagMaxAmountsIn, "100stake"),
				fmt.Sprintf("--%s=%s", cli.FlagShareAmountOut, "1000000000000000000000"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 6,
		},
		{
			"join pool with sufficient balance",
			[]string{ // join-pool --pool-id=1 --max-amounts-in=100stake --share-amount-out=100 --from=validator --keyring-backend=test --chain-id=testing --yes
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", cli.FlagMaxAmountsIn, "100stake"),
				fmt.Sprintf("--%s=%s", cli.FlagShareAmountOut, "10000000000000000000"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewJoinPoolCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewExitPoolCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"ask too much when exit",
			[]string{ // --min-amounts-out=100stake --pool-id=1 --share-amount-in=10 --from=validator --keyring-backend=test --chain-id=testing --yes
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", cli.FlagShareAmountIn, "20000000000000000000"),
				fmt.Sprintf("--%s=%s", cli.FlagMinAmountsOut, "20stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 7,
		},
		{
			"ask enough when exit",
			[]string{ // --min-amounts-out=100stake --pool-id=1 --share-amount-in=10 --from=validator --keyring-backend=test --chain-id=testing --yes
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", cli.FlagShareAmountIn, "20000000000000000000"),
				fmt.Sprintf("--%s=%s", cli.FlagMinAmountsOut, "10stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewExitPoolCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewSwapExactAmountOutCmd() {

	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewSwapExactAmountOut",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewInt64Coin(s.cfg.BondDenom, 20000), sdk.NewInt64Coin("node0token", 20000)), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	testCases := []struct {
		name string
		args []string

		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"swap exact amount out", // imversed tx pools swap-exact-amount-out 10stake 20 --swap-route-pool-ids=1 --swap-route-denoms=node0token --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"10stake", "20",
				fmt.Sprintf("--%s=%d", cli.FlagSwapRoutePoolIds, 1),
				fmt.Sprintf("--%s=%s", cli.FlagSwapRouteDenoms, "node0token"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewSwapExactAmountOutCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewJoinSwapExternAmountInCmd() {

	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewJoinSwapExternAmountIn",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(20000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"join swap extern amount in", // imversed tx pools join-swap-extern-amount-in --pool-id=1 10stake 1 --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"10stake", "1",
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewJoinSwapExternAmountIn()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewExitSwapExternAmountOutCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"exit swap extern amount out", // imversed tx pools exit-swap-extern-amount-out --pool-id=1 10stake 1 --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"10stake", "10000000000000000000",
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewExitSwapExternAmountOut()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewJoinSwapShareAmountOutCmd() {

	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewJoinSwapShareAmountOutAddr", keyring.English,
		ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(20000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"join swap share amount out", // imversed tx pools join-swap-share-amount-out --pool-id=1 stake 10 1 --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"stake", "50", "5000000000000000000",
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewJoinSwapShareAmountOut()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewExitSwapShareAmountInCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"exit swap share amount in", // imversed tx pools exit-swap-share-amount-in --pool-id=1 stake 10 1 --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"stake", "10000000000000000000", "1",
				fmt.Sprintf("--%s=%d", cli.FlagPoolId, 1),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewExitSwapShareAmountIn()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdPools() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pools",
			[]string{
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdPools() // imversed query pools pools
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryPoolsResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())

				s.Require().Greater(len(resp.Pools), 0, out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdNumPools() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query num-pools",
			[]string{
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdNumPools() // imversed query pools num-pools
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryNumPoolsResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())

				s.Require().Greater(resp.NumPools, uint64(0), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdPool() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool by id", // imversed query pools pool 1
			[]string{
				"1",
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdPool()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())

				resp := types.QueryPoolResponse{}
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdPoolParams() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool params by id", // imversed query pools pool-params 1
			[]string{
				"1",
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdPoolParams()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryPoolParamsResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdPoolAssets() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool assets by pool id", // imversed query pools pool-assets 1
			[]string{
				"1",
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdPoolAssets()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryPoolAssetsResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdTotalShares() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool total share by id", // imversed query pools total-share 1
			[]string{
				"1",
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdTotalShares()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryTotalSharesResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdTotalLiquidity() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query total liquidity", // imversed query pools total-liquidity
			[]string{
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryTotalLiquidity()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QueryTotalLiquidityResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestGetCmdSpotPrice() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool spot price", // imversed query pools spot-price 1 stake node0token
			[]string{
				"1", "stake", "node0token",
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdSpotPrice()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QuerySpotPriceResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

//func (s *IntegrationTestSuite) TestGetCmdEstimateSwapExactAmountIn() {
//	val := s.network.Validators[0]
//
//	testCases := []struct {
//		name      string
//		args      []string
//		expectErr bool
//	}{
//		{
//			"query pool estimate swap exact amount in", // imversed query pools estimate-swap-exact-amount-in 1 cosmos1n8skk06h3kyh550ad9qketlfhc2l5dsdevd3hq 10.0stake --swap-route-pool-ids=1 --swap-route-denoms=node0token
//			[]string{
//				"1",
//				"cosmos1n8skk06h3kyh550ad9qketlfhc2l5dsdevd3hq",
//				"10.0stake",
//				fmt.Sprintf("--%s=%d", cli.FlagSwapRoutePoolIds, 1),
//				fmt.Sprintf("--%s=%s", cli.FlagSwapRouteDenoms, "node0token"),
//				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
//			},
//			false,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//
//		s.Run(tc.name, func() {
//			cmd := cli.GetCmdEstimateSwapExactAmountIn()
//			clientCtx := val.ClientCtx
//
//			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
//			if tc.expectErr {
//				s.Require().Error(err)
//			} else {
//				resp := types.QuerySwapExactAmountInResponse{}
//				s.Require().NoError(err, out.String())
//				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
//			}
//		})
//	}
//}

func (s *IntegrationTestSuite) TestGetCmdEstimateSwapExactAmountOut() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query pool estimate swap exact amount in", // imversed query pools estimate-swap-exact-amount-in 1 cosmos1n8skk06h3kyh550ad9qketlfhc2l5dsdevd3hq 10.0stake --swap-route-pool-ids=1 --swap-route-denoms=node0token
			[]string{
				"1",
				val.Address.String(),
				"10.0stake",
				fmt.Sprintf("--%s=%d", cli.FlagSwapRoutePoolIds, 1),
				fmt.Sprintf("--%s=%s", cli.FlagSwapRouteDenoms, "node0token"),
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdEstimateSwapExactAmountOut()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				resp := types.QuerySwapExactAmountOutResponse{}
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &resp), out.String())
			}
		})
	}
}

func (s IntegrationTestSuite) TestNewSwapExactAmountInCmd() {

	val := s.network.Validators[0]

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("NewSwapExactAmountIn",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())

	_, err = banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewInt64Coin(s.cfg.BondDenom, 20000), sdk.NewInt64Coin("node0token", 20000)), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	s.Require().NoError(err)

	testCases := []struct {
		name string
		args []string

		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"swap exact amount in", // imversed tx pools swap-exact-amount-in 10stake 3 --swap-route-pool-ids=1 --swap-route-denoms=node0token --from=validator --keyring-backend=test --chain-id=testing --yes
			[]string{
				"10stake", "3",
				fmt.Sprintf("--%s=%d", cli.FlagSwapRoutePoolIds, 1),
				fmt.Sprintf("--%s=%s", cli.FlagSwapRouteDenoms, "node0token"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewSwapExactAmountInCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
