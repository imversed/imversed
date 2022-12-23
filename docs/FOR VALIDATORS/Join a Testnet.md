# Join a Testnet

## Pick a Testnet
You specify the network you want to join by setting the genesis file and seeds. If you need more information about past networks, check our [testnets repo](https://github.com/imversed/testnets).

| Testnet Chain ID | Description | Site | Version | Status |
| --- | --- | --- | --- | --- |
| `imversed_5555558-1` | imversed_5555558-1 Testnet | imversed_5555558-1 | v3.11 | `Live` |


## Install `imversed`
Follow the [installation](https://docs.imversed.com/validators/quickstart/installation.html) document to install the Imversed binary `imversed`.

> Make sure you have the right version of imversed installed.

### Save Chain ID
We recommend saving the testnet `chain-id` into your `imversed`'s `client.toml`. This will make it so you do not have to manually pass in the chain-id flag for every CLI command.

> See the Official [Chain IDs](https://docs.imversed.com/users/technical_concepts/chain_id.html#official-chain-ids) for reference

```text
imversed config chain-id imversed_5555558-1
```

## Initialize Node

We need to initialize the node to create all the necessary validator and node configuration files:

```text
imversed init <your_custom_moniker> --chain-id imversed_5555558-1
```

> Monikers can contain only ASCII characters. Using Unicode characters will render your node unreachable.

By default, the `init` command creates your `~/.imversed` (i.e `$HOME`) directory with subfolders `config/` and `data/`. In the `config` directory, the most important files for configuration are `app.toml` and `config.toml`.

## Genesis & Seeds

### Copy the Genesis File

Check the `genesis.json` file from the [`archive`](https://archive.imversed.dev/imversed_5555558-1/genesis.json) and copy it over to the `config` directory: `~/.imversed/config/genesis.json`. This is a genesis file with the chain-id and genesis accounts balances.

```text
sudo apt install -y unzip wget
wget -P ~/.imversed/config https://storage.googleapis.com/static.fdvr.co/imversed/testnet/genesis.json
```
Then verify the correctness of the genesis configuration file:

```text
imversed validate-genesis
```

### Add Seed Nodes

Your node needs to know how to find peers. You'll need to add healthy [seed nodes](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#seed) to `$HOME/.imversed/config/config.toml`. The [`testnets`](https://github.com/imversed/testnets) repo contains links to some seed nodes.

Edit the file located in `~/.imversed/config/config.toml` and the `seeds` to the following:

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
SEEDS=`curl -sL https://raw.githubusercontent.com/tharsis/testnets/main/imversed_5555558-1/seeds.txt | awk '{print $1}' | paste -s -d, -`
sed -i.bak -e "s/^seeds =.*/seeds = \"$SEEDS\"/" ~/.imversed/config/config.toml
```

> For more information on seeds and peers, you can the Tendermint [P2P documentation](https://docs.tendermint.com/master/spec/p2p/peer.html).

## Add Persistent Peers
We can set the [`persistent_peers`](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#persistent-peer) field in `~/.imversed/config/config.toml` to specify peers that your node will maintain persistent connections with. You can retrieve them from the list of available peers on the [`testnets`](https://github.com/imversed/testnets) repo.

A list of available persistent peers is also available in the `#find-peers` channel in the [Imversed Discord](https://discord.gg/imversed). You can get a random 10 entries from the `peers.txt` file in the PEERS variable by running the following command:

```text
PEERS=`curl -sL https://raw.githubusercontent.com/tharsis/testnets/main/imversed_5555558-1/peers.txt | sort -R | head -n 10 | awk '{print $1}' | paste -s -d, -`
```
Use `sed` to include them into the configuration. You can also add them manually:

```text
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" ~/.imversed/config/config.toml
```

## Run a Testnet Validator
Claim your testnet tImversed on the [faucet](https://docs.imversed.com/developers/testnet/faucet.html) using your validator account address and submit your validator account address:

> For more details on how to run your validator, follow `these` instructions.

```text
imversed tx staking create-validator \
  --amount=1000000000000aimv \
  --pubkey=$(imversed tendermint show-validator) \
  --moniker="ImversedWhale" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="0.025aimv" \
  --from=<key_name>
```

## Start testnet
The final step is to [start the nodes](https://docs.imversed.com/validators/quickstart/run_node.html#start-node). Once enough voting power (+2/3) from the genesis validators is up-and-running, the testnet will start producing blocks.

```text
imversed start
```

## Upgrading Your Node

> These instructions are for full nodes that have ran on previous versions of and would like to upgrade to the latest testnet version.

### Reset Data

> If the version <new_version> you are upgrading to is not breaking from the previous one, you **should not** reset the data. If this is the case you can skip to [Restart](https://docs.imversed.com/validators/testnet.html#restart)

First, remove the outdated files and reset the data.

```text
rm $HOME/.imversed/config/addrbook.json $HOME/.imversed/config/genesis.json
imversed tendermint unsafe-reset-all --home $HOME/.imversed
```
Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.

> Warning
> Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to double sign.

### Restart
To restart your node, just type:

```text
imversed start
```
## Share your Peer

You can share your peer to posting it in the `#find-peers` channel in the [Imversed Discord](https://discord.gg/imversed).

> To get your Node ID use
> ```text
> imversed tendermint show-node-id
> ```

### State Syncing a Node
If you want to join the network using State Sync (quick, but not applicable for archive nodes), check our [State Sync](https://docs.imversed.com/validators/setup/statesync.html) page


