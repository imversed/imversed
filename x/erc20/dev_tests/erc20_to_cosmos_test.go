package dev_tests

import (
	"context"
	"fmt"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
	"github.com/tharsis/ethermint/server/config"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	"math/big"
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
	//currencyCli "github.com/imversed/imversed/x/currency/client/cli"

	"github.com/tharsis/ethermint/crypto/hd"
	ethermint "github.com/tharsis/ethermint/types"

	//"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/baseapp"
	ercCli "github.com/imversed/imversed/x/erc20/client/cli"
)

type IntegrationTestSuite struct {
	suite.Suite
	app *app.ImversedApp

	cfg        network.Config
	network    *network.Network
	validator  *network.Validator
	ctx        context.Context
	gethClient *gethclient.Client
	ethSigner  ethtypes.Signer
	rpcClient  *rpc.Client
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupTest() {

	network.InitTestConfig()
	s.cfg = network.DefaultConfig().WithDenom(network.DefaultBondDenom, "0.000005")

	consAddress := sdk.ConsAddress(tests.GenerateAddress().Bytes())

	baseapp.SetMinGasPrices(s.cfg.MinGasPrices)
	minGasPriceHelper.Create(baseapp.SetMinGasPrices, s.cfg.MinGasPrices)

	s.cfg.JSONRPCAddress = config.DefaultJSONRPCAddress
	s.cfg.NumValidators = 1
	s.ctx = context.Background()

	s.app = app.Setup(false, nil)

	_ = s.app.BaseApp.NewContext(false, tmproto.Header{
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
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)

	s.Require().NoError(err)
	s.Require().NotNil(s.network)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.validator = s.network.Validators[0]
	_, _, err = s.validator.ClientCtx.Keyring.NewMnemonic("NewCreatePoolAddr",
		keyring.English, ethermint.BIP44HDPath, keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)

	s.Require().NoError(err)

	address := fmt.Sprintf("http://%s", s.network.Validators[0].AppConfig.JSONRPC.Address)

	if s.network.Validators[0].JSONRPCClient == nil {
		s.network.Validators[0].JSONRPCClient, err = ethclient.Dial(address)
		s.Require().NoError(err)
	}

	rpcClient, err := rpc.DialContext(s.ctx, address)
	s.Require().NoError(err)
	s.rpcClient = rpcClient
	s.gethClient = gethclient.New(rpcClient)
	s.Require().NotNil(s.gethClient)
	chainId, err := ethermint.ParseChainID(s.cfg.ChainID)
	s.Require().NoError(err)
	s.ethSigner = ethtypes.LatestSignerForChainID(chainId)

}

func (suite *IntegrationTestSuite) TearDownSuite() {
	suite.network.Cleanup()
	runtime.GC()
}

func (s *IntegrationTestSuite) TestEvmHook() {
	contractHash, contractAddress := s.deployERC20Contract()
	fmt.Println("ContractHash: " + contractHash.String())
	fmt.Println("ContractAddress: " + contractAddress.String())

	//s.sendTransaction(contractAddress)

	s.showCurrentContracts()
	fmt.Println("Done!")
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (s *IntegrationTestSuite) sendTransaction(contractAddress common.Address) {
	blockNum, err := s.network.Validators[0].JSONRPCClient.BlockNumber(s.ctx)
	s.Require().NoError(err)

	s.transferERC20Transaction(contractAddress, common.HexToAddress("0x378c50D9264C63F3F92B806d4ee56E9D86FfB3Ec"), big.NewInt(10))
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNum)),
	}

	logs, err := s.network.Validators[0].JSONRPCClient.FilterLogs(s.ctx, filterQuery)
	s.Require().NoError(err)
	s.Require().NotNil(logs)
	s.Require().Equal(1, len(logs))

	expectedTopics := []common.Hash{
		common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		common.HexToHash("0x000000000000000000000000" + fmt.Sprintf("%x", common.BytesToAddress(s.network.Validators[0].Address))),
		common.HexToHash("0x000000000000000000000000378c50d9264c63f3f92b806d4ee56e9d86ffb3ec"),
	}

	s.Require().Equal(expectedTopics, logs[0].Topics)
}

