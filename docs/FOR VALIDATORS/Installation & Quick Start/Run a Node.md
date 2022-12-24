# Run a Node

## Automated deployment
Run the local node by running the `local_node.sh` script in the base directory of the repository.

> The script below will remove any pre-existing binaries installed. Use the manual deploy if you want to keep your binaries and configuration files.

```text
./local_node.sh
```
## Manual deployment
The instructions for setting up a brand new full node from scratch are the the same as running a [single node local testnet](https://docs.imversed.com/developers/localnet/single_node.html#manual-localnet).

## Start node
To start your node, just type:

```text
imversed start --json-rpc.enable=true --json-rpc.api="eth,web3,net"
```

## Key Management
To run a node with the same key every time: replace imversed keys add $KEY in ./local_node.sh with:

```text
echo "your mnemonic here" | imversed keys add $KEY --recover
```

> Imversed currently only supports 24 word mnemonics.

You can generate a new key/mnemonic with:

```text
imversed keys add $KEY
```

To export your imv key as an Ethereum private key (for use with [Metamask](https://docs.imversed.com/users/wallets/metamask.html) for example):

```text
imversed keys unsafe-export-eth-key $KEY
```

For more about the available key commands, use the `--help` flag

```text
imversed keys -h
```

### Keyring backend options
The instructions above include commands to use `test` as the `keyring-backend`. This is an unsecured keyring that doesn't require entering a password and should not be used in production. Otherwise, Imversed supports using a file or OS keyring backend for key storage. To create and use a file stored key instead of defaulting to the OS keyring, add the flag `--keyring-backend file` to any relevant command and the password prompt will occur through the command line. This can also be saved as a CLI config option with:

```text
imversed config keyring-backend file
```

> For more information about the Keyring and its backend options, click [here](https://docs.imversed.com/users/keys/keyring.html).

## Enable Tracing
To enable tracing when running the node, modify the last line of the `local_node.sh` script to be the following command, where:
* `$TRACER `is the EVM tracer type to collect execution traces from the EVM transaction execution (eg. `json|struct|access_list|markdown`)
* `$TRACESTORE` is the output file which contains KVStore tracing (eg. `store.txt`)

```text
imversed start --evm.tracer $TRACER --tracestore $TRACESTORE --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001aimv --json-rpc.api eth,txpool,personal,net,debug,web3
```

## Clearing data from chain

### Reset Data
Alternatively, you can reset the blockchain database, remove the node's address book files, and reset the p `riv_validator.json` to the genesis state.

> If you are running a **validator node**, always be careful when doing `imversed unsafe-reset-all`. You should never use this command if you are not switching `chain-id`.
> **IMPORTANT:** Make sure that every node has a unique `priv_validator.json`. **Do not** copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to double sign!

First, remove the outdated files and reset the data.

```text
rm $HOME/.imversed/config/addrbook.json $HOME/.imversed/config/genesis.json
imversed tendermint unsafe-reset-all --home $HOME/.imversed
```
Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.

### Delete Data
Data for the imversed binary should be stored at `~/.imversed`, respectively by default. To **delete** the existing binaries and configuration, run:

```text
rm -rf ~/.imversed
```

To clear all data except key storage (if keyring backend chosen) and then you can rerun the full node installation commands from above to start the node again.

## Recording Transactions Per Second (TPS)

In order to get a progressive value of the transactions per second, we use Prometheus to return the values. The Prometheus exporter runs at address http://localhost:8877 so please add this section to your [Prometheus installation](https://opencensus.io/codelabs/prometheus/#1) config.yaml file like this:

```text
global:
  scrape_interval: 10s

  external_labels:
    monitor: 'imv'

scrape_configs:
  - job_name: 'imv'

    scrape_interval: 10s

    static_configs:
      - targets: ['localhost:8877']
```

and then run Prometheus like this

```text
prometheus --config.file=prom_config.yaml
```

and then visit the Prometheus dashboard at http://localhost:9090/ then navigate to the expression area and enter the following expression

```text
rate(imversed_transactions_processed[1m])
```

which will show the rate of transactions processed.