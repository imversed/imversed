package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/verse module sentinel errors
var (
	ErrInvalidVerse       = sdkerrors.Register(ModuleName, 2, "invalid verse")
	ErrVerseAlreadyExists = sdkerrors.Register(ModuleName, 3, "verse already exists")
)
