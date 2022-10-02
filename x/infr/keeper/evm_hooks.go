package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/imversed/imversed/contracts"
	"github.com/imversed/imversed/x/erc20/types"
	erc20types "github.com/imversed/imversed/x/erc20/types"
	infrtypes "github.com/imversed/imversed/x/infr/types"
	"math/big"
)

type Hooks struct {
	k Keeper
}

var _ evmtypes.EvmHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) PostTxProcessing(
	gctx sdk.Context,
	msg core.Message,
	receipt *ethtypes.Receipt,
) error {
	// if not smart-contract deploying
	if msg.To() != nil {
		return nil
	}

	var res *evmtypes.MsgEthereumTxResponse

	ctx, _ := gctx.CacheContext()

	erc20 := contracts.ERC20BurnableContract.ABI

	caller := vm.AccountRef(msg.From())
	nonce, err := h.k.accountKeeper.GetSequence(ctx, msg.From().Bytes())

	if err != nil {
		return nil
	}

	contractAddr := crypto.CreateAddress(caller.Address(), nonce-1)

	account := common.BytesToAddress(msg.From().Bytes())

	// add metadata to new contract
	sc := infrtypes.SmartContract{
		Address:     contractAddr.Hex(),
		Creator:     account.Hex(),
		BlockNumber: ctx.BlockHeight(),
	}
	h.k.SetSmartContractMetadata(ctx, sc)

	// check erc20-compatibility
	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "name")
	if err != nil {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "symbol")
	if err != nil {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "decimals")
	if err != nil {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "totalSupply")
	if err != nil {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "balanceOf", account)
	if err != nil {
		return nil
	}

	am := big.NewInt(1)
	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "transfer", account, am)
	if err != nil {
		return nil
	}

	event, err := erc20.EventByID(common.HexToHash(res.Logs[0].Topics[0]))
	if err != nil {
		return nil
	}
	if event.RawName != "Transfer" {
		return nil
	}

	am2 := big.NewInt(2)
	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, true, "approve", account, am2)
	if err != nil {
		return nil
	}

	event, err = erc20.EventByID(common.HexToHash(res.Logs[0].Topics[0]))
	if event.RawName != "Approval" {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "allowance", account, account)
	if err != nil {
		return nil
	}

	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "transferFrom", account, account, am)
	if err != nil {
		return nil
	}

	event, err = erc20.EventByID(common.HexToHash(res.Logs[0].Topics[0]))
	if err != nil {
		return nil
	}
	if event.RawName != "Transfer" {
		return nil
	}

	cosmosAddress := h.k.accountKeeper.GetAccount(gctx, msg.From().Bytes()).GetAddress()
	_, err = h.k.erc20Keeper.RegisterERC20(sdk.WrapSDKContext(gctx), erc20types.NewMsgRegisterERC20(contractAddr.String(), cosmosAddress.String()))

	return nil
}
