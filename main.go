package main

import (
	"github.com/shamsfk/gosuchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain(1)
	cons := InitConsole(bc)
	cons.Run()
}
