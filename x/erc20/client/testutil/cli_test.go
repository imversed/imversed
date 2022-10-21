package testutil

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/imversed/imversed/x/infr/minGasPriceHelper"
	"strings"
	"testing"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/stretchr/testify/suite"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/imversed/imversed/testutil/network"

	evmosnetwork "github.com/imversed/imversed/testutil/network"
	"github.com/imversed/imversed/x/erc20/client/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	var err error
	s.cfg = evmosnetwork.DefaultConfig()
	minGasPriceHelper.Create(baseapp.SetMinGasPrices, s.cfg.MinGasPrices)
	s.cfg.NumValidators = 1

	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)
	s.Require().NotNil(s.network)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestCmdParams() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			"erc20 params",
			[]string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			`{"params":{"enable_erc20":true,"enable_evm_hook":true}}`,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetParamsCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			s.Require().Equal(strings.TrimSpace(tc.expectedOutput), strings.TrimSpace(out.String()))
		})
	}
}