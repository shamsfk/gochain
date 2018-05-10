package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

// Block holds each individual block info
type Block struct {
	Version   int
	Index     int
	Timestamp int64
	Data      *BlockData
	PrevHash  []byte
	Hash      []byte
}

// NewBlock creates and returns a new Block
func NewBlock(data *BlockData, prevHash []byte, version int) *Block {
	block := Block{
		Version:   version,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.calculateHash()
	return &block
}

// ToString returns string representation of a block
func (b *Block) String() string {
	return fmt.Sprintf("Version:\t%v\nIndex:\t\t%v\nTimestamp:\t%v\nData:\t\t%s\nPrevHash:\t%x\nHash:\t\t%x",
		b.Version,
		b.Index,
		b.Timestamp,
		b.Data.String(),
		b.PrevHash,
		b.Hash,
	)
}

func (b *Block) calculateHash() []byte {
	timestamp := new(bytes.Buffer)
	err := binary.Write(timestamp, binary.LittleEndian, b.Timestamp)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	block := bytes.Join([][]byte{b.PrevHash, b.Data.getHash(), timestamp.Bytes()}, []byte{})
	hash := sha256.Sum256(block)

	return hash[:]
}

// ValidateBlock checks is block is valid and does indeed follows prevBlock
func ValidateBlock(block *Block, prevBlock *Block) bool {
	if block.Index != prevBlock.Index+1 {
		return false
	}

	if string(block.PrevHash) != string(prevBlock.Hash) {
		return false
	}

	if string(block.calculateHash()) != string(block.Hash) {
		return false
	}

	return true
}
