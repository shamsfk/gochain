package main

import (
	"fmt"

	"./blockchain"
	"./console"
)

func main() {
	bc := blockchain.NewBlockchain(1)

	cons := console.NewConsole()

	cons.RegisterFunction("addBlock", func(dataStr string) {
		data := blockchain.BlockData{Data: dataStr}
		bc.AddBlock(&data)
	})

	cons.RegisterFunction("print", func() {
		fmt.Println(bc)
	})

	cons.RegisterFunction("validate", func() string {
		if blockchain.ValidateBlockchain(bc) {
			return "Blockchain is valid"
		}
		return "Blockchain is corrupted"
	})

	cons.Run()
}
