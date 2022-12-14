---
description: Intoruction to Developers.
helpfulVotes: false
---

# Imversed Clients

```List
SYNOPSIS
Learn about all the available services for clients
```

The Imversed supports different clients in order to support Cosmos and Ethereum transactions and queries:

|                            | Description | Default Port |
|----------------------------|-------------|------------|
|Cosmos gRPC                 |Query or send Imversed transactions using gRPC                                      |N/A |                
|Cosmos REST (gRPC-Gateway)  |Query or send Imversed transactions using an HTTP RESTful API                       |N/A |                
|Ethereum JSON-RPC           |Query Ethereum-formatted transactions and blocks or send Ethereum txs using JSON-RPC|N/A | 
|Ethereum Websocket          |Subscribe to Ethereum logs and events emitted in smart contracts.                   |N/A |                
|Tendermint RPC              |Subscribe to Ethereum logs and events emitted in smart contracts.                   |N/A |   
|Tendermint Websocket        |Query transactions, blocks, consensus state, broadcast transactions, etc.           |N/A |                
|Command Line Interface (CLI)|Query or send Imversed transactions using your Terminal or Console.                    |N/A |

## Cosmos gRPC

Imversed exposes gRPC endpoints for all the integrated Cosmos SDK modules. This makes it easier for wallets and block explorers to interact with the Proof-of-Stake logic and native Cosmos transactions and queries.

## Cosmos gRPC-Gateway (HTTP REST)

[gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) reads a gRPC service definition and generates a reverse-proxy server which translates RESTful JSON API into gRPC. With gRPC-Gateway, users can use REST to interact the Cosmos gRPC service.

See the list of supported gRPC-Gateway API endpoints for the Imversed testnet.

## Ethereum JSON-RPC

Imversed supports most of the standard JSON-RPC APIs to connect with existing Ethereum-compatible web3 tooling.

:::tip
Check out the list of supported JSON-RPC API endpoints and namespaces.
:::

## Ethereum Websocket

