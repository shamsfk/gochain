package main

import (
	"fmt"

	bc "./blockchain"
	cons "./console"
)

func main() {
	b := bc.NewBlockchain()

	data := &bc.BlockData{Data: "Send 1 BTC to Ivan"}
	b.AddBlock(data)

	data = &bc.BlockData{Data: "Send 2 more BTC to Ivan"}
	b.AddBlock(data)

	console := cons.NewConsole()

	console.RegisterFunction("printBlocks", func() {
		for _, block := range b.GetBlocks() {
			fmt.Println()
			fmt.Println(block.ToString())
			fmt.Println()
		}
	})

	console.Run()
}
