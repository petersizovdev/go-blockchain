package main

import (
	"fmt"
	"strconv"

	"github.com/petersizovdev/go-blockchain.git/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.Blocks {
		
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("BlockData: %s\n", block.Data)
		fmt.Printf("BlockHash: %x\n", block.Hash)

		pow:= blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
