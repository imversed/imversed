---
description: Intoruction to Users.
helpfulVotes: false
---

# Imversed Governance

## Overview

::: tip

Note: Working on a governance proposal? Make sure to look at the best practices.

:::

Imversed has an on-chain governance mechanism for passing text proposals, changing chain parameters, and spending funds from the community pool.

## On- and off-chain Governance Structure

### Communication Methods

Governance practices and decisions are communicated through different types of documents and design artifacts:

- On-chain governance proposals

- Architecture Decision records

- Technical standards / specifications

### Decision-making and Discussion Venues

Venues involve community members to different degrees and individuals often perform multiple roles in the Cosmos ecosystem (validators, users, developers and core-members of Imversed Team). Because technical direction setting and development is almost always happening in the open, involvement from members in the extended community occurs organically.

- [Imversed Discord](https://discord.com/invite/3s7V2npZw4)

    - For ecosystem cross-pollination with an active developer presence.
        - 🏛│governance channel for discussing proposals, upgrades, etc.
        - 📜│proposals channel for a full list of proposals.
        - ⏫│upgrades channel for upcoming software upgrades.


- [Telegram(@Imversed)](https://t.me/+1BUGmHYcU1U1Y2Ji)
    - General Imversed Telegram group

- [Twitter(imversedhub)](https://twitter.com/imversedhub) 
    - Official Imversed Twitter

## Proposal Process

### Deposit Period

The deposit period lasts either 5 days or until the proposal deposit totals 64 Imversed, whichever happens first.

### Deposits

Deposit amounts are at risk of being burned. Prior to a governance proposal entering the voting period (ie. for the proposal to be voted upon), there must be at least a minimum number of Imversed deposited (64). Anyone may contribute to this deposit. Deposits of passed and failed proposals are returned to the contributors.

In the past, different people have considered contributions amounts differently. There is some consensus that this should be a personal choice. There is also some consensus that this can be an opportunity for supporters to signal their support by adding to the deposit amount, so a proposer may choose to leave contribution room (ie. a deposit below 64 Imversed) so that others may participate. It is important to remember that any contributed Imversed are at risk of being burned.

### Burned deposits

Deposits are burned when proposals:

1. **Expire** - deposits will be burned if the deposit period ends before reaching the minimum deposit (64 Imversed)

2. **Fail to reach quorum** - deposits will be burned for proposals that do not reach quorum ie. 33.4% of all staked Imversed must vote

3. **Are vetoed** - deposits for proposals with 33.4% of voting power backing the `NoWithVeto` option are also burned

### Voting Period

The voting period is currently a fixed 5-day period. During the voting period, participants may select a vote of either `Yes`, `No`, `Abstain`, or `NoWithVeto`. Voters may change their vote at any time before the voting period ends.


### What do the voting options mean?

1. `Abstain`: indicates that the voter is impartial to the outcome of the proposal.

2. `Yes`: indicates approval of the proposal in its current form.

3. `No`: indicates disapproval of the proposal in its current form.

4. `NoWithVeto`: indicates stronger opposition to the proposal than simply voting No. If the number of `NoWithVeto` votes is greater than a third of total votes excluding `Abstain` votes, the proposal is rejected and the deposits are burned.

As accepted by the community in Proposal 6 `(Link to website)`, voters are expected to vote `NoWithVeto` if a proposal leads to undesirable outcomes for the community. It states “if a proposal seems to be spam or is deemed to have caused a negative externality to Imversed community, voters should vote `NoWithVeto`.”

Voting `NoWithVeto` provides a mechanism for a minority group representing a *third* of the participating voting power to reject a proposal that would otherwise pass. This makes explicit an aspect of the consensus protocol: it works as long as only up to [a third of nodes fail](https://docs.tendermint.com/v0.35/introduction/what-is-tendermint.html). In other words, greater than a third of validators are always in a position to cause a proposal to fail outside the formalized governance process and the network's norms, such as by censoring transactions. The purpose of internalizing this aspect of the consensus protocol into the governance process is to discourage validators from relying on collusion and censorship tactics to influence voting outcomes.

### What determines whether or not a governance proposal passes?

There are four criteria:

1. A minimum deposit of 64 Imversed is required for the proposal to enter the voting period
    - anyone may contribute to this deposit
    - the deposit must be reached within 14 days (this is the deposit period)

2. A minimum of 33.4% of the network's voting power (quorum) is required to participate to make the proposal valid

3. A simple majority (greater than 50%) of the participating voting power must back the `Yes` vote during the 5-day voting period

4. Less than 33.4% of participating voting power votes `NoWithVeto`

Currently, the criteria for submitting and passing/failing all proposal types is the same.

### How is voting tallied?

Voting power is determined by stake weight at the end of the 5-day voting period and is proportional to the number of total Imversed participating in the vote. Only bonded Imversed count towards the voting power for a governance proposal. Liquid Imversed will not count toward a vote or quorum.

Inactive validators can cast a vote, but their voting power (including the backing of their delegators) will not count toward the vote if they are not in the active set when the voting period ends. That means that if I delegate to a validator that is either jailed, tombstoned, or ranked lower than 150 in stake-backing at the time that the voting period ends, my stake-weight will not count in the vote.

Though a simple majority `Yes` vote (ie. 50% of participating voting power) is required for a governance proposal vote to pass, a `NoWithVeto` vote of 33.4% of participating voting power or greater can override this outcome and cause the proposal to fail. This enables a minority group representing greater than 1/3 of voting power to fail a proposal that would otherwise pass.

### How is quorum determined?

Voting power, whether backing a vote of `Yes`, `Abstain`, `No`, or `NoWithVeto`, counts toward quorum. Quorum is required for the outcome of a governance proposal vote to be considered valid and for deposit contributors to recover their deposit amounts. If the proposal vote does not reach quorum (ie. less than 33.4% of the network's voting power is participating) within 5 days, any deposit amounts will be burned and the proposal outcome will not be considered to be valid.

## Best Practices

::: tip

Note:

- If users are creating governance proposals which require community pool funding (such as those of type CommunityPoolSpendProposal), refer to this section.

- If users are creating governance proposals concerned with the ERC-20 Module (such as those of type RegisterCoinProposal), refer to this section.

- If users are creating governance proposals concerned with changing parameters (such as those of type ParameterChangeProposal), refer to this section.

:::

### General Advice: Community Outreach

Engagement is likely to be critical to the success of a proposal. The degree to which you engage with Imversed community should be relative to the potential impact that your proposal may have on the stakeholders. This guide does not cover all ways of engaging: you could bring your idea to a podcast or a hackathon, host an AMA on [Reddit](https://www.reddit.com/r/imversed/) or host a Q&A (questions & answers). We encourage you to experiment and use your strengths to introduce proposal ideas and gather feedback.

There are many different ways to engage. One strategy involves a few stages of engagement before and after submitting a proposal on chain. **Why do it in stages?** It's a more conservative approach to save resources. The idea is to check in with key stakeholders at each stage before investing more resources into developing your proposal.

In the first stage of this strategy, you should engage people (ideally experts) informally about your idea. You'll want to start with the minimal, critical components (name, value to cosmos hub, timeline, any funding needs) and check:

- Does it make sense?
- Are there critical flaws?
- Does it need to be reconsidered?

You should be able engaging with key stakeholders (eg. a large validator operator) with a few short sentences to measure their support. Here's an example:

```List
"We are considering a proposal for funding to work on (project). 
We think it will help Imversed to (outcome). 
Timeline is (x), and we're asking for (y) amount. Do you think that this is a proposal 
that (large validator) may support?"
```

**Why a large validator?** They tend to be the de facto decision-makers on Imversed, since their delegators also delegate their voting power. If you can establish a base layer of off-chain support, you can be more confident that it's worth proceeding to the next stage.

:::tip

**Note**: many will likely hesitate to commit support, and that's okay. It will be important to reassure these stakeholders that this isn't a binding a commitment. You're just canvassing the community to get a feel for whether it's worthwhile to proceed. It's also an opportunity to connect with new people and to answer their questions about what it is you're working on. It will be important for them to clearly understand why you think what you're proposing will be valuable to Imversed, and if possible, why it will be valuable to them as long-term stakeholders.

:::

- If you're just developing your idea, start at Stage 1.
- If you're already confident about your idea, skip to Stage 2.
- If you've drafted your proposal, engaged with the community, and submitted your proposal to the testnet, skip to Stage 3.

## Stage 1: Your Idea

### Not yet confident about your idea?

Great! Governance proposals potentially impact many stakeholders. Introduce your idea with known members of the community before investing resources into drafting a proposal. Don't let negative feedback dissuade you from exploring your idea if you think that it's still important.

If you know people who are very involved with Imversed, send them a private message with a concise overview of what you think will result from your idea or proposed changes. Wait for them to ask questions before providing details. Do the same in semi-private channels where people tend to be respectful (and hopefully supportive).

### Confident with your idea?

Great! However, remember that governance proposals potentially impact many stakeholders, which can happen in unexpected ways. Introduce your idea with members of the community before investing resources into drafting a proposal. At this point you should seek out and carefully consider critical feedback in order to protect yourself from confirmation [bias](https://en.wikipedia.org/wiki/Confirmation_bias). This is the ideal time to see a critical flaw, because submitting a flawed proposal will waste resources.

### Are you ready to draft a governance proposal?

There will likely be differences of opinion about the value of what you're proposing to do and the strategy by which you're planning to do it. If you've considered feedback from broad perspectives and think that what you're doing is valuable and that your strategy should work, and you believe that others feel this way as well, it's likely worth drafting a proposal. However, remember that the largest EVMOS stakers have the biggest vote, so a vocal minority isn't necessarily representative or predictive of the outcome of an on-chain vote.

A conservative approach is to have some confidence that you roughly have initial support from a majority of the voting power before proceeding to drafting your proposal. However, there are likely other approaches, and if your idea is important enough, you may want to pursue it regardless of whether or not you are confident that the voting power will support it.


## Stage 2: Your Draft Proposal

The next major section outlines and describes some potential elements of drafting a proposal. Ensure that you have considered your proposal and anticipated questions that the community will likely ask. Once your proposal is on-chain, you will not be able to change it.

### Proposal Elements

It will be important to balance two things: being detailed and being concise. You'll want to be concise so that people can assess your proposal quickly. You'll want to be detailed so that voters will have a clear, meaningful understanding of what the changes are and how they are likely to be impacted.

Every proposal should contain a summary with key details:

- who is submitting the proposal
- the amount of the proposal or parameter(s) being changed;
- and deliverables and timeline
- a reason for the proposal and potential impacts
- a short summary of the history (what compelled this proposal), solution that's being presented, and future expectations

Assume that many people will stop reading at this point. However, it is important to provide in-depth information, so a few more pointers for Parameter-Change, Community Spend, and ERC-20 Module proposals are below.


### Parameter-Change Proposal

1. Problem/Value - generally the problem or value that's motivating the parameter change(s)

2. Solution - generally how changing the parameter(s) will address the problem or improve the network

    - the beneficiaries of the change(s) (ie. who will these changes impact and how?)
        - voters should understand the importance of the change(s) in a simple way

3. Risks & Benefits - clearly describe how making this/these change(s) may expose stakeholders to new benefits and/or risks

4. Supplementary materials - optional materials eg. models, graphs, tables, research, signed petition, etc.


### Community Spend Proposal

1. **Applicant(s)** - the profile of the person(s)/entity making the proposal

- who you are and your involvement in Cosmos and/or other blockchain networks

- an overview of team members involved and their relevant experience

- brief mission statement for your organization/business (if applicable) eg. website

- past work you've done eg. include your Github

- some sort of proof of who you are eg. Keybase

2. **Problem** - generally what you're solving and/or opportunity you're addressing

- provide relevant information about both past and present issues created by this problem

- give suggestions as to the state of the future if this work is not completed

3. **Solution** - generally how you're proposing to deliver the solution

- your plan to fix the problem or deliver value

- the beneficiaries of this plan (ie. who will your plan impact and how?)

    - follow the "as a user" template ie. write a short user story about the problem you are trying to solve and how users will interact with what you're proposing to deliver (eg. benefits and functionality from a user’s perspective)
    
    - voters should understand the value of what you're providing in a simple way

- your reasons for selecting this plan

- your motivation for delivering this solution/value

5. Funding - amount and denomination proposed eg. 5000 Imversed

 - the entity controlling the account receiving the funding
 
 - consider an itemized breakdown of funding per major deliverable
 
 - consider outlining how the funds will be spent

5. Deliverables and timeline - the specifics of what you're delivering and how, and what to expect

    - what are the specific deliverables? (be detailed)

    - when will each of these be delivered?

    - will there be a date at which the project will be considered failed if the deliverables have not been met?

    - how will each of these be delivered?

    - what will happen if you do not deliver on time?

        - what is the deadline for the project to be considered failed?
 
        - do you have a plan to return the funds?

    - how will you be accountable to Imversed stakeholders?

        - how will you communicate updates and how often?

        - how can the community observe your progress?

        - how can the community provide feedback?

    - how should the quality of deliverables be assessed? eg. metrics

6. Relationships and disclosures

    - have you received or applied for grants or funding? for similar work? eg. from the Imversed Grants Program `(coming soon)`

    - how will you and/or your organization benefit?

    - do you see this work continuing in the future and is there a plan?

    - what are the risks involved with this work?

    - do you have conflicts of interest to declare?

## ERC-20 Proposal

1. **Applicant(s)** - the profile of the person(s)/entity making the proposal

    - who you are and your involvement in Cosmos and/or other blockchain networks

    - an overview of team members involved and their relevant experience

    - brief mission statement for your organization/business (if applicable) eg. website
    
    - past work you've done eg. include your Github
    
    - some sort of proof of who you are eg. Keybase

2. **Background information** - promote understanding of the ERC-20 Module

    - a mention of the original blog post `(coming soon)` that introduced the ERC-20 Module
    
    - a brief explanation of what the ERC-20 Module does
    
    - a mention of the [ERC-20 Module documentation](../modules/erc20/README.md)

3. **Solution** - generally how ERC-20 Module changes will be made

    - a brief explanation of what the proposal will do if it passes

    - a brief explanation of the precautions taken, how it was tested, and who was consulted prior to making the proposal

    - a breakdown of the proposal's payload, and third-party review

    - a brief explanation of the risks involved (depending on the direction of IBC Coin, ERC-20)

    - ensure the following are both adhered to and documented:

        - the contracts are verified ( via [Sourcify](https://sourcify.dev/))

        - the contracts are deployed open-source

        - the contracts do not extend the `IERC20.sol` interface through a malicious implementation

        - the contracts use the main libraries for ERC-20s (eg. [OpenZeppelin](https://docs.openzeppelin.com/contracts/4.x/erc20), [dapp.tools](https://dapp.tools/))

        - the transfer logic is not modified (i.e. transfer logic is not directly manipulated)

        - no malicious `Approve` events can directly manipulate users' balance through a delayed granted allowance


    Remember to provide links to the relevant Commonwealth Imversed community `(coming soon)` discussions concerning your proposal, as well as the proposal on testnet.
    
## Begin with a well-considered draft proposal

The ideal format for a proposal is as a Markdown file (ie. `.md`) in a Github repo or [HackMd](https://hackmd.io/). 
Markdown is a simple and accessible format for writing plain text files that is easy to learn. 
See the [Github Markdown Guide](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax) for details on writing markdown files.

## Engage the community with your draft proposal

1. Post a discussion in the Commonwealth Imversed community `(Coming soon)`. Ideally this should contain a link to this repository, either directly to your proposal if it has been merged, or else to a pull-request containing your proposal if it has not been merged yet.

2. Directly engage key members of the community for feedback. These could be large contributors, those likely to be most impacted by the proposal, and entities with high stake-backing (eg. high-ranked validators; large stakers).

3. Target members of the community in a semi-public way before bringing the draft to a full public audience. The burden of public scrutiny in a semi-anonymized environment (eg. Twitter) can be stressful and overwhelming without establishing support. Solicit opinions in places with people who have established reputations first.


## Submit your proposal to the testnet

:::tip

**Note:** Not sure how to submit a proposal to either testnet or mainnet? Check out this document `(Coming soon)`.

:::    


You may want to submit your proposal to the testnet chain before the mainnet for a number of reasons, such as wanting to see what the proposal description will look like, to share what the proposal will look like in advance with stakeholders, and to signal that your proposal is about to go live on the mainnet.

Perhaps most importantly, for parameter change proposals, you can test the parameter changes in advance (if you have enough support from the voting power on the testnet).

Submitting your proposal to the testnet increases the likelihood of engagement and the possibility that you will be alerted to a flaw before deploying your proposal to mainnet.


## Stage 3: Your On-Chain Proposal

A majority of the voting community should probably be aware of the proposal and have considered it before the proposal goes live on-chain. If you're taking a conservative approach, you should have reasonable confidence that your proposal will pass before risking deposit contributions. Make revisions to your draft proposal after each stage of engagement.

See the submitting guide for more on submitting proposals.


## The Deposit Period

The deposit period currently lasts 14 days. If you submitted your transaction with the minimum deposit (64 Imversed), your proposal will immediately enter the voting period. If you didn't submit the minimum deposit amount (currently 64 Imversed), then this may be an opportunity for others to show their support by contributing (and risking) their Imversed as a bond for your proposal. You can request contributions openly and also contact stakeholders directly (particularly stakeholders who are enthusiastic about your proposal). Remember that each contributor is risking their funds, and you can read more about the conditions for burning deposits.

This is a stage where proposals may begin to get broader attention. Most popular explorers currently display proposals that are in the deposit period, but due to proposal spamming, this may change.

A large cross-section of the blockchain/cryptocurrency community exists on Twitter. Having your proposal in the deposit period is a good time to engage the Imversed community to prepare validators to vote and Imversed-holders that are staking.

## The Voting Period

At this point you'll want to track which validator has voted and which has not. You'll want to re-engage directly with top stake-holders, ie. the highest-ranking validator operators, to ensure that:

- they are aware of your proposal;

- they can ask you any questions about your proposal; and

- they are prepared to vote.

Remember that any voter may change their vote at any time before the voting period ends. That historically doesn't happen often, but there may be an opportunity to convince a voter to change their vote. The biggest risk is that stakeholders won't vote at all (for a number of reasons). Validator operators tend to need multiple reminders to vote. How you choose to contact validator operators, how often, and what you say is up to you--remember that no validator is obligated to vote, and that operators are likely occupied by competing demands for their attention. Take care not to stress any potential relationship with validator operators.


## Submit a Proposal

If you have a final draft of your proposal ready to submit, you may want to push your proposal live on the testnet first. These are the three primary steps to getting your proposal live on-chain.

1. (**Optional**) Hosting supplementary materials for your proposal with IPFS (InterPlanetary File System)

2. Formatting the JSON file for the governance proposal transaction that will be on-chain

3. Sending the transaction that submits your governance proposal on-chain

## Hosting supplementary materials

In general we try to minimize the amount of data pushed to the blockchain. Hence, detailed documentation about a proposal is usually hosted on a separate censorship resistant data-hosting platform, like IPFS.

Once you have drafted your proposal, ideally as a Markdown file, you can upload it to the IPFS network:

1. either by [running an IPFS node and the IPFS software](https://ipfs.tech/), or

2. using a service such as [https://pinata.cloud](https://www.pinata.cloud/)

Ensure that you "pin" the file so that it continues to be available on the network. 
You should get a URL like this: `https://ipfs.io/ipfs/QmbkQNtCAdR1CNbFE8ujub2jcpwUcmSRpSCg8gVWrTHSWD` The value `QmbkQNtCAdR1CNbFE8ujub2jcpwUcmSRpSCg8gVWrTHSWD` is called the `CID` of your file - it is effectively the file's hash.

If you uploaded a markdown file, you can use the IPFS markdown viewer to render the document for better viewing. Links for the markdown viewer look like `https://ipfs.io/ipfs/QmTkzDwWqPbnAh5YiV5VwcTLnGdwSNsNTn2aDxdXBFca7D/example#/ipfs/<CID>`, where `<CID>` is your CID. For instance the link above would be: [https://ipfs.io/ipfs/QmTk...HSWD](https://ipfs.io/ipfs/QmTkzDwWqPbnAh5YiV5VwcTLnGdwSNsNTn2aDxdXBFca7D/example#/ipfs/QmbkQNtCAdR1CNbFE8ujub2jcpwUcmSRpSCg8gVWrTHSWD)

Share the URL with others and verify that your file is publicly accessible.

The reason we use IPFS is that it is a decentralized means of storage, making it resistant to censorship or single points of failure. This increases the likelihood that the file will remain available in the future.


## Formatting the JSON file for the governance proposal

Many proposals allow for long form text to be included, usually under the key `description`. These provide the opportunity to include [markdown](https://www.markdownguide.org/) if formatted correctly as well as line breaks with `\n`. Beware, however, that if you are using the CLI to create a proposal, and setting `description` using a flag, the text will be [escaped](https://en.wikipedia.org/wiki/Escape_sequences_in_C) which may have undesired effects. If you're using markdown or line breaks it's recommended to put the proposal text into a json file and include that file as part of the CLI proposal, as opposed to individual fields in flags.

## Text Proposals

TextProposals are used by delegators to agree to a certain strategy, plan, commitment, future upgrade, or any other statement in the form of text. Aside from having a record of the proposal outcome on the Imversed chain, a text proposal has no direct effect on Imversed.

### Real Example

Proposal 1 `(coming soon)` was representative of one of four core network activities that users had to participate in to claim tokens from the Imversed Rektdrop.

```list
Example of code will be coming soon
```

## Community Pool Spend Proposals

For community pool spend proposals, there are five components:

1. **Title** - the distinguishing name of the proposal, typically the way the that explorers list proposals

2. **Description** - the body of the proposal that further describes what is being proposed and details surrounding the proposal

3. **Recipient** - the Imversed (bech32-based) address that will receive funding from the Community Pool

4. **Amount** - the amount of funding that the recipient will receive in atto-Imversed (`aimversed`)

5. **Deposit** - the amount that will be contributed to the deposit (in `aimversed`) from the account submitting the proposal


### Made-Up Example

In this simple example (below), a network explorer will list the governance proposal as a `CommunityPoolSpendProposal`. When an observer selects the proposal, they'll see the description. Not all explorers will show the recipient and amount, so ensure that you verify that the description aligns with the what the governance proposal is programmed to enact. If the description says that a certain address will receive a certain number of Imversed, it should also be programmed to do that, but it's possible that that's not the case (accidentally or otherwise).

The `amount` is `1000000000000000000imversed`. This is equal to 1 Imversed, so `recipient` address `will be coming soon` will receive 1 Imversed if this proposal is passed.

The deposit of `64000000000000000000imversed` results in 64 Imversed being used from the proposal submitter's account. There is a minimum deposit required for a proposal to enter the voting period, and anyone may contribute to this deposit within a 5-day period. If the minimum deposit isn't reached before this time, the deposit amounts will be burned. Deposit amounts will also be burned if quorum isn't met in the vote or if the proposal is vetoed.

```list
Example of code will be coming soon
```

### Real Example

This is a governance protocol which [Flux Protocol](https://www.fluxprotocol.org/), the provider of a cross-chain oracle which provides smart contracts with access to economically secure data feeds, submitted to cover costs of the subsidizied FPO (First Party Oracle) solution which they deployed on the Evmos mainnet.

Users can query the proposal details with the Imversed command-line interface using this command:

```list
Example of code will be coming soon
```

## Params-Change Proposals

Changes to the `gov` module are different from the other kinds of parameter changes because `gov` has subkeys, as discussed [here](https://github.com/cosmos/cosmos-sdk/issues/5800). Only the key part of the JSON file is different for gov parameter-change proposals.

For parameter-change proposals, there are seven components:

1. **Title** - the distinguishing name of the proposal, typically the way the that explorers list proposals

2. **Description** - the body of the proposal that further describes what is being proposed and details surrounding the proposal

3. **Subspace** - the Imversed module with the parameter that is being changed

4. **Key** - the parameter that will be changed

5. **Value** - the value of the parameter that will be changed by the governance mechanism

6. **Denom** - `aimversed` (atto-Imversed) will be the type of asset used as the deposit

7. **Amount** - the amount that will be contributed to the deposit (in `aimversed`) from the account submitting the proposal

### Real Example

In the example below, a network explorer listed the governance proposal by its title: "Increase the minimum deposit for governance proposals." When a user selects the proposal, they'll see the proposal’s description. This proposal can be found on the Imversed network here `(link will be coming soon)`.

Not all explorers will show the proposed parameter changes that are coded into the proposal, so the delegator should verify that the description aligns with what the governance proposal is programmed to enact. If the description says that a certain parameter will be increased, it should also be programmed to do that, but it's possible that that's not the case (accidentally or otherwise).

Users can query the proposal details with the imversedd command-line interface using this command:

```list
Example of command will be coming soon
```

The deposit denom is aimversed and amount is 20100000000000000000. Therefore, a deposit of 20.1 Imversed will be included with this proposal. At the time, the Imversed mainnet had a 10 Imversed minimum deposit, so this proposal was put directly into the voting period (and subsequently passed). There is a minimum deposit required for a proposal to enter the voting period, and anyone may contribute to this deposit within a 5-day period. If the minimum deposit isn't reached before this time, the deposit amounts will be burned.

## Sending the transaction that submits your governance proposal

For information on how to use imversedd binary to submit an on-chain proposal through the governance module, please refer to the quickstart documentation.

### CLI

This is the command format for using imversedd (the command-line interface) to submit your proposal on-chain:\

```list
Example of command will be coming soon
```
:::tip
Use the `imversed tx gov --help` flag to get more info about the governance commands
:::


1. `imversedd` is the command-line interface client that is used to send transactions and query Imversed

2. `tx gov submit-proposal param-change` indicates that the transaction is submitting a parameter-change proposal

3. `--from mykey` is the account key that pays the transaction fee and deposit amount

4. `--gas 500000` is the maximum amount of gas permitted to be used to process the transaction

    - the more content there is in the description of your proposal, the more gas your transaction will consume

    - if this number isn't high enough and there isn't enough gas to process your transaction, the transaction will fail

    - the transaction will only use the amount of gas needed to process the transaction


5. `--gas-prices` is the flat-rate per unit of gas value for a validator to process your transaction

6. `--chain-id evmos_90001-2` is Imversed Mainnet. For current and past chain-id's, please look at the Chain ID documentation.

    - the testnet chain ID is `coming soon` (https://testnet.mintscan.io/imversed). For current and past testnet information, please look at the testnet repository(link will be coming soon)

7. `--node` is using a full node to send the transaction to the Imversed Mainnet


## Verifying your transaction

After posting your transaction, your command line interface (`imversedd`) will provide you with the transaction's hash, which you can either query using `imversedd` or by searching the transaction hash using Mintscan (https://www.mintscan.io/imversed) or any block explorer.

## Depositing funds after a proposal has been submitted

Sometimes a proposal is submitted without having the minimum token amount deposited yet. In these cases you would want to be able to deposit more tokens to get the proposal into the voting stage. In order to deposit tokens, you'll need to know what your proposal ID is after you've submitted your proposal. You can query all proposals by the following command:

```list
Example of command will be coming soon
```

If there are a lot of proposals on the chain already, you can also filter by your own address. For the proposal above, that would be:


```list
Example of command will be coming soon
```

Once you have the proposal ID, this is the command to deposit extra tokens:

```list
Example of command will be coming soon
```

In our case above, the `<proposal-id>` would be 59 as queried earlier. The `<deposit>` is written as 500000aimversed, just like the example above.

## Submit your proposal to the testnet

You may want to submit your proposal to the testnet chain before the mainnet for a number of reasons:

1. To see what the proposal description will look like

2. To signal that your proposal is about to go live on the mainnet

3. To share what the proposal will look like in advance with stakeholders

4. To test the functionality of the governance features

Submitting your proposal to the testnet increases the likelihood that you will discover a flaw before deploying your proposal on mainnet. A few things to keep in mind:

 - you'll need testnet tokens for your proposal (ask around for a faucet)

 - the parameters for testnet proposals are different (eg. voting period timing, deposit amount, deposit denomination)

 - the deposit denomination is in `atimversed` instead of `aimversed`

 # Community Pool

Imversed token-holders can cast a vote to approve spending from the Community Pool to fund development and projects in the Imversed ecosystem.

## Why create a proposal to use Community Pool funds?

There are other funding options, most notably the [Imversed Grants Program](https://medium.com/imversed). Why create a community-spend proposal?

- **As a strategy: you can do both**. You can submit your proposal to the Imversed Grants Program, but also consider submitting your proposal publicly on-chain. If the Imversed community votes in favor, you can withdraw your application.

- **As a strategy: funding is fast**. Besides the time it takes to push your proposal on-chain, the only other limiting factor is a fixed 5-day voting period. As soon as the proposal passes, your account will be credited the full amount of your proposal request.

- **To build rapport**. Engaging publicly with the community is the opportunity to develop relationships with stakeholders and to educate them about the importance of your work. Unforeseen partnerships could arise, and overall the community may value your work more if they are involved as stakeholders.

- **To be more independent**. The [Imversed Grants Program](https://medium.com/imversed) may not always be able to fund work. Having a more consistently funded source and having a report with its stakeholders means you can use your rapport to have confidence in your ability to secure funding without having to be dependent upon the foundation alone.

## FAQ

### How is the Community Pool funded?

10% of all tokens generated (via block rewards) are continually transferred to and accrue within the Community Pool.


## How can funding for the Community Pool change?

Though the rate of funding is currently fixed at 10% of tokens minted per epoch. The current value of funding may be modified with a governance proposal and enacted immediately after the proposal passes.

Funded projects that fail to deliver may return funding to Community Pool and entities may help fund the Community Pool by depositing funds directly to the escrow account.

## What is the balance of the Community Pool?

Community Pool Account: `coming soon`       

## How can funds from the Community Pool be spent?

Funds from the Imversed Community Pool may be spent via successful governance proposal.

## How are funds disbursed after a community-spend proposal is passed?

If a community-spend proposal passes successfully, the number of Imversed encoded in the proposal will be transferred from the community pool to the address encoded in the proposal, and this will happen immediately after the voting period ends.


# Chain Parameters

:::tip
Note: Working on a governance proposal related to the changing of chain parameters? Make sure to look at Imversed Governance, and specifically the best practices.
:::

If a parameter-change proposal is successful, the change takes effect immediately upon completion of the voting period.

## List of Parameters

For a comprehensive list of available module parameters see the table below:

| Module               | Codebase                  | Parameters
| -------------------- | ------------------------- | -----------------
|       `auth`         | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/auth/06_params.html)
|       `bank`         | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/bank/05_params.html)
|       `crisis`       | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/crisis/04_params.html)
|       `distribution` | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/distribution/06_events.html)
|       `governance`   | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/gov/06_params.html)
|       `slashing`     | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/slashing/08_params.html)
|       `staking`      | `cosmos-sdk`              | [reference](https://docs.cosmos.network/main/modules/staking/08_params.html)
|       `transfer`     | `ibc-go`                  | [reference](https://github.com/cosmos/ibc-go/blob/main/docs/ibc/params.md)
|       `evm`          | `ethermint`               | [reference]()
|       `feemarket`    | `ethermint`               | [reference]()
|       `claims`       | `imversed`                | [reference]()
|       `erc20`        | `imversed`                | [reference]()
|       `feesplit`     | `imversed`                | [reference]()
|       `incentives`   | `imversed`                | [reference]()
|       `inflation`    | `imversed`                | [reference]()
