package cli

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/hyperxpizza/blockchain-example/blockchain"
)

type CLI struct {
	blockchain *blockchain.Blockchain
}

func (*CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block <BLOCK_DATA> - add block to the chain")
	fmt.Println(" print - prints the block in the chain")
}

func (c *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		c.PrintUsage()
		runtime.Goexit()
	}
}

func (c *CLI) AddBlock(data string) {
	if err := c.blockchain.AddBlock(data); err != nil {
		log.Fatal(err)
	}
	log.Println("Added Block!")
}

func (c *CLI) printChain() {
	iterator := c.blockchain.Iterator()
	for {
		block, err := iterator.Next()
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Printf("Previous hash : %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", string(block.Data))
	}
}
