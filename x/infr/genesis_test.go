package infr_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
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
	testnet "github.com/imversed/imversed/testutil/network"
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
	// consensus key

	suite.baseMinGasPrice = "0.000005"
	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())
	fmt.Printf("Address: %s", consAddress)
	baseapp.SetMinGasPrices(suite.money(suite.baseMinGasPrice))
	minGasPriceHelper.Create(baseapp.SetMinGasPrices, suite.money(suite.baseMinGasPrice))

	suite.cfg = testnet.DefaultConfig()
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
	accaunt := suite.callCreateNewMember()

	suite.sendMoney(accaunt, suite.money("0.9"), true)
	suite.sendMoney(accaunt, suite.money("1.1"), false)
	suite.callChangeMinGasPrice()
	time.Sleep(10 * time.Second)
	suite.callVote()
	time.Sleep(65 * time.Second)
	suite.showCurrentProposals()
	suite.sendMoney(accaunt, suite.money("19"), true)
	suite.sendMoney(accaunt, suite.money("21"), false)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callCreateNewMember() sdk.AccAddress {
	clientCtx := suite.validator.ClientCtx
	memberNumber := uuid.New().String()

	info, _, err := clientCtx.Keyring.NewMnemonic(fmt.Sprintf("member%s", memberNumber), keyring.English, ethermint.BIP44HDPath,
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
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300))),
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
	propCmd := gov.NewCmdSubmitProposal()
	priceCmd := infrCli.NewChangeMinGasPricesProposalCmd()

	priceCmd.SetArgs([]string{suite.money("0.0001")})

	propCmd.AddCommand(priceCmd)

	propCmdArgs := []string{
		fmt.Sprintf("--%s=%s", gov.FlagTitle, "test_change_min_gas_price"),
		fmt.Sprintf("--%s=%s", gov.FlagDescription, "Changing min gas price"),
		fmt.Sprintf("--%s=%s", gov.FlagDeposit, suite.money("10000000")),
		fmt.Sprintf("--%s=%s", gov.FlagProposalType, "Text"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, suite.validator.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, suite.money("1")),
	}

	out, err := clitestutil.ExecTestCLICmd(suite.validator.ClientCtx, propCmd, propCmdArgs)

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
	return fmt.Sprintf("%s%s", coins, sdk.DefaultBondDenom)
}
