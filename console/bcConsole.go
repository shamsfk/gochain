package console

import (
	"fmt"

	"github.com/shamsfk/gosuchain/blockchain"
)

// RunConsole binds functions to JS vm? inits and runs the console
func RunConsole(bc *blockchain.Blockchain) {
	cons := NewConsole()

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

	consCh := make(chan string)
	go cons.Run(consCh)
	for v := range consCh {
		fmt.Print(v)
	}
}
