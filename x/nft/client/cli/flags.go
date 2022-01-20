package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName = "name"
	FlagTokenURI  = "uri"
	FlagTokenData = "data"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"

	FlagDenomName        = "name"
	FlagDenomID          = "denom-id"
	FlagSchema           = "schema"
	FlagSymbol           = "symbol"
	FlagMintRestricted   = "mint-restricted"
	FlagUpdateRestricted = "update-restricted"
	FlagOracleUrl		 = "oracle-url"
)

var (
	FsIssueDenom    = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateDenom   = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT       = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT       = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT   = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply   = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner    = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferDenom = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagSchema, "", "Denom data structure definition")
	FsIssueDenom.String(FlagDenomName, "", "The name of the denom")
	FsIssueDenom.String(FlagSymbol, "", "The symbol of the denom")
	FsIssueDenom.Bool(FlagMintRestricted, false, "mint restricted of nft under denom")
	FsIssueDenom.Bool(FlagUpdateRestricted, false, "update restricted of nft under denom")
	FsIssueDenom.String(FlagOracleUrl, "","The URL address of a trusted oracle")

	FsUpdateDenom.String(FlagDenomName, "", "The name of the denom")
	FsUpdateDenom.String(FlagSchema, "", "Denom data structure definition")
	FsUpdateDenom.Bool(FlagMintRestricted, false, "mint restricted of nft under denom")
	FsUpdateDenom.Bool(FlagUpdateRestricted, false, "update restricted of nft under denom")
	FsUpdateDenom.String(FlagOracleUrl, "","The URL address of a trusted oracle")

	FsMintNFT.String(FlagTokenURI, "", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagTokenData, "", "The origin data of the nft")
	FsMintNFT.String(FlagTokenName, "", "The name of the nft")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for the supplemental off-chain token data (should return a JSON object)")
	FsEditNFT.String(FlagTokenData, "[do-not-modify]", "The token data of the nft")
	FsEditNFT.String(FlagTokenName, "[do-not-modify]", "The name of the nft")

	FsTransferNFT.String(FlagTokenURI, "[do-not-modify]", "URI for the supplemental off-chain token data (should return a JSON object)")
	FsTransferNFT.String(FlagTokenData, "[do-not-modify]", "The token data of the nft")
	FsTransferNFT.String(FlagTokenName, "[do-not-modify]", "The name of the nft")

	FsQuerySupply.String(FlagOwner, "", "The owner of the nft")

	FsQueryOwner.String(FlagDenomID, "", "The name of the collection")
}
