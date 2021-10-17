package main

import (
	"fmt"
	"strconv"

	"github.com/hyperxpizza/blockchain-example/blockchain"
)

func main() {
	b := blockchain.NewBlockchain()
	b.AddBlock("first block after genesis")
	b.AddBlock("second block after genesis")
	b.AddBlock("third block after genesis")

	for _, block := range b.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
