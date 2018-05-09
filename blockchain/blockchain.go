package blockchain

import (
	"strings"
)

// Blockchain holds a slice of blocks
type Blockchain struct {
	blocks  []Block
	Version int
}

// ToString returns string representation of a Blockchain
func (bc Blockchain) ToString() string {
	builder := strings.Builder{}
	for _, block := range bc.blocks {
		builder.WriteRune('\n')
		builder.WriteString(block.ToString())
		builder.WriteRune('\n')
	}
	return builder.String()
}

// AddBlock adds specified Block to the chain
func (bc Blockchain) AddBlock(data *BlockData) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastBlock.Hash, bc.Version)
	newBlock.Index = len(bc.blocks)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a blockchain with genesis block
func NewBlockchain(version int) Blockchain {
	return Blockchain{
		blocks: []Block{
			NewBlock(&BlockData{"Genesis Block"}, []byte{0}, version),
		},
		Version: version,
	}
}

// ValidateBlockchain checks validity of every block in a given chain
func ValidateBlockchain(bc Blockchain) bool {
	if len(bc.blocks) < 1 {
		return false
	}
	if len(bc.blocks) == 1 {
		return true
	}

	for i := 1; i < len(bc.blocks); i++ {
		if !ValidateBlock(bc.blocks[i], bc.blocks[i-1]) {
			return false
		}
	}

	return true
}
