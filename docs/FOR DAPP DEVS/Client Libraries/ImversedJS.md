# ImversedJS

[`imversedjs`](https://github.com/imversed/js-client) Javascript and Typescript client library for [Imversed](https://imversed.com/).

## Tools
* **Address converter:** convert between `eth` and `imversed` addresses
* **Basic transaction generator:** used to create Imversed transactions
* **EIP-712 transaction generator:** used to create `EIP-712` signed Imversed transactions
* **Protobuf files:** used to create Cosmos Hub and Imversed transactions
* **REST provider:** used to query the Imversed REST API and interact with Imversed nodes

## Usage
Add package with a preferred package manager. For example:

```text
yarn add @imversed/js-client
```

Use queries for retrieving data from imversed:

```js
import { nft } from '@imversed/js-client'
import { NftBaseNFT } from '@imversed/js-client/lib/nft/rest'

const { queryClient } = nft

async function getNft(denomId: string, nftId: string): Promise<NftBaseNFT> {
    const q = await queryClient({ addr: 'https://query-endpoint-test.imversed.com'})

    const res = await q.queryNft(denomId, nftId)

    return res.data.nft
}
```

Use transactions (TXs) to put some data in Imversed:

```js
import { loadWallet, nft } from '@imversed/js-client'

const { txClient } = nft

const mnemonic = "proof fish fun burden differ screen miss vanish three report stereo bamboo purpose doll random blur prepare attack gallery lawn raven glove quantum blade"

async function mintNFT(denomId: string, nftId: string, name: string, uri: string, data: any) {
    const wallet = await loadWallet(mnemonic)
    const [account] = await wallet.getAccounts()
    const tx = await txClient(wallet, { addr: 'https://tx-endpoint-test.imversed.com'})

    const msg = tx.msgMintNFT({
        id: nftId,
        denomId,
        name,
        uri,
        data,
        sender: account.address,
        recipient: account.address
    })

    return tx.signAndBroadcast([msg], {
        fee: {
            amount: [{
                amount: '200',
                denom: 'nimv'
            }],
            gas: '200000'
        }
    })
}
```
