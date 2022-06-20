package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/imversed/imversed/x/erc20/client/cli"
	"github.com/imversed/imversed/x/erc20/client/rest"
)

var (
	ToggleTokenRelayProposalHandler = govclient.NewProposalHandler(cli.NewToggleTokenRelayProposalCmd, rest.ToggleTokenRelayRESTHandler)
)
