package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/verse module sentinel errors
var (
	ErrVerseAlreadyExists    = sdkerrors.Register(ModuleName, 3, "verse already exists")
	ErrContractAlreadyMapped = sdkerrors.Register(ModuleName, 4, "contract already mapped")
	ErrContractNotFound      = sdkerrors.Register(ModuleName, 5, "contract not found")
)
