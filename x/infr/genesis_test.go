package infr_test

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/tests"
	"github.com/imversed/imversed/testutil/network"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	//currencyCli "github.com/imversed/imversed/x/currency/client/cli"

	"github.com/tharsis/ethermint/crypto/hd"
	ethermint "github.com/tharsis/ethermint/types"

	//"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	gov "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/google/uuid"
	infrCli "github.com/imversed/imversed/x/infr/client/cli"
)

type GenesisTestSuite struct {
	suite.Suite
	ctx sdk.Context
	app *app.ImversedApp

	cfg             network.Config
	network         *network.Network
	validator       *network.Validator
	baseMinGasPrice string
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) patchGenesis(app *app.ImversedApp, sapp simapp.GenesisState) simapp.GenesisState {
	govGen := govTypes.DefaultGenesisState()
	govGen.VotingParams.VotingPeriod = 30 * time.Second

	poolsGenJson := suite.cfg.Codec.MustMarshalJSON(govGen)
	sapp[govTypes.ModuleName] = poolsGenJson
	suite.cfg.GenesisState = sapp
	return sapp
}

func (suite *GenesisTestSuite) SetupTest() {
	network.InitTestConfig()

	suite.baseMinGasPrice = "0.000005"
	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())
	fmt.Printf("\nAddress: %s\n", consAddress)
	baseapp.SetMinGasPrices(suite.money(suite.baseMinGasPrice))
	minGasPriceHelper.Create(baseapp.SetMinGasPrices, suite.money(suite.baseMinGasPrice))

	suite.cfg = network.DefaultConfig()
	suite.cfg.NumValidators = 2

	suite.app = app.Setup(false, suite.patchGenesis)

	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{
		Height:          1,
		ChainID:         "imversed_1234-1",
		Time:            time.Now().UTC(),
		ProposerAddress: consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	}) //.WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin(sdk.DefaultBondDenom, sdk.NewInt(1))))

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	suite.ctx.WithLogger(logger)

	var err error
	suite.network, err = network.New(suite.T(), suite.T().TempDir(), suite.cfg)

	suite.Require().NoError(err)
	suite.Require().NotNil(suite.network)

	_, err = suite.network.WaitForHeight(1)
	suite.Require().NoError(err)

	suite.validator = suite.network.Validators[0]
	_, _, err = suite.validator.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)

	suite.Require().NoError(err)
	fmt.Printf("\nValidator: %s\n", suite.validator.Address)
}

func (suite *GenesisTestSuite) TearDownSuite() {
	suite.network.Cleanup()
	suite.app = nil
	suite.network.Logger = nil
	suite.network = nil
	suite.ctx.WithLogger(nil)
	runtime.GC()
	time.Sleep(2 * time.Second)
}

