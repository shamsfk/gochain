package blockchain

import (
	"crypto/sha256"
)

// BlockData holds Block's data
type BlockData struct {
	Data string
}

func (bd *BlockData) getHash() []byte {
	hash := sha256.Sum256([]byte(bd.Data))
	return hash[:]
}

// ToString returns string representation of a BlockData
func (bd *BlockData) String() string {
	return bd.Data
}