Then, start a websocket subscription with [ws](https://github.com/hashrocket/ws)

# connect to tendermint websocet at port 8546 as defined above
ws ws://localhost:8546/

```List
# subscribe to new Ethereum-formatted block Headers
> {"id": 1, "method": "eth_subscribe", "params": ["newHeads", {}]}
< {"jsonrpc":"2.0","result":"0x44e010cb2c3161e9c02207ff172166ef","id":1}
```

## Tendermint Websocket

Tendermint Core provides a Websocket connection to subscribe or unsubscribe to Tendermint ABCI events.

For more info about the how to subscribe to events, please refer to the official [Tendermint documentation](https://docs.tendermint.com/v0.34/tendermint-core/subscription.html).


### List of Tendermint Events

The main events you can subscribe to are:

- `NewBlock`: Contains `events` triggered during `BeginBlock` and `EndBlock`.

- `Tx`: Contains `events` triggered during `DeliverTx` (i.e. transaction processing).

- `ValidatorSetUpdates`: Contains validator set updates for the block.

:::tip
👉 The list of events types and values for each Cosmos SDK module can be found in the Modules Specification section. Check the Events page to obtain the event list of each supported module on Imversed.
:::

List of all Tendermint event keys:

|                            | Event Type | Categories|
|----------------------------|------------|-----------|
|Subscribe to a specific event                        |`"tm.event"`    |`block`|                
|Subscribe to a specific transaction                  |`"tx.hash"`     |`block`|                
|Subscribe to transactions at a specific block height |`"tx.height"`   |`block`| 
|Index `BeginBlock` and `Endblock` events             |`"block.height"`|`block`|                
|Subscribe to ABCI `BeginBlock` events                |`"begin_block"` |`block`|   
|Subscribe to ABCI `EndBlock` events                  |`"end_block"`   |`block`|

Below is a list of values that you can use to subscribe for the tm.event type:

|                            | Event Type | Categories |
|----------------------------|-------------|------------|
|New block                       |`"NewBlock"`      |`block`|                
|New block header                |`"NewBlockHeader"`|`block`|                
|New Byzantine Evidence          |`"NewEvidence"`   |`block`| 
|New transaction                 |`"block.height"`  |`block`|                
|Validator set updated           |`"begin_block"`   |`block`|   
|Block sync status               |`"end_block"`     |`block`|
|lock                            |`"tm.event"`      |`block`|                
|New consensus round             |`"tx.hash"`       |`block`|                
|Polka                           |`"tx.height"`     |`block`| 
|Relock                          |`"block.height"`  |`block`|                
|State sync status               |`"begin_block"`   |`block`|   
|Timeout propose                 |`"end_block"`     |`block`|

### Example

```List
ws ws://localhost:26657/websocket
> { "jsonrpc": "2.0", "method": "subscribe", "params": ["tm.event='ValidatorSetUpdates'"], "id": 1 }
```

**Example response**:
```List
{
    "jsonrpc": "2.0",
    "id": 0,
    "result": {
        "query": "tm.event='ValidatorSetUpdates'",
        "data": {
            "type": "tendermint/event/ValidatorSetUpdates",
            "value": {
              "validator_updates": [
                {
                  "address": "09EAD022FD25DE3A02E64B0FE9610B1417183EE4",
                  "pub_key": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "ww0z4WaZ0Xg+YI10w43wTWbBmM3dpVza4mmSQYsd0ck="
                  },
                  "voting_power": "10",
                  "proposer_priority": "0"
                }
              ]
            }
        }
    }
}
```

## CLI

Users can use the Imversed binary to interact directly with an Imversed node though the CLI.

:::tip
👉 To use the CLI, you will need to provide a Tendermint RPC address for the --node flag. Look for a publicly available addresses for testnet and mainnet in the Quick Connect page.
:::

- **Transactions**: `imversed tx`

The list of available commands, as of `v3.0.0`, are:

```List
Available Commands:
  authz               Authorization transactions subcommands
  bank                Bank transaction subcommands
  broadcast           Broadcast transactions generated offline
  crisis              Crisis transactions subcommands
  decode              Decode a binary encoded transaction string
  distribution        Distribution transactions subcommands
  encode              Encode transactions generated offline
  erc20               erc20 subcommands
  evidence            Evidence transaction subcommands
  evm                 evm transactions subcommands
  feegrant            Feegrant transactions subcommands
  gov                 Governance transactions subcommands
  ibc                 IBC transaction subcommands
  ibc-transfer        IBC fungible token transfer transaction subcommands
  multisign           Generate multisig signatures for transactions generated offline
  multisign-batch     Assemble multisig transactions in batch from batch signatures
  sign                Sign a transaction generated offline
  sign-batch          Sign transaction batch files
  slashing            Slashing transaction subcommands
  staking             Staking transaction subcommands
  validate-signatures validate transactions signatures
  vesting             Vesting transaction subcommands
```

- **Queries**: `imversed` query

The list of available commands, as of `v3.0.0`, are:

```List
Available Commands:
  account                  Query for account by address
  auth                     Querying commands for the auth module
  authz                    Querying commands for the authz module
  bank                     Querying commands for the bank module
  block                    Get verified data for a the block at given height
  claims                   Querying commands for the claims module
  distribution             Querying commands for the distribution module
  epochs                   Querying commands for the epochs module
  erc20                    Querying commands for the erc20 module
  evidence                 Query for evidence by hash or for all (paginated) submitted evidence
  evm                      Querying commands for the evm module
  feegrant                 Querying commands for the feegrant module
  feemarket                Querying commands for the fee market module
  gov                      Querying commands for the governance module
  ibc                      Querying commands for the IBC module
  ibc-transfer             IBC fungible token transfer query subcommands
  incentives               Querying commands for the incentives module
  inflation                Querying commands for the inflation module
  params                   Querying commands for the params module
  recovery                 Querying commands for the recovery module
  slashing                 Querying commands for the slashing module
  staking                  Querying commands for the staking module
  tendermint-validator-set Get the full tendermint validator set at given height
  tx                       Query for a transaction by hash, "<addr>/<seq>" combination or comma-separated signatures in a committed block
  txs                      Query for paginated transactions that match a set of events
  upgrade                  Querying commands for the upgrade module
  vesting                  Querying commands for the vesting module
```

:::tip
**Note**: When querying Ethereum transactions versus Cosmos transactions, the transaction hashes are different. When querying Ethereum transactions, users need to use event query. Here's an example with the CLI:
:::

```List
command will be coming soon
```
