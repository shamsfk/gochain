package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

// BlockData holds Block's data
type BlockData struct {
	Data string
}

// ToBytes retuns []byte representation of data
func (bd *BlockData) ToBytes() []byte {
	return []byte(bd.Data)
}

// Block holds each individual block info
type Block struct {
	Index     int
	Timestamp int64
	Data      *BlockData
	PrevHash  []byte
	Hash      []byte
}

// CalculateHash calculates hash of the Block
func (b *Block) CalculateHash() []byte {
	timestamp := new(bytes.Buffer)
	err := binary.Write(timestamp, binary.LittleEndian, b.Timestamp)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	headers := bytes.Join([][]byte{b.PrevHash, b.Data.ToBytes(), timestamp.Bytes()}, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}

// NewBlock creates and returns a new Block
func NewBlock(data *BlockData, prevHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.CalculateHash()
	return block
}

// ValidateBlock checks is block is valid and does indeed follows prevBlock
func ValidateBlock(block *Block, prevBlock *Block) bool {
	if block.Index != prevBlock.Index+1 {
		return false
	}

	if string(block.PrevHash) != string(prevBlock.Hash) {
		return false
	}

	if string(block.CalculateHash()) != string(block.Hash) {
		return false
	}

	return true
}
