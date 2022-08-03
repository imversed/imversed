package infr_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"

	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/tests"
	"github.com/imversed/imversed/testutil/network"
	testnet "github.com/imversed/imversed/testutil/network"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
	"github.com/imversed/imversed/x/infr/types"

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
	"github.com/google/uuid"
	infrCli "github.com/imversed/imversed/x/infr/client/cli"
)

type GenesisTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.ImversedApp
	genesis types.GenesisState

	cfg     network.Config
	network *network.Network
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) SetupTest() {
	// consensus key
	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())
	fmt.Printf("Address: %s", consAddress)

	minGasPriceHelper.Create(baseapp.SetMinGasPrices, fmt.Sprintf("0.01%s", sdk.DefaultBondDenom))

	suite.app = app.Setup(false, nil)

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
	}) //.WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin(sdk.DefaultBondDenom, sdk.NewInt(100))))

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With()
	suite.ctx.WithLogger(logger)
	suite.genesis = *types.DefaultGenesisState()

	// configurating pools and net

	suite.cfg = testnet.DefaultConfig()
	//suite.cfg.JSONRPCAddress = config.DefaultJSONRPCAddress
	suite.cfg.NumValidators = 1

	suite.cfg.GenesisState = app.ModuleBasics.DefaultGenesis(suite.cfg.Codec)

	var err error
	suite.network, err = network.New(suite.T(), suite.T().TempDir(), suite.cfg)
	suite.Require().NoError(err)
	suite.Require().NotNil(suite.network)

	_, err = suite.network.WaitForHeight(1)
	suite.Require().NoError(err)
}

func (suite *GenesisTestSuite) TearDownSuite() {
	/*suite.network.Cleanup()
	suite.app = nil
	suite.network.Logger = nil
	suite.network = nil
	suite.ctx.WithLogger(nil)
	time.Sleep(2 * time.Second)
	runtime.GC()*/
}

func (suite *GenesisTestSuite) TestMinGasPrice() {

	val := suite.network.Validators[0]
	_, _, err := val.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	suite.Require().NoError(err)

	accaunt := suite.callCreateNewMember(val)
	suite.callChangeMinGasPrice(val)
	suite.callVote(val)
	suite.sendMoney(val, accaunt, 9, true)
	suite.sendMoney(val, accaunt, 11, false)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callCreateNewMember(val *network.Validator) sdk.AccAddress {
	clientCtx := val.ClientCtx
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
		val.Address,
		account,
		sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	suite.Require().NoError(err2)
	var txResp sdk.TxResponse
	suite.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp))

	return account
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callChangeMinGasPrice(val *network.Validator) {
	priceCmd := gov.NewCmdSubmitProposal()
	priceCmd.AddCommand(infrCli.NewChangeMinGasPricesProposalCmd())

	priceCmdArgs := []string{
		fmt.Sprintf("--%s=%s", gov.FlagTitle, "test_change_min_gas_price"),
		fmt.Sprintf("--%s=%s", gov.FlagDescription, "Changing min gas price"),
		fmt.Sprintf("--%s=%s", gov.FlagDeposit, fmt.Sprintf("10%s", sdk.DefaultBondDenom)),

		fmt.Sprintf("--%s=%s", gov.FlagProposalType, "Text"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(11))).String()),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, priceCmd, priceCmdArgs)

	suite.Require().NoError(err)

	var txResp sdk.TxResponse
	suite.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callVote(val *network.Validator) {
	voteCmd := gov.NewCmdVote()

	voteCmdArgs := []string{
		"1",
		"yes",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, voteCmd, voteCmdArgs)
	suite.Require().NoError(err)

	time.Sleep(35 * time.Second)

	var txResp sdk.TxResponse
	suite.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp), out.String())
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) sendMoney(val *network.Validator, tergetAddress sdk.AccAddress, fee int64, expectError bool) {
	out, err := banktestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		tergetAddress,
		sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(fee))).String()),
	)
	fmt.Println("\n-------")
	fmt.Println(out)
	fmt.Println("\n-------")

	suite.Require().NoError(err)

	var txResp sdk.TxResponse

	if expectError {
		suite.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp))
	} else {
		suite.Require().Error(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txResp))
		suite.Require().Equal(uint32(13), txResp.Code, txResp)
	}

}
