package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewCurrency(denom string, owner sdk.AccAddress) Currency {
	return Currency{
		Denom: denom,
		Owner: owner.String(),
	}
}
