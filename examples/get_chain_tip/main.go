package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/setavenger/go-bitcoinkernel/pkg/bitcoinkernel"
)

var datadir string

func init() {
	flag.StringVar(&datadir, "datadir", filepath.Join(os.TempDir(), "bitcoinkernel"), "Data directory for Bitcoin Kernel")
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

	// Create a chain state manager
	if err := os.MkdirAll(datadir, 0755); err != nil {
		log.Fatalf("Failed to create datadir: %v", err)
	}

	chainman, err := ctx.NewChainstateManager(datadir)
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
}
