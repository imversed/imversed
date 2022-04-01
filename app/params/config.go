package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	HumanCoinUnit = "imv"
	BaseCoinUnit  = "nimv"
	OsmoExponent  = 6

	DefaultBondDenom = BaseCoinUnit

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = "imv"

	MaxAddrLen = 32
)

// var (
// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
// Bech32PrefixAccPub = Bech32PrefixAccAddr + "pub"
// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
// Bech32PrefixValAddr = Bech32PrefixAccAddr + "valoper"
// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
// Bech32PrefixValPub = Bech32PrefixAccAddr + "valoperpub"
// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
// Bech32PrefixConsAddr = Bech32PrefixAccAddr + "valcons"
// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
// Bech32PrefixConsPub = Bech32PrefixAccAddr + "valconspub"
// )

func InitConfig() {
	SetAddressPrefixes()
	// RegisterDenoms()
}

func RegisterDenoms() {
	err := sdk.RegisterDenom(HumanCoinUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, OsmoExponent))
	if err != nil {
		panic(err)
	}
}

func SetAddressPrefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccAddr)
	config.SetBech32PrefixForValidator(Bech32PrefixAccAddr, Bech32PrefixAccAddr)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixAccAddr, Bech32PrefixAccAddr)

	config.SetAddressVerifier(func(bytes []byte) error {
		if len(bytes) == 0 {
			return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
		}

		// TODO: Do we want to allow addresses of lengths other than 20 and 32 bytes?
		if len(bytes) != 20 && len(bytes) != 32 {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address length must be 20 or 32 bytes, got %d", len(bytes))
		}

		return nil
	})
}
