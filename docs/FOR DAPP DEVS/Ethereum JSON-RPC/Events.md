# Events

## Subscribing to Events

### Cosmos and Tendermint Events

> Check the [Tendermint Websocket](https://docs.imversed.com/developers/clients.html#tendermint-websocket) in the Clients documentation for reference.

It is possible to subscribe to `Events` via Tendermint's [Websocket](https://docs.tendermint.com/v0.34/tendermint-core/subscription.html). This is done by calling the `subscribe` RPC method via Websocket:

```json
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": "0",
    "params": {
        "query": "tm.event='<event_value>' AND eventType.eventAttribute='<attribute_value>'"
    }
}
```

These events are triggered after a block is committed. You can get the full list of `event` categories and values [here](https://docs.imversed.com/developers/clients.html#list-of-tendermint-events).

The `type` and `attribute` value of the `query` allow you to filter the specific `event` you are looking for. For example, a an Ethereum transaction on Imversed (`MsgEthereumTx`) triggers an `event` of type `ethermint` and has `sender` and `recipient` as `attributes`. Subscribing to this `event` would be done like so:

```text
{
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": "0",
    "params": {
        "query": "tm.event='Tx' AND ethereum.recipient='hexAddress'"
    }
}
```

where `hexAddress` is an Ethereum hex address (eg: `0x1122334455667788990011223344556677889900`).

**Ethereum Events**

Imversed also supports the Ethereum [JSON-RPC](https://docs.imversed.com/developers/json-rpc/server.html) filters calls to subscribe to [state logs](https://eth.wiki/json-rpc/API#eth_newfilter), [blocks](https://eth.wiki/json-rpc/API#eth_newblockfilter) or [pending transactions](https://eth.wiki/json-rpc/API#eth_newpendingtransactionfilter) changes.

Under the hood, it uses the Tendermint RPC client's event system to process subscriptions that are then formatted to Ethereum-compatible events.

```text
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newBlockFilter","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

{"jsonrpc":"2.0","id":1,"result":"0x3503de5f0c766c68f78a03a3b05036a5"}
```

Then you can check if the state changes with the [`eth_getFilterChanges`](https://eth.wiki/json-rpc/API#eth_getfilterchanges) call:

```text
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getFilterChanges","params":["0x3503de5f0c766c68f78a03a3b05036a5"],"id":1}' -H "Content-Type: application/json" http://localhost:8545

{"jsonrpc":"2.0","id":1,"result":["0x7d44dceff05d5963b5bc81df7e9f79b27e777b0a03a6feca09f3447b99c6fa71","0x3961e4050c27ce0145d375255b3cb829a5b4e795ac475c05a219b3733723d376","0xd7a497f95167d63e6feca70f344d9f6e843d097b62729b8f43bdcd5febf142ab","0x55d80a4ba6ef54f2a8c0b99589d017b810ed13a1fda6a111e1b87725bc8ceb0e","0x9e8b92c17280dd05f2562af6eea3285181c562ebf41fc758527d4c30364bcbc4","0x7353a4b9d6b35c9eafeccaf9722dd293c46ae2ffd4093b2367165c3620a0c7c9","0x026d91bda61c8789c59632c349b38fd7e7557e6b598b94879654a644cfa75f30","0x73e3245d4ddc3bba48fa67633f9993c6e11728a36401fa1206437f8be94ef1d3"]}
```

## Websocket Connection

### Tendermint Websocket

To start a connection with the Tendermint websocket you need to define the address with the `--rpc.laddr` flag when starting the node (default `tcp://127.0.0.1:26657`):

```text
imversed start --rpc.laddr="tcp://127.0.0.1:26657"
```
Then, start a websocket subscription with [ws](https://github.com/hashrocket/ws).

```text
# connect to tendermint websocket at port 8080 as defined above
ws ws://localhost:8080/websocket

# subscribe to new Tendermint block headers
> { "jsonrpc": "2.0", "method": "subscribe", "params": ["tm.event='NewBlockHeader'"], "id": 1 }
```

### Ethereum Websocket
Since Imversed runs uses Tendermint Core as it's consensus Engine and it's built with the Cosmos SDK framework, it inherits the event format from them. However, in order to support the native Web3 compatibility for websockets of the [Ethereum's PubSubAPI](https://geth.ethereum.org/docs/rpc/pubsub), Imversed needs to cast the Tendermint responses retrieved into the Ethereum types.

You can start a connection with the Ethereum websocket using the `--json-rpc.ws-address` flag when starting the node (default `"0.0.0.0:8546"`):

```text
imversed start  --json-rpc.address"0.0.0.0:8545" --json-rpc.ws-address="0.0.0.0:8546" --evm.rpc.api="eth,web3,net,txpool,debug" --json-rpc.enable
```
Then, start a websocket subscription with [ws](https://github.com/hashrocket/ws).

```text
# connect to tendermint websocet at port 8546 as defined above
ws ws://localhost:8546/

# subscribe to new Ethereum-formatted block Headers
> {"id": 1, "method": "eth_subscribe", "params": ["newHeads", {}]}
< {"jsonrpc":"2.0","result":"0x44e010cb2c3161e9c02207ff172166ef","id":1}
```
