package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "verse"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName
)

// prefix bytes
const (
	prefixVerse = "Verse/verse"
)

// KVStore key prefixes
var (
	KeyPrefixVerse = []byte(prefixVerse)
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
