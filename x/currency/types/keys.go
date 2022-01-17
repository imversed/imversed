package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "currency"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_currency"

	// CurrencyKeyPrefix is the prefix to retrieve all Currency
	CurrencyKeyPrefix = "Currency/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// CurrencyKey returns the store key to retrieve a Currency from the index fields
func CurrencyKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
