package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/setavenger/go-bitcoinkernel/pkg/bitcoinkernel"
)

var (
	datadir string
	height  int
)

func init() {
	flag.StringVar(&datadir, "datadir", "", "Data directory of Bitcoin Core")
	flag.IntVar(&height, "height", 199980, "Height of the block to show")
	flag.Parse()

	if datadir == "" {
		log.Fatalf("Data directory is required")
	}
}

func main() {
	// Create a new context
	ctx, err := bitcoinkernel.NewContext(bitcoinkernel.ChainTypeSignet)
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	defer ctx.Close()

	chainman, err := bitcoinkernel.NewChainstateManager(ctx, datadir)
	if err != nil {
		log.Fatalf("Failed to create chainstate manager: %v", err)
	}
	defer chainman.Close()

	var blockIndex *bitcoinkernel.BlockIndex
	if height == 0 {
		blockIndex, err = chainman.GetBlockIndexFromTip()
	} else {
		blockIndex, err = chainman.GetBlockIndexFromHeight(height)
	}
	if err != nil {
		log.Fatalf("Failed to get block index: %v", err)
	}
	defer blockIndex.Close()

	blockKernel, err := chainman.ReadBlockData(blockIndex)
	if err != nil {
		log.Fatalf("Failed to read block data: %v", err)
	}
	defer blockKernel.Close()
	hash := blockKernel.GetHash()
	hashBytes := hash.GetBytes()
	reverseBytes(hashBytes[:])

	fmt.Printf("Block hash: %x\n", hashBytes)
	blockData, err := blockKernel.GetData()
	if err != nil {
		log.Fatalf("Failed to get block data: %v", err)
	}

	var prevBlockHash [32]byte
	copy(prevBlockHash[:], blockData[4:36])
	fmt.Printf("Prev block hash: %x\n", reverseBytes(prevBlockHash[:]))
	// fmt.Printf("Block data: %x\n", blockData)

	block, err := btcutil.NewBlockFromBytes(blockData)
	if err != nil {
		log.Fatalf("Failed to create block from bytes: %v", err)
	}

	blockUndo, err := chainman.ReadUndoData(blockIndex)
	if err != nil {
		log.Fatalf("Failed to read block undo data: %v", err)
	}
	defer blockUndo.Close()

	txs := block.Transactions()
	for i, tx := range txs[1:] {
		if i != 1 {
			continue
		}
		fmt.Printf("Transaction: %v\n", tx.Hash())
		for j, input := range tx.MsgTx().TxIn {
			txUndo, err := blockUndo.GetPrevoutByIndex(uint64(i), uint64(j))
			if err != nil {
				log.Fatalf("Failed to get prevout by index (tx: %d, input: %d): %v", i, j, err)
			}

			scriptPubkeyPrevout, err := txUndo.GetScriptPubkey()
			if err != nil {
				log.Fatalf("Failed to get script pubkey (tx: %d, input: %d): %v", i, j, err)
			}
			txUndo.Close()

			scriptPubkeyPrevoutBytes, err := scriptPubkeyPrevout.GetData()
			if err != nil {
				log.Fatalf("Failed to get script pubkey data (tx: %d, input: %d): %v", i, j, err)
			}
			scriptPubkeyPrevout.Close()

			fmt.Printf("(%d) Input:\t\t %v\n", j, input.PreviousOutPoint)
			fmt.Printf("(%d) Script pubkey:\t %x\n", j, scriptPubkeyPrevoutBytes)
			fmt.Printf("(%d) Witness:\t\t %v\n", j, input.Witness)
		}
		for _, output := range tx.MsgTx().TxOut {
			fmt.Printf("Output: %x %v\n", output.PkScript, output.Value)
		}
	}

	// if len(txs) > 8 {
	// 	printTx(txs[8])
	// }
}

func reverseBytes(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func printTx(tx *btcutil.Tx) {
	fmt.Printf("Transaction: %v\n", tx.Hash())
	for _, input := range tx.MsgTx().TxIn {
		fmt.Printf("Input: %v\n", input.PreviousOutPoint)
	}
	for _, output := range tx.MsgTx().TxOut {
		fmt.Printf("Output: %x %v\n", output.PkScript, output.Value)
	}
}
