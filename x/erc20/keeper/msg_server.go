package keeper

import (
	"context"
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/imversed/imversed/contracts"
	"github.com/imversed/imversed/x/erc20/types"
)

var _ types.MsgServer = &Keeper{}

// ConvertCoin converts Cosmos-native Coins into ERC20 tokens for both
// Cosmos-native and ERC20 TokenPair Owners
func (k Keeper) ConvertCoin(
	goCtx context.Context,
	msg *types.MsgConvertCoin,
) (*types.MsgConvertCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Error checked during msg validation
	receiver := common.HexToAddress(msg.Receiver)
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)

	pair, err := k.MintingEnabled(ctx, sender, receiver.Bytes(), msg.Coin.Denom)
	if err != nil {
		return nil, err
	}

	// Remove token pair if contract is suicided
	erc20 := common.HexToAddress(pair.Erc20Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Erc20Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// Check ownership and execute conversion
	switch {
	case pair.IsNativeCoin():
		return k.convertCoinNativeCoin(ctx, pair, msg, receiver, sender) // case 1.1
	case pair.IsNativeERC20():
		return k.convertCoinNativeERC20(ctx, pair, msg, receiver, sender) // case 2.2
	default:
		return nil, types.ErrUndefinedOwner
	}
}

// ConvertERC20 converts ERC20 tokens into Cosmos-native Coins for both
// Cosmos-native and ERC20 TokenPair Owners
func (k Keeper) ConvertERC20(
	goCtx context.Context,
	msg *types.MsgConvertERC20,
) (*types.MsgConvertERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Error checked during msg validation
	receiver, _ := sdk.AccAddressFromBech32(msg.Receiver)
	sender := common.HexToAddress(msg.Sender)

	pair, err := k.MintingEnabled(ctx, sender.Bytes(), receiver, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	// Remove token pair if contract is suicided
	erc20 := common.HexToAddress(pair.Erc20Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Erc20Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// Check ownership
	switch {
	case pair.IsNativeCoin():
		return k.convertERC20NativeCoin(ctx, pair, msg, receiver, sender) // case 1.2
	case pair.IsNativeERC20():
		return k.convertERC20NativeToken(ctx, pair, msg, receiver, sender) // case 2.1
	default:
		return nil, types.ErrUndefinedOwner
	}
}

// convertCoinNativeCoin handles the Coin conversion flow for a native coin
// token pair:
//  - Escrow Coins on module account (Coins are not burned)
//  - Mint Tokens and send to receiver
//  - Check if token balance increased by amount
func (k Keeper) convertCoinNativeCoin(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertCoin,
	receiver common.Address,
	sender sdk.AccAddress,
) (*types.MsgConvertCoinResponse, error) {
	// NOTE: ignore validation from NewCoin constructor
	coins := sdk.Coins{msg.Coin}
	erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
	contract := pair.GetERC20Contract()
	balanceToken := k.balanceOf(ctx, erc20, contract, receiver)

	// Escrow Coins on module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to escrow coins")
	}

	// Mint Tokens and send to receiver
	_, err := k.CallEVM(ctx, erc20, types.ModuleAddress, contract, true, "mint", receiver, msg.Coin.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// Check expected Receiver balance after transfer execution
	tokens := msg.Coin.Amount.BigInt()
	balanceTokenAfter := k.balanceOf(ctx, erc20, contract, receiver)
	exp := big.NewInt(0).Add(balanceToken, tokens)

	if r := balanceTokenAfter.Cmp(exp); r != 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", exp, balanceTokenAfter,
		)
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertCoin,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
			),
		},
	)

	return &types.MsgConvertCoinResponse{}, nil
}

