package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/imversed/imversed/x/erc20/types"
)

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
