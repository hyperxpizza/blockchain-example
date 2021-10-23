package main

import (
	"log"
	"os"

	"github.com/hyperxpizza/blockchain-example/blockchain"
	"github.com/hyperxpizza/blockchain-example/cli"
	"github.com/joho/godotenv"
)

func main() {

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("loading .env file failed: %v\n", err)
	}

	dbPath := os.Getenv("DB_PATH")

	defer os.Exit(0)
	chain, err := blockchain.InitBlockChain(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	defer chain.Close()

	cli := cli.NewCLI(chain)
	cli.Run()
}
