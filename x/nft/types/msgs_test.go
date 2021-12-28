package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fulldivevr/imversed/x/nft/types"
)
var denomE = types.NewDenom(denomID, denom, schema, denomSymbol, address, false, false, oracleUrl)
// ---------------------------------------- Msgs --------------------------------------------------

func TestMsgTransferNFTValidateBasicMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferNFT(denomID, "", id, tokenURI, tokenData, address.String(), address2.String())
	err := newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferNFT(denomID, denom, "", tokenURI, tokenData, "", address2.String())
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferNFT(denomID, denom, "", tokenURI, tokenData, address.String(), "")
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferNFT(denomID, denom, id, tokenURI, tokenData, address.String(), address2.String())
	err = newMsgTransferNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestUpdateDenomValidationBasicMethod(t *testing.T) {
	newMsgUpdateDenom := types.NewMsgUpdateDenom(denomID, denom, schema, address.String(), false, false, oracleUrl)
	err := newMsgUpdateDenom.ValidateBasic()
	require.NoError(t, err)
}

func TestIssueDenomValidationBasicMethod(t *testing.T) {
	newMsgIssueDenom := types.NewMsgIssueDenom(denomID,denom,schema, address.String(), denomSymbol, false, false, oracleUrl)
	err := newMsgIssueDenom.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgTransferNFTGetSignBytesMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferNFT(denomID, denom, id, tokenURI, tokenData, address.String(), address2.String())
	sortedBytes := newMsgTransferNFT.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgTransferNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"denom","name":"id1","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgIssueDenomGetSignBytesMethod(t *testing.T) {
	newMsgIssueDenom := types.NewMsgIssueDenom(denomID, denom, schema, address.String(), denomSymbol, false, false, oracleUrl)
	sortedBytes := newMsgIssueDenom.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgIssueDenom","value":{"id":"denom","name":"denom","oracle_url":"https://oracle-url.com","schema":"https://schema-url.com","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","symbol":"denomSymbol"}}`
	require.Equal(t, expected, string(sortedBytes))
}
func TestMsgUpdateDenomGetSignBytesMethod(t *testing.T) {
	newMsgUpdateDenom := types.NewMsgUpdateDenom(denomID, denom, schema, address.String(), false, false, oracleUrl)
	sortedBytes := newMsgUpdateDenom.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgUpdateDenom","value":{"id":"denom","name":"denom","oracle_url":"https://oracle-url.com","schema":"https://schema-url.com","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgIssueDenomGetSignersMethod(t * testing.T) {
	newMsgIssueDenom := types.NewMsgIssueDenom(denomID, denom, schema, address.String(), denomSymbol, false, false, oracleUrl)
	signers := newMsgIssueDenom.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}
func TestMsgUpdateDenomGetSignersMethod(t * testing.T) {
	newMsgUpdateDenom := types.NewMsgUpdateDenom(denomID, denom, schema, address.String(), false, false, oracleUrl)
	signers := newMsgUpdateDenom.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestMsgTransferNFTGetSignersMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferNFT(denomID, denom, id, tokenURI, tokenData, address.String(), address2.String())
	signers := newMsgTransferNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestMsgEditNFTValidateBasicMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditNFT(id, denom, nftName, tokenURI, tokenData, "")

	err := newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditNFT("", denom, nftName, tokenURI, tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditNFT(id, "", nftName, tokenURI, tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditNFT(id, denom, nftName, tokenURI, tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgEditNFTGetSignBytesMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditNFT(id, denom, nftName, tokenURI, tokenData, address.String())
	sortedBytes := newMsgEditNFT.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgEditNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"id1","name":"report","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgEditNFTGetSignersMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditNFT(id, denom, nftName, tokenURI, tokenData, address.String())
	signers := newMsgEditNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestMsgMsgMintNFTValidateBasicMethod(t *testing.T) {
	newMsgMintNFT := types.NewMsgMintNFT(id, denom, nftName, tokenURI, tokenData, "", address2.String())
	err := newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgMintNFT("", denom, nftName, tokenURI, tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgMintNFT(id, "", nftName, tokenURI, tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgMintNFT(id, denom, nftName, tokenURI, tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgMintNFTGetSignBytesMethod(t *testing.T) {
	newMsgMintNFT := types.NewMsgMintNFT(id, denom, nftName, tokenURI, tokenData, address.String(), address2.String())
	sortedBytes := newMsgMintNFT.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgMintNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"id1","name":"report","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgMsgBurnNFTValidateBasicMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgBurnNFT("", id, denom)
	err := newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgBurnNFT(address.String(), "", denom)
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgBurnNFT(address.String(), id, "")
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgBurnNFT(address.String(), id, denom)
	err = newMsgBurnNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgBurnNFTGetSignBytesMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgBurnNFT(address.String(), id, denom)
	sortedBytes := newMsgBurnNFT.GetSignBytes()
	expected := `{"type":"imversed/nft/MsgBurnNFT","value":{"denom_id":"denom","id":"id1","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgBurnNFTGetSignersMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgBurnNFT(address.String(), id, denom)
	signers := newMsgBurnNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}
