package main

import (
	"github.com/shamsfk/gosuchain/blockchain"
	console "github.com/shamsfk/gosuchain/console"
)

func main() {
	bc := blockchain.NewBlockchain(1)
	console.RunConsole(bc)
}
