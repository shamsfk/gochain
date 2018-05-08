package main

import (
	"fmt"

	bc "./blockchain"
)

func main() {
	b := bc.NewBlockchain()

	data := &bc.BlockData{Data: "Send 1 BTC to Ivan"}
	b.AddBlock(data)

	data = &bc.BlockData{Data: "Send 2 more BTC to Ivan"}
	b.AddBlock(data)

	for _, block := range b.GetBlocks() {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data.ToBytes())
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