func (s *IntegrationTestSuite) transferERC20Transaction(contractAddr, to common.Address, amount *big.Int) common.Hash {
	chainID, err := s.network.Validators[0].JSONRPCClient.ChainID(s.ctx)
	s.Require().NoError(err)

	transferData, err := evmtypes.ERC20Contract.ABI.Pack("transfer", to, amount)
	s.Require().NoError(err)
	owner := common.BytesToAddress(s.network.Validators[0].Address)
	nonce := s.getAccountNonce(owner)

	gas, err := s.network.Validators[0].JSONRPCClient.EstimateGas(s.ctx, ethereum.CallMsg{
		To:   &contractAddr,
		From: owner,
		Data: transferData,
	})
	s.Require().NoError(err)

	gasPrice := s.getGasPrice()
	ercTransferTx := evmtypes.NewTx(
		chainID,
		nonce,
		&contractAddr,
		nil,
		gas,
		gasPrice,
		nil, nil,
		transferData,
		nil,
	)

	ercTransferTx.From = owner.Hex()
	err = ercTransferTx.Sign(s.ethSigner, s.network.Validators[0].ClientCtx.Keyring)
	s.Require().NoError(err)
	err = s.network.Validators[0].JSONRPCClient.SendTransaction(s.ctx, ercTransferTx.AsTransaction())
	s.Require().NoError(err)

	s.waitForTransaction()

	receipt := s.expectSuccessReceipt(ercTransferTx.AsTransaction().Hash())
	s.Require().NotEmpty(receipt.Logs)
	return ercTransferTx.AsTransaction().Hash()

}

func (suite *IntegrationTestSuite) showName() {

}

func (suite *IntegrationTestSuite) showCurrentContracts() {
	queryCmd := ercCli.GetTokenPairsCmd()
	queryOut, queryErr := clitestutil.ExecTestCLICmd(suite.validator.ClientCtx, queryCmd, []string{})
	suite.Require().NoError(queryErr)
	fmt.Println(queryOut)
}

func (s *IntegrationTestSuite) deployERC20Contract() (transaction common.Hash, contractAddr common.Address) {
	owner := common.BytesToAddress(s.network.Validators[0].Address)
	supply := sdk.NewIntWithDecimal(1000, 18).BigInt()

	ctorArgs, err := evmtypes.ERC20Contract.ABI.Pack("", owner, supply)
	s.Require().NoError(err)
	//fmt.Println(hex.EncodeToString(evmtypes.ERC20Contract.Bin))
	//fmt.Println(ctorArgs)
	data := append(evmtypes.ERC20Contract.Bin, ctorArgs...)
	//fmt.Println(hex.EncodeToString(data))
	return s.deployContract(data)
}

func (s *IntegrationTestSuite) getAccountNonce(addr common.Address) uint64 {
	nonce, err := s.network.Validators[0].JSONRPCClient.NonceAt(s.ctx, addr, nil)
	s.Require().NoError(err)
	return nonce
}

func (s *IntegrationTestSuite) getGasPrice() *big.Int {
	gasPrice, err := s.network.Validators[0].JSONRPCClient.SuggestGasPrice(s.ctx)
	s.Require().NoError(err)
	return gasPrice
}

// waits 2 blocks time to keep tests stable
func (s *IntegrationTestSuite) waitForTransaction() {
	err := s.network.WaitForNextBlock()
	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) expectSuccessReceipt(hash common.Hash) *ethtypes.Receipt {
	receipt, err := s.network.Validators[0].JSONRPCClient.TransactionReceipt(s.ctx, hash)
	s.Require().NoError(err)
	s.Require().NotNil(receipt)
	s.Require().Equal(uint64(0x1), receipt.Status)
	return receipt
}

func (s *IntegrationTestSuite) deployContract(data []byte) (transaction common.Hash, contractAddr common.Address) {
	chainID, err := s.network.Validators[0].JSONRPCClient.ChainID(s.ctx)
	s.Require().NoError(err)

	owner := common.BytesToAddress(s.network.Validators[0].Address)
	nonce := s.getAccountNonce(owner)

	gas, err := s.network.Validators[0].JSONRPCClient.EstimateGas(s.ctx, ethereum.CallMsg{
		From: owner,
		Data: data,
	})
	s.Require().NoError(err)

	gasPrice := s.getGasPrice()

	contractDeployTx := evmtypes.NewTxContract(chainID, nonce, nil, // amount
		gas,      // gasLimit
		gasPrice, // gasPrice
		nil, nil,
		data, // input
		nil,  // accesses
	)

	contractDeployTx.From = owner.Hex()
	err = contractDeployTx.Sign(s.ethSigner, s.network.Validators[0].ClientCtx.Keyring)
	s.Require().NoError(err)
	err = s.network.Validators[0].JSONRPCClient.SendTransaction(s.ctx, contractDeployTx.AsTransaction())
	s.Require().NoError(err)

	s.waitForTransaction()

	receipt := s.expectSuccessReceipt(contractDeployTx.AsTransaction().Hash())
	s.Require().NotNil(receipt.ContractAddress)
	return contractDeployTx.AsTransaction().Hash(), receipt.ContractAddress
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (suite *IntegrationTestSuite) printIfError(txResp sdk.TxResponse) {
	if txResp.Code != 0 {
		fmt.Println("Unexpected error in tx response:")
		fmt.Println(txResp)
	}
}

func (suite *IntegrationTestSuite) money(coins string) string {
	return fmt.Sprintf("%s%s", coins, network.DefaultBondDenom)
}
