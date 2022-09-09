package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/imversed/imversed/x/infr/client/cli"
)

var (
	ChangeMinGasPricesProposalHandler = govclient.NewProposalHandler(cli.NewChangeMinGasPricesProposalCmd)
)
