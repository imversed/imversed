# Snapshots & Archive Nodes

## List of Snapshots and Archives
Below is a list of publicly available snapshots that you can use to sync with the Imversed mainnet and archived [imversed_5555555-1 mainne](https://github.com/imversed/mainnet/tree/main/imversed_5555555-1):

### Snapshots
| Name | URL |
| --- | --- |
| `Staketab` | [github.com/staketab/nginx-cosmos-snap](https://github.com/staketab/nginx-cosmos-snap/blob/main/docs/imversed.md) |
| `Polkachu` | [polkachu.com](https://www.polkachu.com/tendermint_snapshots/imversed) |
| `Nodes Guru` | [snapshots.nodes.guru/imversed_5555555-1/](https://docs.imversed.com/validators/snapshots.nodes.guru/imversed_5555555-1/) |
| `Notional` | [mainnet/pruned/imversed_5555555-1(pebbledb)](https://snapshot.notional.ventures/imversed/)
[mainnet/archive/imversed_5555555-1(pebbledb)](https://snapshot.notional.ventures/imversed-archive/)
[testnet/archive/imversed_9000-4(pebbledb)](https://snapshot.notional.ventures/imversed-testnet-archive/) |

### Archives
| Name | URL |
| --- | --- |
| `Nodes Guru` | [snapshots.nodes.guru/imversed_5555555-1](https://snapshots.nodes.guru/imversed_5555555-1/) |
| `Polkachu` | [polkachu.com/tendermint_snapshots/imversed](https://www.polkachu.com/tendermint_snapshots/imversed) |
| `Forbole` | [bigdipper.live/imversed_5555555-1](https://s3.bigdipper.live.eu-central-1.linodeobjects.com/imversed_5555555-1.tar.lz4) |

To access snapshots and archives, follow the process below (this code snippet is to access a snapshot of the current network, imversed_5555555-1, from Nodes Guru):

```text
cd $HOME/.imversed/data
wget https://snapshots.nodes.guru/imversed_5555555-1/imversed_5555555-1-410819.tar
tar xf imversed_5555555-1-410819.tar
```

### PebbleDB
To use PebbleDB instead of GoLevelDB when using snapshots from Notional
Build:

```text
go mod edit -replace github.com/tendermint/tm-db=github.com/baabeetaa/tm-db@pebble
go mod tidy
go install -tags pebbledb -ldflags "-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb" ./...
```

Download snapshot:

```text
cd $HOME/.imversed/
URL_SNAPSHOT="https://snapshot.notional.ventures/imversed/data_20221024_193254.tar.gz"
wget -O - "$URL_SNAPSHOT" |tar -xzf -
```

Start:
`Set db_backend = "pebbledb" in config.toml or start with --db_backend=pebbledb`

```text
imversed start --db_backend=pebbledb
```

**Note:** use this [workaround](https://github.com/notional-labs/cosmosia/blob/main/docs/pebbledb.md) when upgrading a node running PebbleDB.