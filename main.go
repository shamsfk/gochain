package main

import (
	"./blockchain"
)

func main() {
	bc := blockchain.NewBlockchain(1)

	runConsole(bc)
}
