package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/imversed/imversed/x/erc20/client/cli"
	"github.com/imversed/imversed/x/erc20/client/rest"
)

var (
	RegisterERC20ProposalHandler        = govclient.NewProposalHandler(cli.NewRegisterERC20ProposalCmd, rest.RegisterERC20ProposalRESTHandler)
	ToggleTokenRelayProposalHandler     = govclient.NewProposalHandler(cli.NewToggleTokenRelayProposalCmd, rest.ToggleTokenRelayRESTHandler)
)
