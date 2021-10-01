// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgMintNFT } from "./types/nft/tx";
import { MsgTransferDenom } from "./types/nft/tx";
import { MsgTransferNFT } from "./types/nft/tx";
import { MsgEditNFT } from "./types/nft/tx";
import { MsgIssueDenom } from "./types/nft/tx";
import { MsgBurnNFT } from "./types/nft/tx";
const types = [
    ["/metachain.nft.MsgMintNFT", MsgMintNFT],
    ["/metachain.nft.MsgTransferDenom", MsgTransferDenom],
    ["/metachain.nft.MsgTransferNFT", MsgTransferNFT],
    ["/metachain.nft.MsgEditNFT", MsgEditNFT],
    ["/metachain.nft.MsgIssueDenom", MsgIssueDenom],
    ["/metachain.nft.MsgBurnNFT", MsgBurnNFT],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgMintNFT: (data) => ({ typeUrl: "/metachain.nft.MsgMintNFT", value: data }),
        msgTransferDenom: (data) => ({ typeUrl: "/metachain.nft.MsgTransferDenom", value: data }),
        msgTransferNFT: (data) => ({ typeUrl: "/metachain.nft.MsgTransferNFT", value: data }),
        msgEditNFT: (data) => ({ typeUrl: "/metachain.nft.MsgEditNFT", value: data }),
        msgIssueDenom: (data) => ({ typeUrl: "/metachain.nft.MsgIssueDenom", value: data }),
        msgBurnNFT: (data) => ({ typeUrl: "/metachain.nft.MsgBurnNFT", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
