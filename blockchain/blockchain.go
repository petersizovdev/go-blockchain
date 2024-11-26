package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
)

const dbPath = "./tmp/blocks"

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lasthash"))
		if err != nil {
			log.Panic(err)
		}
		lastHash, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = txn.Set([]byte("lasthash"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	if err != nil {
		log.Panic(err)
	}
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			log.Panic(err)
		}
		encodedBlock, err := item.ValueCopy(nil)
		if err != nil {
			log.Panic(err)
		}
		block = Deserialize(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	iter.CurrentHash = block.PrevHash

	return block
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lasthash")); err == badger.ErrKeyNotFound {
			fmt.Println("Blockchain does not exist")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = txn.Set([]byte("lasthash"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lasthash"))
			if err != nil {
				log.Panic(err)
			}
			lastHash, err = item.ValueCopy(nil)
			return err
		}
	})
	if err != nil {
		log.Panic(err)
	}

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}