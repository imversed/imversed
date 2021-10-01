// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgMintNFT: (data: MsgMintNFT): EncodeObject => ({ typeUrl: "/metachain.nft.MsgMintNFT", value: data }),
    msgTransferDenom: (data: MsgTransferDenom): EncodeObject => ({ typeUrl: "/metachain.nft.MsgTransferDenom", value: data }),
    msgTransferNFT: (data: MsgTransferNFT): EncodeObject => ({ typeUrl: "/metachain.nft.MsgTransferNFT", value: data }),
    msgEditNFT: (data: MsgEditNFT): EncodeObject => ({ typeUrl: "/metachain.nft.MsgEditNFT", value: data }),
    msgIssueDenom: (data: MsgIssueDenom): EncodeObject => ({ typeUrl: "/metachain.nft.MsgIssueDenom", value: data }),
    msgBurnNFT: (data: MsgBurnNFT): EncodeObject => ({ typeUrl: "/metachain.nft.MsgBurnNFT", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
