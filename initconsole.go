package main

import (
	"fmt"

	"github.com/shamsfk/gosuchain/blockchain"
	"github.com/shamsfk/gosuchain/console"
)

// InitConsole binds functions to JS vm, inits and runs the console
func InitConsole(bc *blockchain.Blockchain) *console.Console {
	cons := console.NewConsole()

	cons.RegisterFunction("__addBlock", func(dataStr string) {
		data := blockchain.BlockData{Data: dataStr}
		bc.AddBlock(&data)
	})

	cons.RegisterFunction("__print", func() {
		fmt.Println(bc)
	})

	cons.RegisterFunction("__validate", func() string {
		if blockchain.ValidateBlockchain(bc) {
			return "Blockchain is valid"
		}
		return "Blockchain is corrupted"
	})

	initStr := `
		var bc = {
			addBlock: __addBlock,
			print: __print,
			validate: __validate
		}
	`
	cons.ExecuteJS(initStr)

	return &cons
}