// convertERC20NativeCoin handles the erc20 conversion flow for a native coin token pair:
//  - Burn escrowed tokens
//  - Unescrow coins that have been previously escrowed with ConvertCoin
//  - Check if coin balance increased by amount
//  - Check if token balance decreased by amount
func (k Keeper) convertERC20NativeCoin(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC20,
	receiver sdk.AccAddress,
	sender common.Address,
) (*types.MsgConvertERC20Response, error) {
	// NOTE: coin fields already validated
	coins := sdk.Coins{sdk.Coin{Denom: pair.Denom, Amount: msg.Amount}}

	erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
	contract := pair.GetERC20Contract()
	balanceCoin := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceToken := k.balanceOf(ctx, erc20, contract, sender)

	// Burn escrowed tokens
	_, err := k.CallEVM(ctx, erc20, types.ModuleAddress, contract, true, "burnCoins", sender, msg.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// Unescrow Coins and send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, coins); err != nil {
		return nil, err
	}

	// Check expected Receiver balance after transfer execution
	balanceCoinAfter := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	expCoin := balanceCoin.Add(coins[0])
	if ok := balanceCoinAfter.IsEqual(expCoin); !ok {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid coin balance - expected: %v, actual: %v",
			expCoin, balanceCoinAfter,
		)
	}

	// Check expected Sender balance after transfer execution
	tokens := coins[0].Amount.BigInt()
	balanceTokenAfter := k.balanceOf(ctx, erc20, contract, sender)
	expToken := big.NewInt(0).Sub(balanceToken, tokens)
	if r := balanceTokenAfter.Cmp(expToken); r != 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v",
			expToken, balanceTokenAfter,
		)
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC20,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, msg.ContractAddress),
			),
		},
	)

	return &types.MsgConvertERC20Response{}, nil
}

// convertERC20NativeToken handles the erc20 conversion flow for a native erc20 token pair:
//  - Escrow tokens on module account (Don't burn as module is not contract owner)
//  - Mint coins on module
//  - Send minted coins to the receiver
//  - Check if coin balance increased by amount
//  - Check if token balance decreased by amount
//  - Check for unexpected `appove` event in logs
func (k Keeper) convertERC20NativeToken(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC20,
	receiver sdk.AccAddress,
	sender common.Address,
) (*types.MsgConvertERC20Response, error) {
	// NOTE: coin fields already validated
	coins := sdk.Coins{sdk.Coin{Denom: pair.Denom, Amount: msg.Amount}}
	erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
	contract := pair.GetERC20Contract()
	balanceCoin := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceToken := k.balanceOf(ctx, erc20, contract, types.ModuleAddress)

	// Escrow tokens on module account
	transferData, err := erc20.Pack("transfer", types.ModuleAddress, msg.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	res, err := k.CallEVMWithData(ctx, sender, &contract, transferData, true)
	if err != nil {
		return nil, err
	}

	// Check unpackedRet execution
	var unpackedRet types.ERC20BoolResponse
	if err := erc20.UnpackIntoInterface(&unpackedRet, "transfer", res.Ret); err != nil {
		return nil, err
	}

	if !unpackedRet.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to execute transfer")
	}

	// Check expected escrow balance after transfer execution
	tokens := coins[0].Amount.BigInt()
	balanceTokenAfter := k.balanceOf(ctx, erc20, contract, types.ModuleAddress)
	expToken := big.NewInt(0).Add(balanceToken, tokens)

	if r := balanceTokenAfter.Cmp(expToken); r != 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v",
			expToken, balanceTokenAfter,
		)
	}

	// Mint coins
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, err
	}

	// Send minted coins to the receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, coins); err != nil {
		return nil, err
	}

	// Check expected Receiver balance after transfer execution
	balanceCoinAfter := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	expCoin := balanceCoin.Add(coins[0])

	if ok := balanceCoinAfter.IsEqual(expCoin); !ok {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid coin balance - expected: %v, actual: %v",
			expCoin, balanceCoinAfter,
		)
	}

	// Check for unexpected `appove` event in logs
	if err := k.monitorApprovalEvent(res); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC20,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, msg.ContractAddress),
			),
		},
	)

	return &types.MsgConvertERC20Response{}, nil
}

// convertCoinNativeERC20 handles the Coin conversion flow for a native ERC20
// token pair:
//  - Escrow Coins on module account
//  - Unescrow Tokens that have been previously escrowed with ConvertERC20 and send to receiver
//  - Burn escrowed Coins
//  - Check if token balance increased by amount
//  - Check for unexpected `appove` event in logs
func (k Keeper) convertCoinNativeERC20(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertCoin,
	receiver common.Address,
	sender sdk.AccAddress,
) (*types.MsgConvertCoinResponse, error) {
	// NOTE: ignore validation from NewCoin constructor
	coins := sdk.Coins{msg.Coin}

	erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
	contract := pair.GetERC20Contract()
	balanceToken := k.balanceOf(ctx, erc20, contract, receiver)

	// Escrow Coins on module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to escrow coins")
	}

	// Unescrow Tokens and send to receiver
	res, err := k.CallEVM(ctx, erc20, types.ModuleAddress, contract, true, "transfer", receiver, msg.Coin.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// Check unpackedRet execution
	var unpackedRet types.ERC20BoolResponse
	if err := erc20.UnpackIntoInterface(&unpackedRet, "transfer", res.Ret); err != nil {
		return nil, err
	}

	if !unpackedRet.Value {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "failed to execute unescrow tokens from user")
	}

	// Check expected Receiver balance after transfer execution
	tokens := msg.Coin.Amount.BigInt()
	balanceTokenAfter := k.balanceOf(ctx, erc20, contract, receiver)
	exp := big.NewInt(0).Add(balanceToken, tokens)

	if r := balanceTokenAfter.Cmp(exp); r != 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", exp, balanceTokenAfter,
		)
	}

	// Burn escrowed Coins
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to burn coins")
	}

	// Check for unexpected `appove` event in logs
	if err := k.monitorApprovalEvent(res); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertCoin,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
			),
		},
	)

	return &types.MsgConvertCoinResponse{}, nil
}

