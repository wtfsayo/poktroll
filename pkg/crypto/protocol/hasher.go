package protocol

import "crypto/sha256"

const (
	RelayHasherSize      = sha256.Size
	TrieHasherSize       = sha256.Size
	TrieRootSize         = TrieHasherSize + trieRootMetadataSize
	trieRootMetadataSize = 16 // TODO_CONSIDERATION: Export this from the SMT package.
)

var (
	NewRelayHasher = sha256.New
	NewTrieHasher  = sha256.New
)
