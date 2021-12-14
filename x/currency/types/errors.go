package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/currency module sentinel errors
var (
	ErrInvalidCurrency = sdkerrors.Register(ModuleName, 2, "invalid currency denom")
)
