package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/hyperxpizza/blockchain-example/blockchain"
)

type CLI struct {
	blockchain *blockchain.Blockchain
}

func NewCLI(chain *blockchain.Blockchain) *CLI {
	return &CLI{
		blockchain: chain,
	}
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

		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", string(block.Data))
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Proof of work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("")

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (c *CLI) Run() {
	c.ValidateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	default:
		c.PrintUsage()
		runtime.Goexit()

	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}

		c.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}

}