// balanceOf queries an account's balance for a given ERC20 contract
func (k Keeper) balanceOf(
	ctx sdk.Context,
	abi abi.ABI,
	contract, account common.Address,
) *big.Int {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, "balanceOf", account)
	if err != nil {
		return nil
	}

	unpacked, err := abi.Unpack("balanceOf", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return nil
	}

	balance, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil
	}

	return balance
}

// monitorApprovalEvent returns an error if the given transactions logs include
// an unexpected `approve` event
func (k Keeper) monitorApprovalEvent(res *evmtypes.MsgEthereumTxResponse) error {
	if res == nil || len(res.Logs) == 0 {
		return nil
	}

	logApprovalSig := []byte("Approval(address,address,uint256)")
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)

	for _, log := range res.Logs {
		if log.Topics[0] == logApprovalSigHash.Hex() {
			return sdkerrors.Wrapf(
				types.ErrUnexpectedEvent, "unexpected Approval event",
			)
		}
	}

	return nil
}

// RegisterCoin deploys an erc20 contract and creates the token pair for the existing cosmos coin
func (k Keeper) RegisterCoin(goCtx context.Context, msg *types.MsgRegisterCoin) (*types.MsgRegisterCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coinMetadata := msg.Metadata
	// check if the conversion is globally enabled
	params := k.GetParams(ctx)
	if !params.EnableErc20 {
		return nil, sdkerrors.Wrap(types.ErrERC20Disabled, "registration is currently disabled by governance")
	}
	/*
		// prohibit denominations that contain the evm denom
		if strings.Contains(coinMetadata.Base, "evm") {
			return nil, sdkerrors.Wrapf(types.ErrEVMDenom, "cannot register the EVM denomination %s", coinMetadata.Base)
		}
	*/
	// check if the denomination already registered
	if k.IsDenomRegistered(ctx, coinMetadata.Name) {
		return nil, sdkerrors.Wrapf(types.ErrTokenPairAlreadyExists, "coin denomination already registered: %s", coinMetadata.Name)
	}

	// check if the coin exists by ensuring the supply is set
	if !k.bankKeeper.HasSupply(ctx, coinMetadata.Base) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidCoins,
			"base denomination '%s' cannot have a supply of 0", coinMetadata.Base,
		)
	}

	if err := k.verifyMetadata(ctx, coinMetadata); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "coin metadata is invalid %s", coinMetadata.Name)
	}

	addr, err := k.DeployERC20Contract(ctx, coinMetadata)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to create wrapped coin denom metadata for ERC20")
	}

	pair := types.NewTokenPair(addr, coinMetadata.Base, true, types.OWNER_MODULE, msg.Sender)
	k.SetTokenPair(ctx, pair)
	k.SetDenomMap(ctx, pair.Denom, pair.GetID())
	k.SetERC20Map(ctx, common.HexToAddress(pair.Erc20Address), pair.GetID())

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterCoin,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
			sdk.NewAttribute(types.AttributeKeyERC20Token, pair.Erc20Address),
		),
	)

	return &types.MsgRegisterCoinResponse{}, nil
}

