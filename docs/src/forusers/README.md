---
description: Intoruction to Users.
helpfulVotes: false
---

# Basic concept

## Transactions





Transactions are objects created by end-users to trigger state changes in the application.

::: tip

This documentation page is currently under work in progress.

:::

## Transaction Confirmations


## Transaction Types



## Ethereum Transactions

Ethereum transactions refer to actions initiated by EOAs (externally-owned accounts, managed by humans), rather than internal smart contract calls. Ethereum transactions transform the state of the EVM and therefore must be broadcasted to the entire network.

Ethereum transactions also require a fee, known as `gas`. [(EIP-1559)](https://eips.ethereum.org/EIPS/eip-1559) 
introduced the idea of a base fee, along with a priority fee which serves as an incentive for miners to include specific transactions in blocks.

There are several categories of Ethereum transactions:

- regular transactions: transactions from one account to another
- contract deployment transactions: transactions without a `to` address, where the contract code is sent in the `data` field
- execution of a contract: transactions that interact with a deployed smart contract, where the `to` address is the smart contract address

For more information on Ethereum transactions and the transaction lifecycle, [go here](https://ethereum.org/en/developers/docs/transactions/)

Imversed supports the following Ethereum transactions.

::: tip 

Note: Unprotected legacy transactions are not supported by default.

:::

- Dynamic Fee Transactions [(EIP-1559)](https://eips.ethereum.org/EIPS/eip-1559)
- Access List Transactions [(EIP-2930)](https://eips.ethereum.org/EIPS/eip-2930)
- Legacy Transactions [(EIP-2718)](https://eips.ethereum.org/EIPS/eip-2718)

## Interchain Transactions




## Transaction Receipts

# Tokens

::: tip

Learn about the the different types of tokens available in Imversed.

:::

## Introduction

Imversed is a Cosmos-based chain with full Ethereum Virtual Machine (EVM) support. Because of this [architecture], tokens and assets in the network may come from different independent sources.

## The Imversed Token

The denomination used for staking, governance and gas consumption on the EVM is the Imversed. The Imversed provides the utility of: securing the Proof-of-Stake chain, token used for governance proposals, distribution of fees to validator and users, and as a mean of gas for running smart contracts on the EVM.

Imversed uses [Atto](https://en.wikipedia.org/wiki/Atto-) as the base denomination to maintain parity with Ethereum.

1 IMV =  1*10^18 AIMV

This matches Ethereum denomination of:

1 ETH = 1018 wei

## Imversed Coins

Accounts can own Cosmos coins in their balance, which are used for operations with other Imversed and transactions. Examples of these are using the coins for staking, IBC transfers, governance deposits and IVM.

## IVM Tokens

Imversed is compatible with ERC20 tokens and other non-fungible token standards (EIP721, EIP1155) that are natively supported by the EVM.

## Gas and Fees


The concept of Gas represents the amount of computational effort required to execute specific operations on the state machine.

Gas was created on Ethereum to disallow the EVM (Ethereum Virtual Machine) from running infinite loops by allocating a small amount of monetary value into the system. A unit of gas, usually in the form of a fraction of the native coin, is consumed for every operation on the EVM and requires a user to pay for these operations. These operations consist in state transitions such as sending a transaction or calling a contract.

Exactly like Ethereum, Imvered utilizes the concept of gas and this is how Imversed tracks the resource usage of operations during execution. Operations on Imversed are represented as read or writes done to the chain's store.

In Cosmos, a fee is calculated and charged to the user during a message execution. This fee is calculated from the sum of all gas consumed in a message execution. So, the fee is equivalent to the gas multiplied by the gas price.

In both networks, gas is used to make sure that operations do not require an excess amount of computational power to complete and as a way to deter bad-acting users from spamming the network.


## Imversed Gas

In the Imversed SDK, gas is tracked in the main `GasMeter` and the `BlockGasMeter`:

- `GasMeter`: keeps track of the gas consumed during executions that lead to state transitions. It is reset on every transaction execution.
- `BlockGasMeter`: keeps track of the gas consumed in a block and enforces that the gas does not go over a predefined limit. This limit is defined in the Tendermint consensus parameters and can be changed via governance parameter change proposals.

More information regarding gas in Imversed SDK can be found `here (link to page)`.

## Matching EVM Gas consumption

Imversed is an EVM-compatible chain that supports Ethereum Web3 tooling. For this reason, gas consumption must be equitable with other EVMs, most importantly Ethereum.

The main difference between EVM and Imversed state transitions, is that the EVM uses a [gas table](https://github.com/ethereum/go-ethereum/blob/master/params/protocol_params.go) for each OPCODE, whereas Imversed uses a `GasConfig` that charges gas for each CRUD operation by setting a flat and per-byte cost for accessing the database.

```list
// GasConfig defines gas cost for each operation on KVStores
type GasConfig struct {
	HasCost          Gas
	DeleteCost       Gas
	ReadCostFlat     Gas
	ReadCostPerByte  Gas
	WriteCostFlat    Gas
	WriteCostPerByte Gas
	IterNextCostFlat Gas
}
```

In order to match the gas consumed by the EVM, the gas consumption logic from the SDK is ignored, and instead the gas consumed is calculated by subtracting the state transition leftover gas plus refund from the gas limit defined on the message.

To ignore the SDK gas consumption, we reset the transaction `GasMeter` count to 0 and manually set it to the `gasUsed` value computed by the EVM module at the end of the execution.

```list

package keeper

This documentation part is currently under work in progress.

```

# `AnteHandler`

The Imversed SDK `AnteHandler (Link to page)` performs basic checks prior to transaction execution. These checks are usually signature verification, transaction field validation, transaction fees, etc.

Regarding gas consumption and fees, the `AnteHandler` checks that the user has enough balance to cover for the tx cost (amount plus fees) as well as checking that the gas limit defined in the message is greater or equal than the computed intrinsic gas for the message.

## Gas Refunds

In the EVM, gas can be specified prior to execution. The totality of the gas specified is consumed at the beginning of the execution (during the `AnteHandler` step) and the remaining gas is refunded back to the user if any gas is left over after the execution. Additionally the EVM can also define gas to be refunded back to the user but those will be capped to a fraction of the used gas depending on the fork/version being used.


## O Fee Transactions

In Imversed, a minimum gas price is not enforced by the `AnteHandler` as the `min-gas-prices` is checked against the local node/validator. In other words, the minimum fees accepted are determined by the validators of the network, and each validator can specify a different minimum value for their fees. This potentially allows end users to submit 0 fee transactions if there is at least one single validator that is willing to include transactions with `0` gas price in their blocks proposed.

For this same reason, in Imversed it is possible to send transactions with `0` fees for transaction types other than the ones defined by the `emv` module. EVM module transactions cannot have `0` fees as gas is required inherently by the EVM. This check is done by the EVM transactions stateless validation (i.e `ValidateBasic`) function as well as on the custom `AnteHandler` defined by Imversed.

## Gas estimation

Ethereum provides a JSON-RPC endpoint `eth_estimateGas` to help users set up a correct gas limit in their transactions.

Unfortunately, we cannot make use of the SDK tx simulation for gas estimation because the pre-check in the Ante Handlers would require a valid signature, and the sender balance to be enough to pay for the gas. But in Ethereum, this endpoint can be called without specifying any sender address.

For that reason, a specific query API `EstimateGas` is implemented in Imversed. It will apply the transaction against the current block/state and perform a binary search in order to find the optimal gas value to return to the user (the same transaction will be applied over and over until we find the minimum gas needed before it fails). The reason we need to use a binary search is that the gas required for the transaction might be higher than the value returned by the EVM after applying the transaction, so we need to try until we find the optimal value.

A cache context will be used during the whole execution to avoid changes be persisted in the state.

``` List

package keeper

This documentation part is currently under work in progress.

```

## Smart contracts

`Learn more about smart contracts and how are they supported on Imversed`


::: tip

This documentation page is currently under work in progress.

:::



