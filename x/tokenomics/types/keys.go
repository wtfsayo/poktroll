package types

import "bytes"

const (
	// ModuleName defines the module name
	ModuleName = "tokenomics"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenomics"
)

var (
	ParamsKey = []byte("p_tokenomics")

	// KeyDelimiter is the delimiter for composite keys.
	KeyDelimiter = []byte("/")
)

func KeyPrefix(p string) []byte { return []byte(p) }

// KeyComposite combines the given keys into a single key for use with KVStore.
func KeyComposite(keys ...[]byte) []byte {
	return bytes.Join(keys, KeyDelimiter)
}
