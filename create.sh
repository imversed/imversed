#!/bin/bash
set -eu

PATH=build:$PATH

metachaind init val-one --chain-id my-test-chain

metachaind keys add val-one --keyring-backend os

# Put the generated address in a variable for later use.
MY_VALIDATOR_ADDRESS=$(metachaind keys show val-one -a --keyring-backend os)

metachaind add-genesis-account meta16flua38pmsfnf2j08hcwa2gdj4t52dcs36wyh3 100000000fulldive

# Create a gentx.
metachaind gentx val-one 100000000fulldive --chain-id metachain --keyring-backend os

# Add the gentx to the genesis file.
metachaind collect-gentxs

metachaind start