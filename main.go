package main

import (
	"fmt"

	"./blockchain"
	"./console"
)

func main() {
	bc := blockchain.NewBlockchain()

	data := &blockchain.BlockData{Data: "Send 1 BTC to Ivan"}
	bc.AddBlock(data)

	data = &blockchain.BlockData{Data: "Send 2 more BTC to Ivan"}
	bc.AddBlock(data)

	cons := console.NewConsole()

	cons.RegisterFunction("print", func() {
		for _, block := range bc.GetBlocks() {
			fmt.Println()
			fmt.Println(block.ToString())
			fmt.Println()
		}
	})

	cons.Run()
}
