package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/imversed/imversed/x/infr/client/cli"
	"github.com/imversed/imversed/x/infr/client/rest"
)

var (
	ChangeMinGasPricesProposalHandler = govclient.NewProposalHandler(cli.NewChangeMinGasPricesProposalCmd, rest.ChangeMinGasPricesRESTHandler)
)
