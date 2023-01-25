package types

import (
	"encoding/binary"
	"strings"
)

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
	prefixCreatorToVerses = "Xverse/creatorMapping"
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
	hash = strings.ToLower(hash)
	var key []byte

	hashBytes := []byte(hash)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}

// KeyPrefixCreatorToVerse prefix for mapping storage
func KeyPrefixCreatorToVerse(address string) []byte {
	key := []byte(prefixCreatorToVerses)
	key = append(key, []byte("/")...)
	key = append(key, []byte(address)...)
	return key
}

func OwnerKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)

	return key
}
