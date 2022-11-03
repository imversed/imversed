# Chain ID

## Official Chain IDs
> NOTE: The latest Chain ID (i.e highest Version Number) is the latest version of the software and mainnet.

:::: tabs ::: tab Mainnet

| **Name** | **Chain ID** | **Identifier** | **EIP155 Number** | **Version Number** |
| --- | --- | --- | --- | --- |
| Imversed Canary | `5555555` | `imv` | `5555555` | `1` |

::: ::: tab Testnets

| **Name** | **Chain ID** | **Identifier** | **EIP155 Number** | **Version Number** |
| --- | --- | --- | --- | --- |
| Imversed Testnet | `5555558` | `ivm` | `5555558` | `1` |

::: ::::

> You can also lookup the [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) Chain ID by referring to [chainlist.org](https://chainlist.org/).

## The Chain Identifier
Every chain must have a unique identifier or chain-id. Tendermint requires each application to define its own `chain-id` in [the genesis.json fields](https://docs.tendermint.com/master/spec/core/genesis.html#genesis-fields). However, in order to comply with both EIP155 and Cosmos standard for chain upgrades, Imversed-compatible chains must implement a special structure for their chain identifiers.

## Structure
The Imversed Chain ID contains 3 main components

* **Identifier:** Unstructured string that defines the name of the application.
* **EIP155 Number:** Immutable [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `CHAIN_ID` that defines the replay attack protection number.
* **Version Number:** Is the version number (always positive) that the chain is currently running. This number **MUST** be incremented every time the chain is upgraded or forked in order to avoid network or consensus errors.

### Format
The format for specifying and Imversed compatible chain-id in genesis is the following:

```text
{identifier}_{EIP155}-{version}
```
The following table provides an example where the second row corresponds to an upgrade from the first one:

| ChainID | Identifier | EIP155 Number | Version Number |
| --- | --- | --- | --- |
| imv_5555555 _1 | imv | 5555555 | 1 |