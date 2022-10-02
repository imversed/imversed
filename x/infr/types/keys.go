package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "infr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// SmartContractKeyPrefix is the prefix to retrieve all Currency
	SmartContractKeyPrefix = "Smart_contract/address/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// InfrSmartContractKey returns the store key to retrieve a smart-contract from the index fields
func InfrSmartContractKey(
	address string,
) []byte {
	var key []byte

	denomBytes := []byte(address)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
