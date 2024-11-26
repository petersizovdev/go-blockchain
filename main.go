package main

import (
	"os"

	"github.com/petersizovdev/go-blockchain.git/blockchain"
	"github.com/petersizovdev/go-blockchain.git/cli"
)

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cl := cli.NewCommandLine(chain)
	cl.Run()
}
