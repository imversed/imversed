package infr_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/imversed/imversed/x/infr/baseAppHelper"
	"runtime"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

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

	"github.com/evmos/ethermint/crypto/hd"
	ethermint "github.com/evmos/ethermint/types"

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

	cfg       network.Config
	network   *network.Network
	validator *network.Validator
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
	suite.cfg = network.DefaultConfig().WithDenom(network.DefaultBondDenom, "0.000005")

	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())

	baseapp.SetMinGasPrices(suite.cfg.MinGasPrices)
	baseAppHelper.Create(baseapp.SetMinGasPrices, suite.cfg.MinGasPrices)

	suite.cfg.NumValidators = 1
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
	})

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
}

func (suite *GenesisTestSuite) TearDownSuite() {
	suite.network.Cleanup()
	suite.app = nil
	suite.network.Logger = nil
	suite.network = nil
	suite.ctx.WithLogger(nil)
	time.Sleep(1 * time.Second)
	runtime.GC()
	time.Sleep(1 * time.Second)
}

func (suite *GenesisTestSuite) TestMinGasPrice() {
	account := suite.callCreateNewMember()
	suite.sendMoney(account, suite.money("0.9"), true)
	suite.sendMoney(account, suite.money("1.1"), false)
	suite.callChangeMinGasPrice()
	time.Sleep(10 * time.Second)
	suite.callVote()
	time.Sleep(45 * time.Second)
	//suite.showCurrentProposals()
	suite.sendMoney(account, suite.money("19"), true)
	suite.sendMoney(account, suite.money("21"), false)
	//fmt.Println("Done!")
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
func (suite *GenesisTestSuite) callChangeMinGasPrice() {
	priceCmd := infrCli.NewChangeMinGasPricesProposalCmd()
	flags.AddTxFlagsToCmd(priceCmd)
	propCmdArgs := []string{
		suite.money("0.0001"),
		fmt.Sprintf("--%s=%s", gov.FlagTitle, "test_change_min_gas_price"),
		fmt.Sprintf("--%s=%s", gov.FlagDescription, "Changing min gas price"),
		fmt.Sprintf("--%s=%s", gov.FlagDeposit, suite.money("10000000")),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, suite.validator.Address.String()),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	}

	out, err := clitestutil.ExecTestCLICmd(suite.validator.ClientCtx, priceCmd, propCmdArgs)

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
		fmt.Println("Unexpected error in tx response:")
		fmt.Println(txResp)
	}
}

func (suite *GenesisTestSuite) money(coins string) string {
	return fmt.Sprintf("%s%s", coins, network.DefaultBondDenom)
}
