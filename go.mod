module github.com/imversed/imversed

go 1.16

require (
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.45.2
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/holiman/uint256 v1.2.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/cors v1.8.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969
	github.com/stretchr/testify v1.7.1
	//github.com/tendermint/spm v0.1.5
	github.com/tendermint/tendermint v0.34.16
	github.com/tendermint/tm-db v0.6.7
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/genproto v0.0.0-20220401170504-314d38edb7de
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/tendermint/spm v0.1.9
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
