---
description: Intoruction to Users.
helpfulVotes: false
---

# Account Keys

## Keyring

```List

SYNOPSIS
Create, import, export and delete keys using the CLI keyring

```

The keyring holds the private/public keypairs used to interact with the node. For instance, a validator key needs to be set up before running the node, so that blocks can be correctly signed. The private key can be stored in different locations, called "backends", such as a file or the operating system's own key storage.

## Add keys

You can use the following commands for help with the `keys` command and for more information about a particular subcommand, respectively:

```List

Coming soon

```

```List

Coming soon

```

To create a new key in the keyring, run the `add` subcommand with a `<key_name>` argument. You will have to provide a password for the newly generated key. This key will be used in the next section.

```List

Coming soon

```

This command generates a new 24-word mnemonic phrase, persists it to the relevant backend, and outputs information about the keypair. If this keypair will be used to hold value-bearing tokens, be sure to write down the mnemonic phrase somewhere safe!

By default, the keyring generates a `eth_secp256k1` key. The keyring also supports `ed25519` keys, which may be created by passing the `--algo` flag. A keyring can of course hold both types of keys simultaneously.


::: tip 

Note: The Ethereum address associated with a public key can be derived by taking the full Ethereum public key of type eth_secp256k1, computing the Keccak-256 hash, and truncating the first twelve bytes.

:::



<!-- ::: tip  

NOTE: Cosmos `secp256k1` keys are not supported on Imversed due to compatibility issues with Ethereum transactions.

:::

-->

## Keyring Backends

### OS

::: tip

`os` is the default option since operating system's default credentials managers are designed to meet users' most common needs and provide them with a comfortable experience without compromising on security.

:::

The `os` backend relies on operating system-specific defaults to handle key storage securely. Typically, an operating system's credential sub-system handles password prompts, private keys storage, and user sessions according to the user's password policies. Here is a list of the most popular operating systems and their respective passwords manager:

- macOS (since Mac OS 8.6): [Keychain](https://support.apple.com/en-gb/guide/keychain-access/welcome/mac)
- Windows: [Credentials Management API](https://docs.microsoft.com/en-us/windows/win32/secauthn/credentials-management)
- GNU/Linux:
    - [libsecret](https://gitlab.gnome.org/GNOME/libsecret)
    - [kwallet](https://api.kde.org/frameworks/kwallet/html/index.html)


GNU/Linux distributions that use GNOME as default desktop environment typically come with [Seahorse](https://wiki.gnome.org/Apps/Seahorse). Users of KDE based distributions are commonly provided with [KDE Wallet Manager](https://userbase.kde.org/KDE_Wallet_Manager). Whilst the former is in fact a `libsecret` convenient frontend, the latter is a `kwallet` client.

The recommended backends for headless environments are `file` and `pass`.


### File

The `file` stores the keyring encrypted within the app's configuration directory. This keyring will request a password each time it is accessed, which may occur multiple times in a single command resulting in repeated password prompts. If using bash scripts to execute commands using the `file` option you may want to utilize the following format for multiple prompts:

```List

Coming soon

```

::: tip

The first time you add a key to an empty keyring, you will be prompted to type the password twice.

:::

### Password Store

The `pass` backend uses the [pass](https://www.passwordstore.org/) utility to manage on-disk encryption of keys' sensitive data and metadata. Keys are stored inside `gpg` encrypted files within app-specific directories. `pass` is available for the most popular UNIX operating systems as well as GNU/Linux distributions. Please refer to its manual page for information on how to download and install it.

::: tip

`pass` uses [GnuPG](https://gnupg.org/) for encryption. `gpg` automatically invokes the gpg-agent daemon upon execution, which handles the caching of GnuPG credentials. Please refer to `gpg-agent` man page for more information on how to configure cache parameters such as credentials TTL and passphrase expiration.

:::

The password store must be set up prior to first use:

```List

pass init <GPG_KEY_ID>

```

Replace `<GPG_KEY_ID>` with your GPG key ID. You can use your personal GPG key or an alternative one you may want to use specifically to encrypt the password store.

### KDE Wallet Manager

The `kwallet` backend uses `KDE Wallet Manager`, which comes installed by default on the GNU/Linux distributions that ships KDE as default desktop environment. Please refer to [KWallet Handbook](https://docs.kde.org/stable5/en/kwalletmanager/kwallet5/) for more information.

### Testing

The `test` backend is a password-less variation of the `file` backend. Keys are stored **unencrypted** on disk. This keyring is provided for testing purposes only. Use at your own risk!


::: tip

🚨 DANGER: **Never** create your mainnet validator keys using a `test` keying backend. Doing so might result in a loss of funds by making your funds remotely accessible via the `eth_sendTransaction` JSON-RPC endpoint.

Ref: [Security Advisory: Insecurely configured geth can make funds remotely accessible](https://blog.ethereum.org/2015/08/29/security-alert-insecurely-configured-geth-can-make-funds-remotely-accessible)

:::

### In Memory

The `memory` backend stores keys in memory. The keys are immediately deleted after the program has exited.

::: tip

**IMPORTANT**: Provided for testing purposes only. The `memory` backend is **not** recommended for use in production environments. Use at your own risk!

:::

## Multisig

```List

SYNOPSIS
Learn how to generate, sign and broadcast a transaction using the keyring multisig

```

A **multisig** account is an Imversed account with a special key that can require more than one signature to sign transactions. This can be useful for increasing the security of the account or for requiring the consent of multiple parties to make transactions. Multisig accounts can be created by specifying:

- threshold number of signatures required
- the public keys involved in signing

To sign with a multisig account, the transaction must be signed individually by the different keys specified for the account. Then, the signatures will be combined into a multisignature which can be used to sign the transaction. If fewer than the threshold number of signatures needed are present, the resultant multisignature is considered invalid.


## Generate a Multisig key

```List

Coming soon

```

`K` is the minimum number of private keys that must have signed the transactions that carry the public key's address as signer.

The `--multisig` flag must contain the name of public keys that will be combined into a public key that will be generated and stored as `new_key_name` in the local database. All names supplied through `--multisig` must already exist in the local database.

Unless the flag `--nosort` is set, the order in which the keys are supplied on the command line does not matter, i.e. the following commands generate two identical keys:

```List

Coming soon

```

Multisig addresses can also be generated on-the-fly and printed through the which command:

```List

Coming soon

```

## Signing a transaction

### Step 1: Create the multisig key

Let's assume that you have `test1` and `test2` want to make a multisig account with `test3`.

First import the public keys of `test3` into your keyring.

```List

Coming soon

```
Generate the multisig key with 2/3 threshold.

```List

Coming soon

```

You can see its address and details:

```List

Coming soon

```

Let's add 10 Imversed to the multisig wallet:

```List

Coming soon

```

### Step 2: Create the multisig transaction

We want to send 5 Imversed from our multisig account to `Adress will Coming soon`

The file `unsignedTx.json` contains the unsigned transaction encoded in JSON.

```List

Coming soon

```

### Step 3: Sign individually

Sign with `test1` and `test2` and create individual signatures.

```List

Coming soon

```

```List

Coming soon

```

### Step 4: Create multisignature

Combine signatures to sign transaction.

```List

Coming soon

```

The TX is now signed:

```List

Coming soon

```

Step 5: Broadcast transaction

```List

Coming soon

```


