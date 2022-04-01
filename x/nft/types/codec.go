package types

// DONTCOVER

import (
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/imversed/imversed/x/nft/exported"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "imversed/nft/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgUpdateDenom{}, "imversed/nft/MsgUpdateDenom", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "imversed/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(&MsgEditNFT{}, "imversed/nft/MsgEditNFT", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "imversed/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "imversed/nft/MsgBurnNFT", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "imversed/nft/MsgTransferDenom", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "imversed/nft/BaseNFT", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgUpdateDenom{},
		&MsgTransferNFT{},
		&MsgEditNFT{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
		&MsgTransferDenom{},
	)

	registry.RegisterImplementations((*exported.NFT)(nil),
		&BaseNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.BinaryCodec, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyWrap)
}

// return th supply
func MustUnMarshalSupply(cdc codec.BinaryCodec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

// return the tokenID protobuf code
func MustMarshalTokenID(cdc codec.BinaryCodec, tokenID string) []byte {
	tokenIDWrap := gogotypes.StringValue{Value: tokenID}
	return cdc.MustMarshal(&tokenIDWrap)
}

// return th tokenID
func MustUnMarshalTokenID(cdc codec.Codec, value []byte) string {
	var tokenIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &tokenIDWrap)
	return tokenIDWrap.Value
}
