# Validator Backup

It is **crucial** to backup your validator's private key. It's the only way to restore your validator in the event of a disaster.

The validator private key is a Tendermint Key: a unique key used to sign consensus votes.

To backup everything you need to restore your validator, note that if you are using the "software sign" (the default signing method of Tendermint), your Tendermint key is located at:

```text
~/.imversed/config/priv_validator_key.json
```

Then do the following:

1. Backup the `json` file mentioned above (or backup the whole `config` folder).
2. Backup the self-delegator wallet. [See backing up wallets with the Imversed Daemon](https://docs.imversed.com/users/wallets/backup.html).

To see your validator's associated public key:

```text
imversed tendermint show-validator
```

To see your validator's associated `bech32` address:

```text
imversed tendermint show-address
```

You can also use hardware to store your Tendermint Key much more safely, such as [YubiHSM2](https://developers.yubico.com/YubiHSM2/).
