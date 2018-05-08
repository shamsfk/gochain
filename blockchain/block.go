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
	Timestamp int64
	Data      *BlockData
	PrevHash  []byte
	Hash      []byte
}

// SetHash calculates and sets Hash to the Block
func (b *Block) SetHash() {
	timestamp := new(bytes.Buffer)
	err := binary.Write(timestamp, binary.LittleEndian, b.Timestamp)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	headers := bytes.Join([][]byte{b.PrevHash, b.Data.ToBytes(), timestamp.Bytes()}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns a new Block
func NewBlock(data *BlockData, prevHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.SetHash()
	return block
}
