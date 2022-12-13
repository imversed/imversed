# Join Mainnet

## Mainnet
You need to set the **genesis file** and **seeds**. If you need more information about past networks, check our [mainnet repo](https://github.com/imversed/mainnet). The table below gives an overview of all Mainnet Chain IDs. Note that, the displayed version might differ when an active Software Upgrade proposal exists on chain.

| Chain ID | Description | Site | Version | Status |
| --- | --- | --- | --- | --- |
| `imversed_5555555-1` | imversed_5555555-1  | Imversed | v3.11 | `Live` |

> **IMPORTANT:** If you join mainnet as a validator make sure you follow all the [security](https://docs.imversed.com/validators/security/security.html) recommendations!

## Install `imversed`
Follow the [installation](https://docs.imversed.com/validators/quickstart/installation.html) document to install the Imversed binary `imversed`.

> Make sure you have the right version of `imversed` installed.

### Save Chain ID
We recommend saving the mainnet `chain-id` into your `imversed`'s `client.toml`. This will make it so you do not have to manually pass in the `chain-id` flag for every CLI command.

> See the Official [Chain IDs](https://docs.imversed.com/users/technical_concepts/chain_id.html#official-chain-ids) for reference.

```text
imversed config chain-id imversed_9001-2
```

## Initialize Node
We need to initialize the node to create all the necessary validator and node configuration files:

```text
imversed init <your_custom_moniker> --chain-id imversed_9001-2
```

> Monikers can contain only ASCII characters. Using Unicode characters will render your node unreachable.

By default, the `init` command creates your `~/.imversed` (i.e `$HOME`) directory with subfolders `config/` and `data/`. In the config directory, the most important files for configuration are `app.toml` and `config.toml`.

## Genesis & Seeds

### Copy the Genesis File
Download the `genesis.json` file from the [`archive`](https://archive.imversed.org/mainnet/genesis.json) and copy it over to the `config` directory: ~/.imversed/config/genesis.json. This is a genesis file with the chain-id and genesis accounts balances.

```text
wget https://archive.imversed.org/mainnet/genesis.json
mv genesis.json ~/.imversed/config/
```

Then verify the correctness of the genesis configuration file:

```text
imversed validate-genesis
```

### Add Seed Nodes
Your node needs to know how to find [peers](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#peers). You'll need to add healthy [`seed nodes`](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#seed) to `$HOME/.imversed/config/config.toml`. The [`mainnet`](https://github.com/imversed/mainnet) repo contains links to some seed nodes.

Edit the file located in `~/.imversed/config/config.toml` and the seeds to the following:

```text
#######################################################
###           P2P Configuration Options             ###
#######################################################
[p2p]

# ...

# Comma separated list of seed nodes to connect to
seeds = "<node-id>@<ip>:<p2p port>"
```

You can use the following code to get seeds from the repo and add it to your config:

```text
SEEDS=`curl -sL https://raw.githubusercontent.com/tharsis/mainnet/main/imversed_9001-2/seeds.txt | awk '{print $1}' | paste -s -d, -`
sed -i.bak -e "s/^seeds =.*/seeds = \"$SEEDS\"/" ~/.imversed/config/config.toml
```

> For more information on seeds and peers, you can the Tendermint [P2P documentation](https://docs.tendermint.com/master/spec/p2p/peer.html).

### Add Persistent Peers
We can set the [`persistent_peers`](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#persistent-peer) field in `~/.imversed/config/config.toml` to specify peers that your node will maintain persistent connections with. You can retrieve them from the list of available peers on the [`mainnet`](https://github.com/imversed/mainnet) repo.

A list of available persistent peers is also available in the `#find-peers` channel in the [Imversed Discord](https://discord.gg/imversed). You can get a random 10 entries from the `peers.txt` file in the `PEERS` variable by running the following command:

```text
PEERS=`curl -sL https://raw.githubusercontent.com/tharsis/mainnet/main/imversed_9001-2/peers.txt | sort -R | head -n 10 | awk '{print $1}' | paste -s -d, -`
```
Use sed to include them into the configuration. You can also add them manually:

```text
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" ~/.imversed/config/config.toml
```

## Run a Mainnet Validator

> For more details on how to run your validator, follow the validator these instructions.

```text
imversed tx staking create-validator \
  --amount=1000000000000aimv \
  --pubkey=$(imversed tendermint show-validator) \
  --moniker="ImversedWhale" \
  --chain-id=<chain_id> \
  --commission-rate="0.05" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="0.025aimv" \
  --from=<key_name>
```

> 🚨 DANGER: Never create your validator keys using a test keying backend. Doing so might result in a loss of funds by making your funds remotely accessible via the eth_sendTransaction JSON-RPC endpoint.
> 
> Ref: [Security Advisory: Insecurely configured geth can make funds remotely accessible](https://blog.ethereum.org/2015/08/29/security-alert-insecurely-configured-geth-can-make-funds-remotely-accessible/)

### Start mainnet
The final step is to [start the nodes](https://docs.imversed.com/validators/quickstart/run_node.html#start-node). Once enough voting power (+2/3) from the genesis validators is up-and-running, the node will start producing blocks.

```text
imversed start
```

### Share your Peer
You can share your peer to posting it in the #find-peers channel in the [Imversed Discord](https://discord.com/channels/911489720536690688/1019664266804080720).

> To get your Node ID use
> ```text
> imversed tendermint show-node-id
> ```

## State Syncing a Node
If you want to join the network using State Sync (quick, but not applicable for archive nodes), check our [State Sync](https://docs.imversed.com/validators/setup/statesync.html) page.

