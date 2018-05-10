package main

import (
	"fmt"

	"github.com/shamsfk/gosuchain/blockchain"
	"github.com/shamsfk/gosuchain/console"
)

// InitConsole binds functions to JS vm, inits and runs the console
func InitConsole(bc *blockchain.Blockchain) *console.Console {
	cons := console.NewConsole()

	cons.RegisterFunction("_addBlock", func(dataStr string) {
		data := blockchain.BlockData{Data: dataStr}
		bc.AddBlock(&data)
	})

	cons.RegisterFunction("_print", func() {
		fmt.Println(bc)
	})

	cons.RegisterFunction("_validate", func() string {
		if blockchain.ValidateBlockchain(bc) {
			return "Blockchain is valid"
		}
		return "Blockchain is corrupted"
	})

	initStr := `
		var bc = {
			addBlock: _addBlock,
			print: _print,
			validate: _validate
		}
	`
	cons.RunJS(initStr)

	return &cons
}
