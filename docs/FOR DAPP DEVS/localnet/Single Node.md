# Single Node
## Automated Localnet (script)
You can customize the local testnet script by changing values for convenience for example:

```bash
# customize the name of your key, the chain-id, moniker of the node, keyring backend, and log level
KEY="dev0"
CHAINID="5555555"
MONIKER="localtestnet"
KEYRING="test"
LOGLEVEL="info"


# Allocate genesis accounts (cosmos formatted addresses)
imversed add-genesis-account $KEY 100000000000000000000000000aimv --keyring-backend $KEYRING

# Sign genesis transaction
imversed gentx $KEY 1000000000000000000000aimv --keyring-backend $KEYRING --chain-id $CHAINID
```

The default configuration will generate a single validator localnet with the chain-id `imversed-1` and one predefined account (`dev0`) with some allocated funds at the genesis.

You can start the local chain using:

```bash
local_node.sh
```

## Manual Localnet
This guide helps you create a single validator node that runs a network locally for testing and other development related uses.

### Initialize the chain
Before actually running the node, we need to initialize the chain, and most importantly its genesis file. This is done with the `init` subcommand:

```bash
$MONIKER=testing
$KEY=dev0
$CHAINID="5555555"

# The argument $MONIKER is the custom username of your node, it should be human-readable.
imversed init $MONIKER --chain-id=$CHAINID
```

> You can [edit](https://docs.imversed.com/validators/quickstart/binary.html#configuring-the-node) this `moniker` later by updating the `config.toml` file.

The command above creates all the configuration files needed for your node and validator to run, as well as a default genesis file, which defines the initial state of the network. All these [configuration files](https://docs.imversed.com/validators/quickstart/binary.html#configuring-the-node) are in `~/.imversed` by default, but you can overwrite the location of this folder by passing the `--home` flag.

## Genesis Procedure

### Adding Genesis Accounts
Before starting the chain, you need to populate the state with at least one account using the keyring:

```linux
imversed keys add my_validator
```

Once you have created a local account, go ahead and grant it some aimv tokens in your chain's genesis file. Doing so will also make sure your chain is aware of this account's existence:

```linux
imversed add-genesis-account my_validator 10000000000aimv
```

Now that your account has some tokens, you need to add a validator to your chain.

For this guide, you will add your local node (created via the `init` command above) as a validator of your chain. Validators can be declared before a chain is first started via a special transaction included in the genesis file called a `gentx`:

```linux
# Create a gentx
# NOTE: this command lets you set the number of coins. 
# Make sure this account has some coins with the genesis.app_state.staking.params.bond_denom denom
imversed add-genesis-account my_validator 1000000000stake,10000000000aimv
```

A gentx does three things:

1. Registers the `validator` account you created as a validator operator account (i.e. the account that
   controls the validator).
2. Self-delegates the provided `amount` of staking tokens.
3. Link the operator account with a Tendermint node pubkey that will be used for signing blocks. If no
   `--pubkey` flag is provided, it defaults to the local node pubkey created via the `imversed init` command above.

For more information on `gentx`, use the following command:

```linux
imversed gentx --help
```

### Collecting `gentx`
By default, the genesis file do not contain any `gentxs`. A `gentx` is a transaction that bonds staking token present in the genesis file under `accounts` to a validator, essentially creating a validator at genesis. The chain will start as soon as more than 2/3rds of the validators (weighted by voting power) that are the recipient of a valid `gentx` come online after `genesis_time`.

A `gentx` can be added manually to the genesis file, or via the following command:

```linux
# Add the gentx to the genesis file
imversed collect-gentxs
```

This command will add all the gentxs stored in `~/.imversed/config/gentx` to the genesis file.

### Run Testnet
Finally, check the correctness of the `genesis.json` file:

```text
imversed validate-genesis
```

Now that everything is set up, you can finally start your node:

```text
imversed start
```

> To check all the available customizable options when running the node, use the `--help` flag.

You should see blocks come in.

The previous command allow you to run a single node. This is enough for the next section on interacting with this node, but you may wish to run multiple nodes at the same time, and see how consensus happens between them.

You can then stop the node using `Ctrl+C`.