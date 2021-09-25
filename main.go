package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{blocks: []*Block{Genesis()}}
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()
	return &block
}

func main() {
	b := NewBlockchain()
	b.AddBlock("first block after genesis")
	b.AddBlock("second block after genesis")
	b.AddBlock("third block after genesis")

	for _, block := range b.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}
