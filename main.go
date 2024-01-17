package main

import (
	"fmt"

	"github.com/GuolinY/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("first")
	chain.AddBlock("second")

	for i, block := range chain.Blocks {
		fmt.Printf("Idx %v, Data %s, Hash %x, Prev %x \n", i, block.Data, block.Hash, block.PrevHash)
	}
}