// DeployERC20Contract creates and deploys an ERC20 contract on the EVM with the
// erc20 module account as owner.
func (k Keeper) DeployERC20Contract(
	ctx sdk.Context,
	coinMetadata banktypes.Metadata,
) (common.Address, error) {
	decimals := uint8(coinMetadata.DenomUnits[0].Exponent)
	ctorArgs, err := contracts.ERC20MinterBurnerDecimalsContract.ABI.Pack(
		"",
		coinMetadata.Name,
		coinMetadata.Symbol,
		decimals,
	)
	if err != nil {
		return common.Address{}, sdkerrors.Wrapf(types.ErrABIPack, "coin metadata is invalid %s: %s", coinMetadata.Name, err.Error())
	}

	data := make([]byte, len(contracts.ERC20MinterBurnerDecimalsContract.Bin)+len(ctorArgs))
	copy(data[:len(contracts.ERC20MinterBurnerDecimalsContract.Bin)], contracts.ERC20MinterBurnerDecimalsContract.Bin)
	copy(data[len(contracts.ERC20MinterBurnerDecimalsContract.Bin):], ctorArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, sdkerrors.Wrapf(err, "failed to deploy contract for %s", coinMetadata.Name)
	}

	return contractAddr, nil
}

// RegisterERC20 creates a cosmos coin and registers the token pair between the coin and the ERC20
func (k Keeper) RegisterERC20(goCtx context.Context, msg *types.MsgRegisterERC20) (*types.MsgRegisterERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract := common.HexToAddress(msg.Erc20Address)
	params := k.GetParams(ctx)
	if !params.EnableErc20 {
		return nil, sdkerrors.Wrap(types.ErrERC20Disabled, "registration is currently disabled by governance")
	}

	if k.IsERC20Registered(ctx, contract) {
		return nil, sdkerrors.Wrapf(types.ErrTokenPairAlreadyExists, "token ERC20 contract already registered: %s", contract.String())
	}

	metadata, err := k.CreateCoinMetadata(ctx, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to create wrapped coin denom metadata for ERC20")
	}

	pair := types.NewTokenPair(contract, metadata.Name, true, types.OWNER_EXTERNAL, msg.Sender)
	k.SetTokenPair(ctx, pair)
	k.SetDenomMap(ctx, pair.Denom, pair.GetID())
	k.SetERC20Map(ctx, common.HexToAddress(pair.Erc20Address), pair.GetID())
	return &types.MsgRegisterERC20Response{}, nil
}

// CreateCoinMetadata generates the metadata to represent the ERC20 token on evmos.
func (k Keeper) CreateCoinMetadata(ctx sdk.Context, contract common.Address) (*banktypes.Metadata, error) {
	strContract := contract.String()

	erc20Data, err := k.QueryERC20(ctx, contract)
	if err != nil {
		return nil, err
	}

	_, found := k.bankKeeper.GetDenomMetaData(ctx, types.CreateDenom(strContract))
	if found {
		// metadata already exists; exit
		return nil, sdkerrors.Wrap(types.ErrInternalTokenPair, "denom metadata already registered")
	}

	if k.IsDenomRegistered(ctx, types.CreateDenom(strContract)) {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "coin denomination already registered: %s", erc20Data.Name)
	}

	// base denomination
	base := types.CreateDenom(strContract)

	// create a bank denom metadata based on the ERC20 token ABI details
	// metadata name is should always be the contract since it's the key
	// to the bank store
	metadata := banktypes.Metadata{
		Description: types.CreateDenomDescription(strContract),
		Base:        base,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    base,
				Exponent: 0,
			},
		},
		Name:    types.CreateDenom(strContract),
		Symbol:  erc20Data.Symbol,
		Display: base,
	}

	// only append metadata if decimals > 0, otherwise validation fails
	if erc20Data.Decimals > 0 {
		nameSanitized := types.SanitizeERC20Name(erc20Data.Name)
		metadata.DenomUnits = append(
			metadata.DenomUnits,
			&banktypes.DenomUnit{
				Denom:    nameSanitized,
				Exponent: uint32(erc20Data.Decimals),
			},
		)
		metadata.Display = nameSanitized
	}

	if err := metadata.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(err, "ERC20 token data is invalid for contract %s", strContract)
	}

	k.bankKeeper.SetDenomMetaData(ctx, metadata)

	return &metadata, nil
}

