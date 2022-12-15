---
description: Intoruction to Users.
helpfulVotes: false
---


# Technical Concepts

## Architecture

```List
SYNOPSIS
Learn how Imversed's architecture leverages the Cosmos SDK Proof-of-Stake functionality, 
    EVM compatibility and fast-finality from Tendermint Core's BFT consensus.
```

:::tip
This documentation page is currently under work in progress.
:::

## Cosmos SDK

Imversed enables the full composability and modularity of the [Cosmos SDK](https://docs.cosmos.network/).


## Tendermint Core & the Application Blockchain Interface (ABCI)

Tendermint consists of two chief technical components: a blockchain consensus engine and a generic application interface. The consensus engine, called [Tendermint Core](https://docs.tendermint.com/), ensures that the same transactions are recorded on every machine in the same order. The application interface, called the [Application Blockchain Interface (ABCI)](https://github.com/tendermint/tendermint/blob/v0.34.x/spec/abci/README.md), enables the transactions to be processed in any programming language.

Tendermint has evolved to be a general purpose blockchain consensus engine that can host arbitrary application states. Since Tendermint can replicate arbitrary applications, it can be used as a plug-and-play replacement for the consensus engines of other blockchains. Imversed is such an example of an ABCI application replacing Ethereum's PoW via Tendermint's consensus engine.

Another example of a cryptocurrency application built on Tendermint is the Cosmos network. Tendermint is able to decompose the blockchain design by offering a very simple API (ie. the ABCI) between the application process and consensus process.

## EVM module

Imversed enables EVM compatibility by implementing various components that together support all the EVM state transitions while ensuring the same developer experience as Ethereum:

- Ethereum transaction format as a Cosmos SDK `Tx` and `Msg` interface

- Ethereum's `secp256k1` curve for the Cosmos Keyring

- ` StateDB` interface for state updates and queries

- JSON-RPC client for interacting with the EVM

# Accounts

```List
SYNOPSIS
This document describes the in-built accounts system of Imversed.
```

## Imversed Accounts

Imversed defines its own custom `Account` type that uses Ethereum's ECDSA secp256k1 curve for keys. This satisfies the [EIP84](https://github.com/ethereum/EIPs/issues/84) for full [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) paths. The root HD path for Imversed-based accounts is `m/44'/60'/0'/0`.

```List
// EthAccount implements the authtypes.AccountI interface and embeds an
// authtypes.BaseAccount type. It is compatible with the auth AccountKeeper.
type EthAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	CodeHash           string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
}
```

## Addresses and Public Keys

[BIP-0173](https://github.com/satoshilabs/slips/blob/master/slip-0173.md) defines a new format for segregated witness output addresses that contains a human-readable part that identifies the Bech32 usage. Imversed uses the following HRP (human readable prefix) as the base HRP:


| Network      | Mainnet        | Testnet
| ------------ | ---------------|-------------
| Imversed     | `Imversed`     | `Imversed`


There are 3 main types of HRP for the `Addresses`/`PubKeys` available by default on Imversed:

  - Addresses and Keys for **accounts**, which identify users (e.g. the sender of a message). They are derived using the **`eth_secp256k1 curve`**.

  - Addresses and Keys for **validator operators**, which identify the operators of validators. They are derived using the **`eth_secp256k1`** curve.

  - Addresses and Keys for **consensus nodes**, which identify the validator nodes participating in consensus. They are derived using the **`ed25519`** curve.


|           | Address bech32 Prefix | Pubkey bech32 Prefix | Curve | Address byte length | Address byte length |
|---------- |-----------------------|----------------------|-------|---------------------|---------------------|
|Accounts   |`Imversed`             | `Imversedhub`        |`eth_secp256k1` | `20`       |`33` (compressed)    |
|Validator Operator| `Imversedvaloper`|`Imversedvaloperpub`| `eth_secp256k1`| `20`       |`33` (compressed)    |
|Consensus Nodes| `Imversedvalcons`   |`Imversedvalconspub`| `ed25519`      | `20`       |`32`                 |


## Address formats for clients

`EthAccount` can be represented in both [Bech32](https://en.bitcoin.it/wiki/Bech32) (`imveresed1`...) and hex (`0x`...) formats for Ethereum's Web3 tooling compatibility.

The Bech32 format is the default format for Cosmos-SDK queries and transactions through CLI and REST clients. The hex format on the other hand, is the Ethereum `common.Address` representation of a Cosmos `sdk.AccAddress`.

- **Address (Bech32)**: `coming soon`

- **Address ([EIP55](https://eips.ethereum.org/EIPS/eip-55)Hex**): 0x91defC7fE5603DFA8CC9B655cF5772459BF10c6f

- **Compressed Public Key**: {"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AsV5oddeB+hkByIJo/4lZiVUgXTzNfBPKC73cZ4K1YD2"}

## Address conversion

The `imversed debug addr <address>` can be used to convert an address between hex and bech32 formats. For example:

:::: tabs ::: tab Bech32

```List
command will be coming soon
```

::: ::: tab Hex

```List
command will be coming soon
```

::: :::

## Key output

:::tip
The Cosmos SDK Keyring output (i.e `imversed keys`) only supports addresses and public keys in Bech32 format.
:::

We can use the `keys show` command of `imversed` with the flag `--bech <type> (acc|val|cons)` to obtain the addresses and keys as mentioned above,

:::: tabs ::: tab Account

```List
command will be coming soon
```

::: ::: tab Validator

```List
command will be coming soon
```

::: ::: tab Consensus

```List
command will be coming soon
```

::: ::::

## Querying an Account

You can query an account address using the CLI, gRPC or

### Command Line Interface

```List
# NOTE: the --output (-o) flag will define the output format in JSON or YAML (text)
imversed q auth account $(imversed keys show mykey -a) -o text
|
  '@type': /ethermint.types.v1.EthAccount
  base_account:
    account_number: "0"
    address: imv1z3t55m0l9h0eupuz3dp5t5cypyv674jj7mz2jw
    pub_key:
      '@type': /ethermint.crypto.v1.ethsecp256k1.PubKey
      key: AsV5oddeB+hkByIJo/4lZiVUgXTzNfBPKC73cZ4K1YD2
    sequence: "1"
  code_hash: 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470

```
### Cosmos gRPC and REST

```List
# GET /cosmos/auth/v1beta1/accounts/{address}
curl -X GET "http://localhost:10337/cosmos/auth/v1beta1/accounts/imv14au322k9munkmx5wrchz9q30juf5wjgz2cfqku" -H "accept: application/json"
```

### SON-RPC

To retrieve the Ethereum hex address using Web3, use the JSON-RPC `eth_accounts` or `personal_listAccounts` endpoints:
```List
# query against a local node
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0","method":"personal_listAccounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545
```

# Chain ID

```List
SYNOPSIS
Learn about the Imversed chain-id format
```

## Official Chain IDs

:::tip
NOTE: The latest Chain ID (i.e highest Version Number) is the latest version of the software and mainnet.
:::

:::: tabs ::: tab Mainnet

|Name       | Chain ID  | Identifier | EIP155 Number | Version Number |
|---------- |-----------|------------|---------------|----------------|
|Imversed 2 |           |            |               |                |
|Imversed 1 |           |            |               |                |


::: ::: tab Testnets

|Name       | Chain ID  | Identifier | EIP155 Number | Version Number |
|---------- |-----------|------------|---------------|----------------|
|Imversed Public Testnet |           |            |               |                |
|Imversed Public Testnet |           |            |               |                |
|Olympus Mons Incentivized Testnet |           |            |               |                |
|Arsia Mons Testnet |           |            |               |                |

::: ::::

:::tip
You can also lookup the [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `Chain ID` by referring to [chainlist.org](https://chainlist.org/).
:::

```List 
Coming soon
```

## The Chain Identifier

Every chain must have a unique identifier or `chain-id`. Tendermint requires each application to define its own `chain-id` in the [genesis.json fields](https://github.com/tendermint/tendermint/blob/v0.34.x/spec/core/genesis.md). However, in order to comply with both EIP155 and Cosmos standard for chain upgrades, Imversed-compatible chains must implement a special structure for their chain identifiers.

## Structure

The Imversed Chain ID contains 3 main components

- **Identifier**: Unstructured string that defines the name of the application.

- **EIP155 Number**: Immutable [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `CHAIN_ID` that defines the replay attack protection number.

- **Version Number**: Is the version number (always positive) that the chain is currently running. This number **MUST** be incremented every time the chain is upgraded or forked in order to avoid network or consensus errors.


## Format

The format for specifying and Imversed compatible chain-id in genesis is the following:

```List
{identifier}_{EIP155}-{version}
```

The following table provides an example where the second row corresponds to an upgrade from the first one:

| Chain ID  | Identifier | EIP155 Number | Version Number |
|---------- |-----------|----------------|----------------|
|           |           |                |                |
|           |           |                |                |
|           |           |                |                |
|           |           |                |                |


## Encoding

```List
SYNOPSIS
Learn about the encoding formats used on Imversed.
```

# Encoding Formats

## Protocol Buffers

The Cosmos [Stargate](https://stargate.cosmos.network/) release introduces [protobuf](https://developers.google.com/protocol-buffers)as the main encoding format for both client and state serialization. All the EVM module types that are used for state and clients (transaction messages, genesis, query services, etc) will be implemented as protocol buffer messages.

## RLP

Recursive Length Prefix ([RLP](https://ethereum.org/en/)), is an encoding/decoding algorithm that serializes a message and allows for quick reconstruction of encoded data. Imversed uses RLP to encode/decode Ethereum messages for JSON-RPC handling to conform messages to the proper Ethereum format. This allows messages to be encoded and decoded in the exact format as Ethereum's.

The `x/evm` transactions (`MsgEthereumTx`) encoding is performed by casting the message to a go-ethereum's `Transaction` and then marshaling the transaction data using RLP:

```List
// TxEncoder overwrites sdk.TxEncoder to support MsgEthereumTx
func (g txConfig) TxEncoder() sdk.TxEncoder {
  return func(tx sdk.Tx) ([]byte, error) {
    msg, ok := tx.(*evmtypes.MsgEthereumTx)
    if ok {
      return msg.AsTransaction().MarshalBinary()
   }
    return g.TxConfig.TxEncoder()(tx)
  }
}

// TxDecoder overwrites sdk.TxDecoder to support MsgEthereumTx
func (g txConfig) TxDecoder() sdk.TxDecoder {
  return func(txBytes []byte) (sdk.Tx, error) {
    tx := &ethtypes.Transaction{}

    err := tx.UnmarshalBinary(txBytes)
    if err == nil {
      msg := &evmtypes.MsgEthereumTx{}
      msg.FromEthereumTx(tx)
      return msg, nil
    }

    return g.TxConfig.TxDecoder()(txBytes)
  }
}
```

# Pending State

```List
SYNOPSIS
Learn how Imversed handles pending state queries.
```
## Imversed vs Ethereum

In Ethereum, pending blocks are generated as they are queued for production by miners. These pending blocks include pending transactions that are picked out by miners, based on the highest reward paid in gas. This mechanism exists as block finality is not possible on the Ethereum network. Blocks are committed with probabilistic finality, which means that transactions and blocks become less likely to become reverted as more time (and blocks) passes.

Imversed is designed quite differently on this front as there is no concept of a "pending state". Imversed uses [Tendermint Core](https://docs.tendermint.com/) BFT consensus which provides instant finality for transaction. For this reason, Ethermint does not require a pending state mechanism, as all (if not most) of the transactions will be committed to the next block (avg. block time on Cosmos chains is ~8s). However, this causes a few hiccups in terms of the Ethereum Web3-compatible queries that can be made to pending state.

Another significant difference with Ethereum, is that blocks are produced by validators or block producers, who include transactions from their local mempool into blocks in a first-in-first-out (FIFO) fashion. Transactions on Imversed cannot be ordered or cherry picked out from the Tendermint node [mempool](https://docs.tendermint.com/master/tendermint-core/mempool/).

## Pending State Queries

Imversed will make queries which will account for any unconfirmed transactions present in a node's transaction mempool. A pending state query made will be subjective and the query will be made on the target node's mempool. Thus, the pending state will not be the same for the same query to two different nodes.

## JSON-RPC Calls on Pending Transactions

- `eth_getBalance`

- `eth_getTransactionCount`

- `eth_getBlockTransactionCountByNumber`

- `eth_getBlockByNumber`

- `eth_getTransactionByHash`

- `eth_getTransactionByBlockNumberAndIndex`

- `eth_sendTransaction`
