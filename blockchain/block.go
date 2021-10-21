package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Nonce: 0}
	pow := NewProof(&block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return &block
}

func (b *Block) Serialize() ([]byte, error) {
	var resp bytes.Buffer
	encoder := gob.NewEncoder(&resp)
	err := encoder.Encode(b)
	if err != nil {
		log.Printf("encoder.Encode failed: %s\n", err.Error())
		return nil, err
	}

	return resp.Bytes(), nil
}

func Deserialize(data []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Printf("decoder.Decode failed: %s\n", err.Error())
		return nil, err
	}

	return &block, nil
}
