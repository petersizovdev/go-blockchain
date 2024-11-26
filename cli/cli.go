package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/petersizovdev/go-blockchain.git/blockchain"
)

type CommandLine struct {
	Blockchain *blockchain.BlockChain
}

func NewCommandLine(blockchain *blockchain.BlockChain) *CommandLine {
	return &CommandLine{
		Blockchain: blockchain,
	}
}

func (cli *CommandLine) printUsage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(":::::: add -block <BLOCK_DATA> :::::: Add a block to the chain")
	fmt.Println(":::::: print :::::::::::::::::::::::: Prints the blocks in the chain")
	fmt.Println()
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Block added")
}

func (cli *CommandLine) printChain() {
	iter := cli.Blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("BlockData: %s\n", block.Data)
		fmt.Printf("BlockHash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printChain()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printChain()

	}
}
