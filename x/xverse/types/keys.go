package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "xverse"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName
)

// prefix bytes
const (
	prefixVerse           = "Xverse/xverse"
	prefixContractToVerse = "Xverse/contracts"
)

// KVStore key prefixes
var (
	KeyPrefixVerse    = []byte(prefixVerse)
	KeyPrefixContract = []byte(prefixContractToVerse)
)

// VerseKey returns the store key to retrieve a Verse from the index fields
func VerseKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ContractKey(
	hash string,
) []byte {
	var key []byte

	nameBytes := []byte(hash)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