// UpdateTokenPairERC20 updates the ERC20 token address for the registered token pair
func (k Keeper) UpdateTokenPairERC20(
	goCtx context.Context,
	msg *types.MsgUpdateTokenPairERC20,
) (*types.MsgUpdateTokenPairERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	fmt.Println(sender)
	erc20Addr := common.HexToAddress(msg.Erc20Address)
	id := k.GetERC20Map(ctx, erc20Addr)
	if len(id) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "token %s not registered", erc20Addr)
	}

	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTokenPairNotFound, "token '%s' not registered", erc20Addr)
	}
	owner, err := sdk.AccAddressFromBech32(pair.AccountOwner)

	if !sender.Equals(owner) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidTokenPairAccountOwner, "sender is not token pair account owner")
	}

	// Get current stored metadata
	metadata, found := k.bankKeeper.GetDenomMetaData(ctx, pair.Denom)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "could not get metadata for %s", pair.Denom)
	}

	// safety check
	if len(metadata.DenomUnits) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "metadata denom units for %s cannot be empty", pair.Erc20Address)
	}

	// Get new erc20 values
	newERC20Addr := common.HexToAddress(msg.NewErc20Address)
	erc20Data, err := k.QueryERC20(ctx, newERC20Addr)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "could not get token %s erc20Data", newERC20Addr.String())
	}

	// compare metadata and ERC20 details
	if metadata.Display != erc20Data.Name ||
		metadata.Symbol != erc20Data.Symbol ||
		metadata.Description != types.CreateDenomDescription(erc20Addr.String()) {
		return nil, sdkerrors.Wrapf(types.ErrInternalTokenPair, "metadata details (display, symbol, description) don't match the ERC20 details from %s ", pair.Erc20Address)
	}

	// check that the denom units contain one item with the same
	// name and decimal values as the ERC20
	found = false
	for _, denomUnit := range metadata.DenomUnits {
		// iterate denom units until we found the one with the ERC20 Name
		if denomUnit.Denom != erc20Data.Name {
			continue
		}

		// once found, check it has the same exponent
		if denomUnit.Exponent != uint32(erc20Data.Decimals) {
			return nil, sdkerrors.Wrapf(
				types.ErrInternalTokenPair, "metadata denom unit exponent doesn't match the ERC20 details from %s, expected %d, got %d",
				pair.Erc20Address, erc20Data.Decimals, denomUnit.Exponent,
			)
		}

		// break as metadata might contain denom units for higher exponents
		found = true
		break
	}

	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrInternalTokenPair,
			"metadata doesn't contain denom unit found for ERC20 %s (%s)",
			erc20Data.Name, pair.Erc20Address,
		)
	}

	// Update the metadata description with the new address
	metadata.Description = types.CreateDenomDescription(newERC20Addr.String())
	k.bankKeeper.SetDenomMetaData(ctx, metadata)
	// Delete old token pair (id is changed because the ERC20 address was modifed)
	k.DeleteTokenPair(ctx, pair)
	// Update the address
	pair.Erc20Address = newERC20Addr.Hex()
	newID := pair.GetID()
	// Set the new pair
	k.SetTokenPair(ctx, pair)
	// Overwrite the value because id was changed
	k.SetDenomMap(ctx, pair.Denom, newID)
	// Add the new address
	k.SetERC20Map(ctx, newERC20Addr, newID)

	return &types.MsgUpdateTokenPairERC20Response{}, nil
}

// ToggleTokenRelay toggles relaying for a given token pair
func (k Keeper) ToggleTokenRelay(
	goCtx context.Context,
	msg *types.MsgToggleTokenRelay,
) (*types.MsgToggleTokenRelayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)

	id := k.GetTokenPairID(ctx, msg.Token)
	if len(id) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrTokenPairNotFound, "token '%s' not registered by id", msg.Token)
	}

	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTokenPairNotFound, "token '%s' not registered", msg.Token)
	}

	owner, err := sdk.AccAddressFromBech32(pair.AccountOwner)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidTokenPairAccountOwner, "invalid account owner", pair.AccountOwner)
	}

	if !sender.Equals(owner) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidTokenPairAccountOwner, "sender is not token pair account owner")
	}

	pair.Enabled = !pair.Enabled

	k.SetTokenPair(ctx, pair)
	return &types.MsgToggleTokenRelayResponse{}, nil
}

// verify if the metadata matches the existing one, if not it sets it to the store
func (k Keeper) verifyMetadata(ctx sdk.Context, coinMetadata banktypes.Metadata) error {
	meta, found := k.bankKeeper.GetDenomMetaData(ctx, coinMetadata.Base)
	if !found {
		k.bankKeeper.SetDenomMetaData(ctx, coinMetadata)
		return nil
	}
	// If it already existed, Check that is equal to what is stored
	return types.EqualMetadata(meta, coinMetadata)
}
