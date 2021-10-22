package main

import (
	"log"
	"os"

	"github.com/hyperxpizza/blockchain-example/blockchain"
	"github.com/hyperxpizza/blockchain-example/cli"
)

func main() {
	defer os.Exit(0)
	chain, err := blockchain.InitBlockChain()
	if err != nil {
		log.Fatal(err)
	}

	defer chain.Close()

	cli := cli.NewCLI(chain)
	cli.Run()
}
