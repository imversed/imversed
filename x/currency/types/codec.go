package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCurrency{}, "currency/CreateCurrency", nil)
	cdc.RegisterConcrete(&MsgUpdateCurrency{}, "currency/UpdateCurrency", nil)
	cdc.RegisterConcrete(&MsgDeleteCurrency{}, "currency/DeleteCurrency", nil)
	cdc.RegisterConcrete(&MsgIssue{}, "currency/Issue", nil)
	cdc.RegisterConcrete(&MsgMint{}, "currency/Mint", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCurrency{},
		&MsgUpdateCurrency{},
		&MsgDeleteCurrency{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssue{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)