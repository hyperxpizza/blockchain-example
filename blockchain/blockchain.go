package blockchain

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

func (chain *Blockchain) AddBlock(data string) error {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	newBlock := NewBlock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		serialized, err := newBlock.Serialize()
		if err != nil {
			return err
		}

		err = txn.Set(newBlock.Hash, serialized)
		if err != nil {
			return err
		}

		err = txn.Set([]byte("lh"), newBlock.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	chain.LastHash = newBlock.Hash

	return nil
}

func InitBlockChain(dbPath string) (*Blockchain, error) {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	if err != nil {
		log.Printf("badger.Open failed: %v\n", err)
		return nil, err
	}

	err = db.Update(func(txn *badger.Txn) error {
		//lh stands for lastHash
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			log.Println("No existing blockchain found")
			genesis := Genesis()
			log.Println("Genesis proved")

			serializedGenesis, err := genesis.Serialize()
			if err != nil {
				return err
			}

			err = txn.Set(genesis.Hash, serializedGenesis)
			if err != nil {
				return err
			}

			err = txn.Set([]byte("lh"), genesis.Hash)
			if err != nil {
				return err
			}
		} else {
			item, err := txn.Get([]byte("lh"))
			if err != nil {
				return err
			}

			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	blockchain := Blockchain{lastHash, db}
	return &blockchain, nil

}

type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (chain *Blockchain) Iterator() *BlockchainIterator {
	iterator := BlockchainIterator{chain.LastHash, chain.Database}
	return &iterator
}

func (iterator *BlockchainIterator) Next() (*Block, error) {
	var block *Block

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			block, err = Deserialize(val)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	iterator.CurrentHash = block.PrevHash
	return block, nil

}

func (b *Blockchain) Close() {
	b.Database.Close()
}
