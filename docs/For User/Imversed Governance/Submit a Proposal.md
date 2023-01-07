# Submit a Proposal

If you have a final draft of your proposal ready to submit, you may want to push your proposal live on the testnet first. These are the three primary steps to getting your proposal live on-chain.

1. (Optional) [Hosting supplementary materials](https://docs.imversed.com/users/governance/submitting.html#hosting-supplementary-materials) for your proposal with IPFS (InterPlanetary File System)
2. [Formatting the JSON file](https://docs.imversed.com/users/governance/submitting.html#formatting-the-json-file-for-the-governance-proposal) for the governance proposal transaction that will be on-chain
3. [Sending the transaction](https://docs.imversed.com/users/governance/submitting.html#sending-the-transaction-that-submits-your-governance-proposal) that submits your governance proposal on-chain

## Hosting supplementary materials
In general we try to minimize the amount of data pushed to the blockchain. Hence, detailed documentation about a proposal is usually hosted on a separate censorship resistant data-hosting platform, like IPFS.

Once you have drafted your proposal, ideally as a Markdown file, you can upload it to the IPFS network:

1. either by [running an IPFS node and the IPFS softwar](https://ipfs.io/), or
2. using a service such as [pinata.cloud](https://pinata.cloud).

Ensure that you "pin" the file so that it continues to be available on the network. You should get a URL like this: `https://ipfs.io/ipfs/QmbkQNtCAdR1CNbFE8ujub2jcpwUcmSRpSCg8gVWrTHSWD` The value `QmbkQNtCAdR1CNbFE8ujub2jcpwUcmSRpSCg8gVWrTHSWD` is called the `CID` of your file - it is effectively the file's hash.

If you uploaded a markdown file, you can use the IPFS markdown viewer to render the document for better viewing. Links for the markdown viewer look like `https://ipfs.io/ipfs/QmTkzDwWqPbnAh5YiV5VwcTLnGdwSNsNTn2aDxdXBFca7D/example#/ipfs/<CID>`, where `<CID>` is your CID. For instance the link above would be: https://ipfs.io/ipfs/QmTk...HSWD

Share the URL with others and verify that your file is publicly accessible.

The reason we use IPFS is that it is a decentralized means of storage, making it resistant to censorship or single points of failure. This increases the likelihood that the file will remain available in the future.

## Formatting the JSON file for the governance proposal
Many proposals allow for long form text to be included, usually under the key `description`. These provide the opportunity to include [markdown](https://www.markdownguide.org/) if formatted correctly as well as line breaks with `\n`. Beware, however, that if you are using the CLI to create a proposal, and setting `description` using a flag, the text will be [escaped](https://en.wikipedia.org/wiki/Escape_sequences_in_C) which may have undesired effects. If you're using markdown or line breaks it's recommended to put the proposal text into a json file and include that file as part of the CLI proposal, as opposed to individual fields in flags.

### Text Proposals
`TextProposals` are used by delegators to agree to a certain strategy, plan, commitment, future upgrade, or any other statement in the form of text. Aside from having a record of the proposal outcome on the Imversed chain, a text proposal has no direct effect on Imversed.

### Community Pool Spend Proposals
For community pool spend proposals, there are five components:

1. **Title** - the distinguishing name of the proposal, typically the way the that explorers list proposals
2. **Description** - the body of the proposal that further describes what is being proposed and details surrounding the proposal
3. **Recipient** - the Imversed (bech32-based) address that will receive funding from the Community Pool
4. **Amount** - the amount of funding that the recipient will receive in atto-IMV (`aimv`)
5. **Deposit** - the amount that will be contributed to the deposit (in `aimv`) from the account submitting the proposal

#### Made-Up Example
In this simple example (below), a network explorer will list the governance proposal as a `CommunityPoolSpendProposal`. When an observer selects the proposal, they'll see the description. Not all explorers will show the recipient and amount, so ensure that you verify that the description aligns with the what the governance proposal is programmed to enact. If the description says that a certain address will receive a certain number of IMV, it should also be programmed to do that, but it's possible that that's not the case (accidentally or otherwise).

The `amount` is `1000000000000000000aimv`. This is equal to 1 IMV, so `recipient` address `imv1mx9nqk5agvlsvt2yc8259nwztmxq7zjq50mxkp` will receive 1 IMV if this proposal is passed.

The `deposit` of `192000000000000000000aimv` results in 192 IMV being used from the proposal submitter's account. A minimum deposit is required for a proposal to enter the voting period, and anyone may contribute to this deposit within 3 days. If the minimum deposit isn't reached before this time, the deposit amounts will be burned. Deposit amounts will also be burned if the quorum isn't met in the vote or the proposal is vetoed.

```json
{
  "title": "Community Pool Spend",
  "description": "This is the summary of the key information about this proposal. Include the URL to a PDF version of your full proposal.",
  "recipient": "imv1mx9nqk5agvlsvt2yc8259nwztmxq7zjq50mxkp",
  "amount": [
    {
      "denom": "aimv",
      "amount": "1000000000000000000"
    }
  ],
  "deposit": "64000000000000000000aimv"
}
```

#### Real Example
This is a governance protocol which [Flux Protocol](https://www.fluxprotocol.org/), the provider of a cross-chain oracle which provides smart contracts with access to economically secure data feeds, submitted to cover costs of the subsidizied FPO (First Party Oracle) solution which they deployed on the Imversed mainnet.

Users can query the proposal details with the imversed command-line interface using this command:

```linux
`imversed --node https://tx-endpoint-test.imversed.com:443 query gov proposal 23`.
```

```json
{
  "id": "2",
  "messages": [
    {
      "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
      "content": {
        "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
        "title": "v3.5",
        "description": "upgrade",
        "plan": {
          "name": "v3.5",
          "time": "0001-01-01T00:00:00Z",
          "height": "1322760",
          "info": "",
          "upgraded_client_state": null
        }
    },
    "authority": "imv10d07y265gmmuvt4z0w9aw880jnsr700jemq2lk"
    }
  ],
  "status": "PROPOSAL_STATUS_PASSED",
  "final_tally_result": {
  "yes_count": "500000000000000000000000",
  "abstain_count": "0",
  "no_count": "0",
  "no_with_veto_count": "0"
},
"total_deposit": [
  {
    "denom": "aimv",
    "amount": "500000000"
  }
],
"metadata": ""
}
```

### Params-Change Proposals
> Changes to the [`gov` module](https://docs.imversed.com/users/governance/overview.html) are different from the other kinds of parameter changes because gov has subkeys, [as discussed here](https://github.com/cosmos/cosmos-sdk/issues/5800). Only the key part of the JSON file is different for gov parameter-change proposals.

For parameter-change proposals, there are seven components:

1. **Title** - the distinguishing name of the proposal, typically the way the that explorers list proposals
2. **Description** - the body of the proposal that further describes what is being proposed and details surrounding the proposal
3. **Subspace** - the Imversed module with the parameter that is being changed
4. **Key** - the parameter that will be changed
5. **Value** - the value of the parameter that will be changed by the governance mechanism
6. **Denom** - `aimv` (atto-IMV) will be the type of asset used as the deposit
7. **Amount** - the amount that will be contributed to the deposit (in aimv) from the account submitting the proposal

#### Real Example
In the example below, a network explorer listed the governance proposal by its title: "Increase the minimum deposit for governance proposals." When a user selects the proposal, they'll see the proposal’s description.
Not all explorers will show the proposed parameter changes that are coded into the proposal, so the delegator should verify that the description aligns with what the governance proposal is programmed to enact. If the description says that a certain parameter will be increased, it should also be programmed to do that, but it's possible that that's not the case (accidentally or otherwise).

Users can query the proposal details with the imversed command-line interface using this command:

```linux
imversed --node https://tx-endpoint-test.imversed.com:443 query gov proposal 2
```
```json
{
  "id": "2",
  "messages": [
  {
      "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
      "content": {
        "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
        "title": "v3.5",
        "description": "upgrade",
        "plan": {
          "name": "v3.5",
          "time": "0001-01-01T00:00:00Z",
          "height": "1322760",
          "info": "",
          "upgraded_client_state": null
        }
      },
      "authority": "imv10d07y265gmmuvt4z0w9aw880jnsr700jemq2lk"
    }
  ],
  "status": "PROPOSAL_STATUS_PASSED",
  "final_tally_result": {
    "yes_count": "500000000000000000000000",
    "abstain_count": "0",
    "no_count": "0",
    "no_with_veto_count": "0"
  },
  "total_deposit": [
    {
    "denom": "aimv",
    "amount": "500000000"
    }
  ],
  "metadata": ""
}
```

The deposit `denom` is `aimv` and `amount` is `20100000000000000000`. Therefore, a deposit of 20.1 IMV will be included with this proposal. At the time, the IMV mainnet had a 10 IMV minimum deposit, so this proposal was put directly into the voting period (and subsequently passed). The minimum deposit amount is currently 192 IMV. There is a minimum deposit required for a proposal to enter the voting period, and anyone may contribute to this deposit within a 3-day period. If the minimum deposit isn't reached before this time, the deposit amounts will be burned.

### Sending the transaction that submits your governance proposal

For information on how to use `imversed` binary to submit an on-chain proposal through the governance module, please refer to the [quickstart](https://docs.imversed.com/validators/quickstart/binary.html) documentation.

#### CLI
This is the command format for using imversed (the command-line interface) to submit your proposal on-chain:

```linux
imversed tx gov submit-proposal \
  --title=<title> \
  --description=<description> \
  --type="Text" \
  --deposit="1000000aimv" \
  --from=<mykey> \
  --chain-id=<chain_id>
  --node <address>
```

> Use the `imversed tx gov --help` flag to get more info about the governance commands

1. `imversed` is the command-line interface client that is used to send transactions and query Imversed
2. `tx gov submit-proposal param-change` indicates that the transaction is submitting a parameter-change proposal
3. `--from mykey` is the account key that pays the transaction fee and deposit amount
4. `--gas 500000` is the maximum amount of gas permitted to be used to process the transaction
   * the more content there is in the description of your proposal, the more gas your transaction will consume
   * if this number isn't high enough and there isn't enough gas to process your transaction, the transaction will fail
   * the transaction will only use the amount of gas needed to process the transaction
5. `--gas-prices` is the flat-rate per unit of gas value for a validator to process your transaction
6. `--chain-id imversed_5555555-1` is Imversed Mainnet. For current and past chain-id's, please look at the Chain ID documentation.
   * the testnet chain ID is imversed_5555558-1
7. `--node` is using a full node to send the transaction to the Imversed Mainnet

### Verifying your transaction
After posting your transaction, your command line interface (`imversed`) will provide you with the transaction's hash, which you can either query using `imversed` or by searching the transaction hash using [Mintscan](https://txe.imversed.com/) or any block explorer.

### Depositing funds after a proposal has been submitted
Sometimes a proposal is submitted without having the minimum token amount deposited yet. In these cases you would want to be able to deposit more tokens to get the proposal into the voting stage. In order to deposit tokens, you'll need to know what your proposal ID is after you've submitted your proposal. You can query all proposals by the following command:

```linux
imversed q gov proposals
```

If there are a lot of proposals on the chain already, you can also filter by your own address. For the proposal above, that would be:

```linux
imversed q gov proposals --depositor imv159dctezcw6u077gwthl0ytfgtj2nxufs53c5lg
```

Once you have the proposal ID, this is the command to deposit extra tokens:

```linux
imversed tx gov deposit <proposal-id> <deposit> --from <name>
```

In our case above, the `<proposal-id>` would be 59 as queried earlier. The `<deposit>` is written as `500000aimv`, just like the example above.

### Submit your proposal to the testnet
You may want to submit your proposal to the testnet chain before the mainnet for a number of reasons:

1. To see what the proposal description will look like
2. To signal that your proposal is about to go live on the mainnet
3. To share what the proposal will look like in advance with stakeholders
4. To test the functionality of the governance features

Submitting your proposal to the testnet increases the likelihood that you will discover a flaw before deploying your proposal on mainnet. A few things to keep in mind:

* you'll need testnet tokens for your proposal (ask around for a [faucet](https://docs.imversed.com/developers/testnet/faucet.html))
* the parameters for testnet proposals are different (eg. voting period timing, deposit amount, deposit denomination)
* the deposit denomination is in `'atimv'` instead of `'aimv'`
