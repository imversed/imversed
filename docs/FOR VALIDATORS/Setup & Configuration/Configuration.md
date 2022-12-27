# Configuration

## Block Time

The timeout-commit value in the node config defines how long we wait after committing a block, before starting on the new height (this gives us a chance to receive some more pre-commits, even though we already have +2/3). The current default value is `"1s"`.

> Note: From v6, this is handled automatically by the server when initializing the node. Validators will need to ensure their local node configurations in order to speed up the network to ~2s block times.

```text
# In ~/.imversed/config/config.toml

#######################################################
###         Consensus Configuration Options         ###
#######################################################
[consensus]

### ... 

# How long we wait after committing a block, before starting on the new
# height (this gives us a chance to receive some more precommits, even
# though we already have +2/3).
timeout_commit = "1s"
```

## Peers

In `~/.imversed/config/config.toml` you can set your peers.

See the [Add persistent peers section](https://docs.imversed.com/validators/testnet.html#add-persistent-peers) in our docs for an automated method, but field should look something like a comma separated string of peers (do not copy this, just an example):

### Sharing your Peer
You can see and share your peer with the tendermint show-node-id command:

```text
imversed tendermint show-node-id
ac29d21d0a6885465048a4481d16c12f59b2e58b
```

* **Peer Format:** `node-id@ip:port`
* **Example:** `ac29d21d0a6885465048a4481d16c12f59b2e58b@143.198.224.124:26656`

### Healthy peers
If you are relying on just seed node and no persistent peers or a low amount of them, please increase the following params in the `config.toml`:

```text
# Maximum number of inbound peers
max_num_inbound_peers = 120

# Maximum number of outbound peers to connect to, excluding persistent peers
max_num_outbound_peers = 60
```
