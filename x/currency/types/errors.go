package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/currency module sentinel errors
var (
	ErrInvalidCurrency = sdkerrors.Register(ModuleName, 2, "invalid currency denom")
	ErrUnknownCurrency = sdkerrors.Register(ModuleName, 3, "unknown currency denom")
	ErrInvalidOwner    = sdkerrors.Register(ModuleName, 4, "invalid owner")
)
