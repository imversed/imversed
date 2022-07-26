package infr_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"

	"github.com/imversed/imversed/app"
	"github.com/imversed/imversed/tests"
	"github.com/imversed/imversed/testutil/network"
	testnet "github.com/imversed/imversed/testutil/network"
	"github.com/imversed/imversed/x/infr/types"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	//currencyCli "github.com/imversed/imversed/x/currency/client/cli"
	currencyTypes "github.com/imversed/imversed/x/currency/types"
	"github.com/tharsis/ethermint/crypto/hd"
	ethermint "github.com/tharsis/ethermint/types"

	//"github.com/spf13/cast"

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
	})

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
	suite.network.Cleanup()
	suite.app = nil
	suite.network = nil
	runtime.GC()
}

func (suite *GenesisTestSuite) TestMinGasPrice() {

	val := suite.network.Validators[0]

	//minGasPriceHelper.Create(baseapp.SetMinGasPrices, fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom))

	_, _, err := val.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	suite.Require().NoError(err)

	suite.callCreateNewMember(val)
	//suite.callChangeMinGasPrice(val)
	//suite.callVote(val)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callCreateNewMember(val *network.Validator) {
	clientCtx := val.ClientCtx
	memberNumber := uuid.New().String()

	info, _, err := clientCtx.Keyring.NewMnemonic(fmt.Sprintf("member%s", memberNumber), keyring.English, ethermint.BIP44HDPath,
		keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	suite.Require().NoError(err)

	pk := info.GetPubKey()

	account := sdk.AccAddress(pk.Address())
	clientAccaunt := account.String()
	fmt.Printf("\n Client acc: %s \n\n", clientAccaunt)

	_, err = banktestutil.MsgSendExec(
		clientCtx,
		val.Address,
		account,
		sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(2000))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(10))).String()),
		//fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
	)
	suite.Require().NoError(err)
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
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(suite.cfg.BondDenom, sdk.NewInt(20))).String()),
	}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, priceCmd, priceCmdArgs)

	fmt.Println("---------")
	fmt.Println(out)
	fmt.Println("---------")

	suite.Require().NoError(err)

	var priceCmdResp currencyTypes.QueryAllCurrencyResponse
	suite.Require().NoError(suite.network.Config.Codec.UnmarshalJSON(out.Bytes(), &priceCmdResp))
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *GenesisTestSuite) callVote(val *network.Validator) {
	voteCmd := gov.NewCmdVote()

	voteCmdArgs := []string{}

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, voteCmd, voteCmdArgs)
	suite.Require().NoError(err)

	time.Sleep(35 * time.Second)

	var voteResp currencyTypes.QueryAllCurrencyResponse
	suite.Require().NoError(suite.network.Config.Codec.UnmarshalJSON(out.Bytes(), &voteResp))
}
