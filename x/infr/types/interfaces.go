package types

import (
	"context"
	accounttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	erc20types "github.com/imversed/imversed/x/erc20/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(sdk.Context, sdk.AccAddress) accounttypes.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}

type Erc20Keeper interface {
	RegisterERC20(goCtx context.Context, msg *erc20types.MsgRegisterERC20) (*erc20types.MsgRegisterERC20Response, error)
}

type BankKeeper interface {
	authtypes.BankKeeper
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}
