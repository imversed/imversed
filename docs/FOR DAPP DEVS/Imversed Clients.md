# Imversed Clients

The Imversed supports different clients in order to support Cosmos and Ethereum transactions and queries:

| | Description | Default Port |
| --- | --- | ---|
| **Cosmos** [**gRPC**](https://docs.imversed.com/developers/clients.html#cosmos-grpc) | Query or send Imversed transactions using gRPC | 9090 |
| **Cosmos REST (**[**gRPC-Gateway**](https://docs.imversed.com/developers/clients.html#cosmos-grpc-gateway)**)** | Query or send Imversed transactions using an HTTP RESTful API | 9091 |
| **Ethereum** [**JSON-RPC**](https://docs.imversed.com/developers/clients.html#ethereum-json-rpc) | Query Ethereum-formatted transactions and blocks or send Ethereum txs using JSON-RPC | 8545
| **Ethereum** [Websocket](https://docs.imversed.com/developers/clients.html#ethereum-websocket) | Subscribe to Ethereum logs and events emitted in smart contracts. | 8586 |
| **Tendermint** [**RPC**](https://docs.imversed.com/developers/clients.html#tendermint-rpc) | Subscribe to Ethereum logs and events emitted in smart contracts. | 26657 |
| **Tendermint** [**Websocket**](https://docs.imversed.com/developers/clients.html#tendermint-websocket) | Query transactions, blocks, consensus state, broadcast transactions, etc. | 26657 |
| **Command Line Interface (**[**CLI**](https://docs.imversed.com/developers/clients.html#cli)**)** | Query or send Imversed transactions using your Terminal or Console. | N/A |

## Cosmos gRPC
Imversed exposes gRPC endpoints for all the integrated Cosmos SDK modules. This makes it easier for wallets and block explorers to interact with the Proof-of-Stake logic and native Cosmos transactions and queries.

### Cosmos gRPC-Gateway (HTTP REST)
[gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) reads a gRPC service definition and generates a reverse-proxy server which translates RESTful JSON API into gRPC. With gRPC-Gateway, users can use REST to interact with the Cosmos gRPC service.

See the list of supported gRPC-Gateway API endpoints for the Imversed testnet [here](https://docs.imversed.com/modules/erc20/08_clients.html#queries-2).

## Ethereum JSON-RPC
Imversed supports most of the standard [JSON-RPC APIs](https://docs.imversed.com/developers/json-rpc/server.html) to connect with existing Ethereum-compatible web3 tooling.

> Check out the list of supported JSON-RPC API [endpoints](https://docs.imversed.com/developers/json-rpc/endpoints.html) and [namespaces](https://docs.imversed.com/developers/json-rpc/namespaces.html).

## Ethereum Websocket
Then, start a websocket subscription with [ws](https://github.com/hashrocket/ws).

```text
# connect to tendermint websocket at port 8546 as defined above
ws ws://localhost:8546/

# subscribe to new Ethereum-formatted block Headers
> {"id": 1, "method": "eth_subscribe", "params": ["newHeads", {}]}
< {"jsonrpc":"2.0","result":"0x44e010cb2c3161e9c02207ff172166ef","id":1}
```

## Tendermint Websocket
Tendermint Core provides a Websocket connection to subscribe or unsubscribe to Tendermint ABCI events.

For more info about how to subscribe to events, please refer to the official [Tendermint documentation](https://docs.tendermint.com/v0.34/tendermint-core/subscription.html).

```text
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": "0",
    "params": {
        "query": "tm.event='<event_value>' AND eventType.eventAttribute='<attribute_value>'"
    }
}
```

### List of Tendermint Events
The main events you can subscribe to are:

* `NewBlock`: Contains `events` triggered during BeginBlock and `EndBlock`.
* `Tx`: Contains events triggered during `DeliverTx` (i.e. transaction processing).
* `ValidatorSetUpdates`: Contains validator set updates for the block.

> 👉 The list of events types and values for each Cosmos SDK module can be found in the Modules Specification section. Check the Events page to obtain the event list of each supported module on Imversed.