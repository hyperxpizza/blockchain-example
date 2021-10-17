package blockchain

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{Genesis()}}
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

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
