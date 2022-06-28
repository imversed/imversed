#!/bin/bash

KEY="validator"

CHAINID="imversed_1234-1"
MONIKER="localtestnet"

KEYRING="test"
KEYALGO="eth_secp256k1"
MNEMONIC="kangaroo buffalo margin access fiscal manage firm coral case tattoo salt stadium crystal kid poverty document confirm coach bronze use cram uphold bridge input"

LOGLEVEL="debug"
# to trace evm
TRACE="--trace"
# TRACE=""

## validate dependencies are installed
#command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.imversed
rm -rf ~/go/bin/imversed
#
#ignite chain build
starport c build

#mv ~/go/bin/imversedd ~/go/bin/imversed

#make install

~/go/bin/imversed config keyring-backend $KEYRING
~/go/bin/imversed config chain-id $CHAINID

# if $KEY exists it should be deleted
yes "$MNEMONIC" | ~/go/bin/imversed keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
~/go/bin/imversed init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to aimv
cat $HOME/.imversed/config/genesis.json | jq '.app_state["bank"]["denom_metadata"] += [{"name":"Imversed Token","symbol":"IMV","base":"aimv","display":"imv","description":"The native staking token of the Imversed.","denom_units":[{"denom":"aimv","exponent":0,"aliases":["attoimversed"]},{"denom":"imv","exponent":18,"aliases":["imversed"]}]}]' > $HOME/.imversed/config/tmp_genesis.json && mv $HOME/.imversed/config/tmp_genesis.json $HOME/.imversed/config/genesis.json
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
~/go/bin/imversed gentx $KEY 1000000000000000000000aimv 0x4e34a7E46e4Ad9756D03E8Fc40b605C5b023555F imv1f0en8j95fwh7rpljh7pk3tw8442spl2wvlm5lr 0xbba7a8fbf5fff2763828d2b7da885cbb760a3cc064f4117bcf9822750b065a623d149c847ffe87e69c4be5d682604fd69a1c50e515bf5b53eeb750be1672be811b --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
~/go/bin/imversed collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
# ~/go/bin/imversed validate-genesis //TODO check it, doesn't work correctly with new gentx command

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
#~/go/bin/imversed start --pruning=nothing --evm.tracer=json $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001aimv --json-rpc.api eth,txpool,personal,net,debug,web3,miner --api.enable
~/go/bin/imversed start
