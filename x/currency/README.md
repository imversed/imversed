# Currency

This module provides messages to issue and mint user coins.

## Issue

Before mint user coins denom needs to be issued.

    imversed tx currency issue <user coins' denom> --from <user's address>

## Mint

To mint user coins it needs to broadcast transaction with currency mint message:

    imversed tx currency mint 1000000000mycoin --from <user's address>