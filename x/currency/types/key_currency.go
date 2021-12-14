package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CurrencyKeyPrefix is the prefix to retrieve all Currency
	CurrencyKeyPrefix = "Currency/value/"
)

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
