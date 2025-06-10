package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/setavenger/go-bitcoinkernel/pkg/bitcoinkernel"
)

var datadir string

func init() {
	flag.StringVar(&datadir, "datadir", "", "Data directory of Bitcoin Core")
	flag.Parse()

	if datadir == "" {
		log.Fatalf("Data directory is required")
	}
}

func main() {
	// Create a new context
	ctx, err := bitcoinkernel.NewContext()
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	defer ctx.Close()

	chainman, err := bitcoinkernel.NewChainstateManager(ctx, datadir)
	if err != nil {
		log.Fatalf("Failed to create chainstate manager: %v", err)
	}
	defer chainman.Close()

	// Get the block index at the chain tip
	blockIndex, err := chainman.GetBlockIndexFromTip()
	if err != nil {
		log.Fatalf("Failed to get block index from tip: %v", err)
	}
	defer blockIndex.Close()

	fmt.Println("Successfully got chain tip block index pointer:", blockIndex)
	// Get block height
	height := blockIndex.GetHeight()
	hash := blockIndex.GetBlockHash()

	fmt.Printf("Chain tip at height %d\n", height)
	fmt.Printf("Block hash: %x\n", hash.GetBytes())
}
