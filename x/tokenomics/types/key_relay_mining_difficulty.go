package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // RelayMiningDifficultyKeyPrefix is the prefix to retrieve all RelayMiningDifficulty
	RelayMiningDifficultyKeyPrefix = "RelayMiningDifficulty/value/"
)

// RelayMiningDifficultyKey returns the store key to retrieve a RelayMiningDifficulty from the index fields
func RelayMiningDifficultyKey(
serviceId string,
) []byte {
	var key []byte
    
    serviceIdBytes := []byte(serviceId)
    key = append(key, serviceIdBytes...)
    key = append(key, []byte("/")...)
    
	return key
}