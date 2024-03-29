# Pool parameters

This details the types of messages supported by the x/pools module.

At launch, Pools have the following parameters:

```list
SwapFee
ExitFee
Weights
SmoothWeightChangeParams
```

We go through these in sequence.

1. SwapFee
    The swap fee is the cut of all swaps that goes to the Liquidity Providers (LPs) for a pool. Suppose a pool has a swap fee `s`. Then if a user wants to swap T tokens in the pool, `sT` tokens go to the LP's, and then `(1 - s)T` tokens are swapped according to the AMM swap function.
2. ExitFee
    The exit fee is a fee that is applied to LP's that want to remove their liquidity from the pool. Suppose a pool has an exit fee `e`. If they currently have `S` LP shares, then when they remove their liquidity they get tokens worth `(1 - e)S` shares back. The remaining `eS` shares are then burned, and the tokens corresponding to these shares are kept as liquidity.
3. Weights
    This defines the weights of the pool. https://balancer.fi/whitepaper.pdf
    TODO Add better description of how the weights affect things here.
4. SmoothWeightChangeParams
    SmoothWeightChangeParams allows pool governance to smoothly change the weights of the assets it holds in the pool. So it can slowly move from a 2:1 ratio, to a 1:1 ratio. Currently, smooth weight changes are implemented as a linear change in weight ratios over a given duration of time. So weights changed from 4:1 to 2:2 over 2 days, then at day 1 of the change, the weights would be 3:1.5, and at day 2 its 2:2, and will remain at these weight ratios.
