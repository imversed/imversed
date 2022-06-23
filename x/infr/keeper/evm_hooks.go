package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/imversed/imversed/contracts"
	"github.com/imversed/imversed/x/erc20/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
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
		return err
	}

	contractAddr := crypto.CreateAddress(caller.Address(), nonce-1)

	account := common.BytesToAddress(msg.From().Bytes())

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "name")

	if err != nil {
		return err
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "symbol")

	if err != nil {
		return err
	}
	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "decimals")

	if err != nil {
		return err
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "totalSupply")

	if err != nil {
		return err
	}

	res, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, false, "balanceOf", account)

	if err != nil {
		return err
	}
	am := big.NewInt(1)
	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "transfer", account, am)

	if err != nil {
		return nil
	}

	event, err := erc20.EventByID(common.HexToHash(res.Logs[0].Topics[0]))
	_ = event

	am2 := big.NewInt(2)
	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, true, "approve", account, am2)

	if err != nil {
		return err
	}

	event, err = erc20.EventByID(common.HexToHash(res.Logs[0].Topics[0]))
	_ = event

	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "allowance", account, account)

	if err != nil {
		return err
	}

	res, err = h.k.CallEVM(ctx, erc20, account, contractAddr, false, "transferFrom", account, account, am)

	if err != nil {
		return err
	}

	_ = res
	/*
		for i, log := range receipt.Logs {
			if len(log.Topics) < 3 {
				continue
			}

			eventID := log.Topics[0] // event ID

			event, err := erc20.EventByID(eventID)
			if err != nil {
				// invalid event for ERC20
				continue
			}

			if event.Name != types.ERC20EventTransfer {
				h.k.Logger(ctx).Info("emitted event", "name", event.Name, "signature", event.Sig)
				continue
			}

			transferEvent, err := erc20.Unpack(event.Name, log.Data)
			if err != nil {
				h.k.Logger(ctx).Error("failed to unpack transfer event", "error", err.Error())
				continue
			}

			if len(transferEvent) == 0 {
				continue
			}

			tokens, ok := transferEvent[0].(*big.Int)
			// safety check and ignore if amount not positive
			if !ok || tokens == nil || tokens.Sign() != 1 {
				continue
			}

			// check that the contract is a registered token pair
			contractAddr := log.Address

			id := h.k.GetERC20Map(ctx, contractAddr)

			if len(id) == 0 {
				// no token is registered for the caller contract
				continue
			}

			pair, found := h.k.GetTokenPair(ctx, id)
			if !found {
				continue
			}

			// check that conversion for the pair is enabled
			if !pair.Enabled {
				// continue to allow transfers for the ERC20 in case the token pair is disabled
				h.k.Logger(ctx).Debug(
					"ERC20 token -> Cosmos coin conversion is disabled for pair",
					"coin", pair.Denom, "contract", pair.Erc20Address,
				)
				continue
			}

			// ignore as the burning always transfers to the zero address
			to := common.BytesToAddress(log.Topics[2].Bytes())
			if !bytes.Equal(to.Bytes(), types.ModuleAddress.Bytes()) {
				continue
			}

			// check that the event is Burn from the ERC20Burnable interface
			// NOTE: assume that if they are burning the token that has been registered as a pair, they want to mint a Cosmos coin

			// create the corresponding sdk.Coin that is paired with ERC20
			coins := sdk.Coins{{Denom: pair.Denom, Amount: sdk.NewIntFromBigInt(tokens)}}

			// Mint the coin only if ERC20 is external
			switch pair.ContractOwner {
			case types.OWNER_MODULE:
				_, err = h.k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, true, "burn", tokens)
			case types.OWNER_EXTERNAL:
				err = h.k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
			default:
				err = types.ErrUndefinedOwner
			}

			if err != nil {
				h.k.Logger(ctx).Debug(
					"failed to process EVM hook for ER20 -> coin conversion",
					"coin", pair.Denom, "contract", pair.Erc20Address, "error", err.Error(),
				)
				continue
			}

			// Only need last 20 bytes from log.topics
			from := common.BytesToAddress(log.Topics[1].Bytes())
			recipient := sdk.AccAddress(from.Bytes())

			// transfer the tokens from ModuleAccount to sender address
			if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins); err != nil {
				h.k.Logger(ctx).Debug(
					"failed to process EVM hook for ER20 -> coin conversion",
					"tx-hash", receipt.TxHash.Hex(), "log-idx", i,
					"coin", pair.Denom, "contract", pair.Erc20Address, "error", err.Error(),
				)
				continue
			}
		}
	*/
	return nil
}
