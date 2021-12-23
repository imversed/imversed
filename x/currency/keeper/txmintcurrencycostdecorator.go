package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fulldivevr/imversed/x/currency/types"
)

type ConsumeTxMintCurrencyGasDecorator struct {
	k Keeper
}

func NewConsumeTxMintCurrencyGasDecorator(k Keeper) ConsumeTxMintCurrencyGasDecorator {
	return ConsumeTxMintCurrencyGasDecorator{
		k: k,
	}
}

func (d ConsumeTxMintCurrencyGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	params := d.k.GetParams(ctx)

	for _, msg := range tx.GetMsgs() {
		switch msg.(type) {
		case *types.MsgMint:
			ctx.GasMeter().ConsumeGas(params.TxMintCurrencyCost, "txMintCurrency")
		}
	}

	return next(ctx, tx, simulate)
}
