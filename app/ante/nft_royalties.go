package ante

//
//import (
//	"errors"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//)
//
//type ValidateNFTRoyaltyDecorator struct {
//	nk nftkeeper.Keeper
//}
//
//func NewValidateNFTRoyaltyDecorator(nk nftkeeper.Keeper) ValidateNFTRoyaltyDecorator {
//	return ValidateNFTRoyaltyDecorator{
//		nk: nk,
//	}
//}
//
//func (vnr ValidateNFTRoyaltyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
//
//	msgs := tx.GetMsgs()
//
//	for i, msg := range msgs {
//		if _, ok := msg.(*nfttypes.MsgTransferNFT); ok {
//			if err := vnr.CheckRoyaltiesForNFT(ctx, msgs, i); err != nil {
//				return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, err.Error())
//			}
//		}
//	}
//
//	return next(ctx, tx, simulate)
//}
//
//func (vnr ValidateNFTRoyaltyDecorator) CheckRoyaltiesForNFT(ctx sdk.Context, msgs []sdk.Msg, i int) error {
//	msgTransfer := msgs[i].(*nfttypes.MsgTransferNFT)
//	denomId := msgTransfer.DenomId
//
//	denom, ok := vnr.nk.GetDenom(ctx, denomId)
//	if !ok {
//		return errors.New("denom not found")
//	}
//
//	recipient := msgTransfer.Recipient
//	sender := msgTransfer.Sender
//	creator := denom.Creator
//
//	flagPay, flagRoyalty := false, false
//
//	for _, msg := range msgs {
//		if bm, ok := msg.(*banktypes.MsgSend); ok {
//			if bm.FromAddress == recipient {
//				if !flagPay && bm.ToAddress == sender {
//					flagPay = true
//				}
//				if !flagRoyalty && bm.ToAddress == creator {
//					flagRoyalty = true
//				}
//			}
//			if flagPay && flagRoyalty {
//				break
//			}
//		}
//	}
//
//	if !flagPay || !flagRoyalty {
//		return errors.New("tx contains transfer messages, this also must contains 2 bank messages: transfer money and royalties")
//	}
//	return nil
//}
