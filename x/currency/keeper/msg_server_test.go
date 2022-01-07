package keeper_test

import (
	"github.com/fulldivevr/imversed/testutil/sample"
	"github.com/stretchr/testify/mock"
	"testing"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fulldivevr/imversed/testutil/keeper"
	"github.com/fulldivevr/imversed/x/currency/keeper"
	"github.com/fulldivevr/imversed/x/currency/types"
)

func setupMsgServer(t testing.TB, bank *BankMock) (types.MsgServer, sdk.Context) {
	k, ctx := keepertest.CurrencyKeeper(t, bank)
	return keeper.NewMsgServerImpl(*k), ctx
}

func setupDenom(srv types.MsgServer, ctx sdk.Context, sender string, denom string) error {
	wCtx := sdk.WrapSDKContext(ctx)
	if _, err := srv.Issue(wCtx, &types.MsgIssue{
		Sender: sender,
		Denom:  denom,
	}); err != nil {
		return err
	}
	return nil
}

func TestSendMintEvents(t *testing.T) {
	bankMock := new(BankMock)
	srv, ctx := setupMsgServer(t, bankMock)

	denomName := "IMV"
	accAddress := sample.AccAddress()

	if err := setupDenom(srv, ctx, accAddress, denomName); err != nil {
		t.Fatal(err)
	}

	// m.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Coin))
	bankMock.On("SendCoinsFromModuleToAccount", types.ModuleName, mock.Anything, mock.Anything).Return(nil)
	bankMock.On("MintCoins", types.ModuleName, mock.Anything).Return(nil)

	msgMint := new(types.MsgMint)
	msgMint.Sender = accAddress
	msgMint.Coin = cosmostypes.NewCoin(denomName, sdk.NewInt(100500))

	wContext := sdk.WrapSDKContext(ctx)

	if _, err := srv.Mint(wContext, msgMint); err != nil {
		t.Fatal(err)
	}

	ctx.EventManager().Events()
	for _, event := range ctx.EventManager().Events() {
		// TODO: make better checking
		println(event.Type)
		//fmt.Print(event)
	}
}

// ====================================================================================================

type BankMock struct {
	mock.Mock
}

func (b BankMock) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	res := b.Called(senderModule, recipientAddr, amt)
	return res.Error(0)
}

func (b BankMock) MintCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error {
	res := b.Called(moduleName, amounts)
	return res.Error(0)
}
