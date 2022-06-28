#!/bin/bash

KEY="validator"
CHAINID="imversed_1234-1"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="debug"
# to trace evm
TRACE="--trace"
# TRACE=""

## validate dependencies are installed
#command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.imversed
#rm -fr ~/go/bin/imversed
#
#starport chain build
#
#mv ~/go/bin/imversedd ~/go/bin/imversed

#make install

~/go/bin/imversed config keyring-backend $KEYRING
~/go/bin/imversed config chain-id $CHAINID

# if $KEY exists it should be deleted
~/go/bin/imversed keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
~/go/bin/imversed init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to aimv
cat $HOME/.imversed/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aimv"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json
cat $HOME/.imversed/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aimv"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json
cat $HOME/.imversed/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aimv"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json
cat $HOME/.imversed/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="aimv"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json

# increase block time (?)
cat $HOME/.imversed/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="1000"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json

# Set gas limit in genesis
cat $HOME/.imversed/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json
sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = "0.001aimv"/g' ~/.imversed/config/app.toml &&
cat <<< $(jq '.app_state.gov.voting_params.voting_period = "30s"' $HOME/.imversed/config/genesis.json) > $HOME/.imversed/config/genesis.json &&
sed -i -e 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3,miner"/g' ~/.imversed/config/app.toml &&

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.imversed/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.imversed/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.imversed/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.imversed/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.imversed/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.imversed/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
~/go/bin/imversed add-genesis-account $KEY 100000000000000000000000000aimv --keyring-backend $KEYRING

# Sign genesis transaction
~/go/bin/imversed gentx $KEY 1000000000000000000000aimv --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
~/go/bin/imversed collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
~/go/bin/imversed validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin &&
# copy current binary
cp ~/go/bin/imversed $DAEMON_HOME/cosmovisor/genesis/bin &&

starport c build --release &&
tar -zxvf release/imversed_darwin_amd64.tar.gz  &&
#mv imversedd imversed &&

mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v3.2/bin &&
# copy binary with upgrade

cp ~/projects/imversed/imversed $DAEMON_HOME/cosmovisor/upgrades/v3.2/bin &&

~/go/bin/cosmovisor start