func (suite *GenesisTestSuite) TestMinGasPrice() {
	suite.createAccountForValidator()
	account := suite.callCreateNewMember()

	suite.sendMoney(account, suite.money("0.9"), true)
	suite.sendMoney(account, suite.money("1.1"), false)
	suite.callChangeMinGasPrice()
	time.Sleep(10 * time.Second)
	suite.callVote()
	time.Sleep(65 * time.Second)
	suite.showCurrentProposals()
	suite.sendMoney(account, suite.money("19"), true)
	suite.sendMoney(account, suite.money("21"), false)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callCreateNewMember() sdk.AccAddress {
	clientCtx := suite.validator.ClientCtx
	memberNumber := uuid.New().String()

	info, _, err := clientCtx.Keyring.NewMnemonic(fmt.Sprintf("member%s", memberNumber),
		keyring.English, ethermint.BIP44HDPath,
		keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	suite.Require().NoError(err)

	pk := info.GetPubKey()

	account := sdk.AccAddress(pk.Address())
	clientAccaunt := account.String()
	fmt.Printf("\n Client acc: %s \n\n", clientAccaunt)

	out, err2 := banktestutil.MsgSendExec(
		clientCtx,
		suite.validator.Address,
		account,
		sdk.NewCoins(sdk.NewCoin(network.DefaultBondDenom, sdk.NewInt(300))),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	)
	suite.Require().NoError(err2)
	var txResp sdk.TxResponse
	suite.Require().NoError(suite.validator.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp))
	suite.printIfError(txResp)
	suite.Require().Equal(uint32(0), txResp.Code)

	return account
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) createAccountForValidator() {
	val1 := suite.validator
	val2 := suite.network.Validators[1]

	k, err := val1.ClientCtx.Keyring.KeyByAddress(val1.Address)
	suite.Require().NoError(err)
	keyName := k.GetName()
	suite.Require().NotNil(keyName)
	addr := k.GetAddress()
	suite.Require().NotNil(addr)

	account, err := val1.ClientCtx.AccountRetriever.GetAccount(val1.ClientCtx, addr)
	suite.Require().NoError(err)
	suite.Require().NotNil(account)

	sendTokens := sdk.NewCoins(sdk.NewCoin(network.DefaultBondDenom, sdk.NewInt(1000)))
	args := []string{
		keyName,
		val2.Address.String(),
		sendTokens.String(),
		//fmt.Sprintf("--%s=true", flags.FlagGenerateOnly),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}
	generatedStd, err := clitestutil.ExecTestCLICmd(val1.ClientCtx, bankCli.NewSendTxCmd(), args)
	suite.Require().NoError(err)
	suite.Require().NotNil(generatedStd)

}

func ExecTestCLICmd(clientCtx client.Context, cmd *cobra.Command) (testutil.BufferWriter, error) {
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	if err := cmd.ExecuteContext(ctx); err != nil {
		return out, err
	}

	return out, nil
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callChangeMinGasPrice() {
	key, err := suite.validator.ClientCtx.Keyring.KeyByAddress(suite.validator.Address)
	suite.Require().NoError(err)

	propCmd := gov.NewCmdSubmitProposal()

	priceCmd := infrCli.NewChangeMinGasPricesProposalCmd()
	priceCmd.SetArgs([]string{
		suite.money("0.0001"),
		fmt.Sprintf("--%s=%s", gov.FlagTitle, "test_change_min_gas_price"),
		fmt.Sprintf("--%s=%s", gov.FlagDescription, "Changing min gas price"),
		fmt.Sprintf("--%s=%s", gov.FlagDeposit, suite.money("10000000")),
	})

	propCmd.AddCommand(priceCmd)
	propCmd.SetArgs([]string{
		//fmt.Sprintf("--%s=%s", gov.FlagDeposit, suite.money("10000000")),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, key.GetName()),
		//fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	})

	/*cmdArgs := []string{
		suite.money("0.0001"),
		fmt.Sprintf("--%s=%s", gov.FlagTitle, "test_change_min_gas_price"),
		fmt.Sprintf("--%s=%s", gov.FlagDescription, "Changing min gas price"),
		fmt.Sprintf("--%s=%s", gov.FlagDeposit, suite.money("10000000")),

		fmt.Sprintf("--%s=%s", flags.FlagFrom, suite.validator.Address.String()),
		fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	}

	priceCmd.SetArgs(cmdArgs)*/
	var out testutil.BufferWriter
	out, err = ExecTestCLICmd(suite.validator.ClientCtx, priceCmd)

	suite.Require().NoError(err)

	var txResp sdk.TxResponse
	suite.Require().NoError(suite.validator.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	suite.printIfError(txResp)
	suite.Require().Equal(uint32(0), txResp.Code)

}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) showCurrentProposals() {
	queryCmd := gov.GetCmdQueryProposals()
	queryOut, queryErr := clitestutil.ExecTestCLICmd(suite.validator.ClientCtx, queryCmd, []string{})
	suite.Require().NoError(queryErr)
	fmt.Println(queryOut)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callVote() {
	voteCmd := gov.NewCmdVote()

	voteCmdArgs := []string{
		"1",
		"yes",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, suite.validator.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	}

	out, err := clitestutil.ExecTestCLICmd(suite.validator.ClientCtx, voteCmd, voteCmdArgs)
	suite.Require().NoError(err)

	var txResp sdk.TxResponse
	suite.Require().NoError(suite.validator.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
	suite.printIfError(txResp)
	suite.Require().Equal(uint32(0), txResp.Code)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) sendMoney(tergetAddress sdk.AccAddress, fee string, expectError bool) {
	out, err := banktestutil.MsgSendExec(
		suite.validator.ClientCtx,
		suite.validator.Address,
		tergetAddress,
		sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, fee),
	)
	suite.Require().NoError(err)

	var txResp sdk.TxResponse
	suite.Require().NoError(suite.validator.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp))

	//fmt.Println(txResp)
	if expectError {
		suite.Require().Equal(uint32(13), txResp.Code)
	} else {
		suite.printIfError(txResp)
		suite.Require().Equal(uint32(0), txResp.Code)
	}

}

func (suite *GenesisTestSuite) printIfError(txResp sdk.TxResponse) {
	if txResp.Code != 0 {
		fmt.Println(txResp)
	}
}

func (suite *GenesisTestSuite) money(coins string) string {
	return fmt.Sprintf("%s%s", coins, network.DefaultBondDenom)
}
