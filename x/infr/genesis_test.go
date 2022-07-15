package infr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"

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
	currencyCli "github.com/imversed/imversed/x/currency/client/cli"
	currencyTypes "github.com/imversed/imversed/x/currency/types"
	"github.com/tharsis/ethermint/crypto/hd"
	ethermint "github.com/tharsis/ethermint/types"
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
	suite.cfg.NumValidators = 1

	// modification to pay fee with test bond denom "stake"
	genesisState := app.ModuleBasics.DefaultGenesis(suite.cfg.Codec)
	suite.cfg.GenesisState = genesisState

	var err error
	suite.network, err = network.New(suite.T(), suite.T().TempDir(), suite.cfg)
	suite.Require().NoError(err)

	_, err = suite.network.WaitForHeight(1)
	suite.Require().NoError(err)

	//val := suite.network.Validators[0]

	// create a new pool
	//_, err = poolstestutil.MsgCreatePool(suite.T(), val.ClientCtx, val.Address, "5stake,5node0token", "100stake,100node0token", "0.01", "0.01")
	//suite.Require().NoError(err)

	_, err = suite.network.WaitForHeight(1)
	suite.Require().NoError(err)
}

func (suite *GenesisTestSuite) TearDownSuite() {
	suite.network.Cleanup()
	suite.app = nil
	suite.network = nil

}

func (suite *GenesisTestSuite) TestMinGasPrice() {

	val := suite.network.Validators[0]

	_, _, err := val.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
	suite.Require().NoError(err)

	//newAddr := sdk.AccAddress(info.GetPubKey().Address())

	ctx := suite.network.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}

	args := request(nil, uint64(0), uint64(2), false)
	out, err := clitestutil.ExecTestCLICmd(ctx, currencyCli.CmdListCurrency(), args)
	suite.Require().NoError(err)
	var resp currencyTypes.QueryAllCurrencyResponse
	suite.Require().NoError(suite.network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
}
