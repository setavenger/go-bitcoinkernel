package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"

type BlockIndex struct {
	ptr *C.kernel_BlockIndex
}

func (b *BlockIndex) Close() {
	if b.ptr != nil {
		C.kernel_block_index_destroy(b.ptr)
		b.ptr = nil
	}
}

// GetHeight returns the height of the block
func (b *BlockIndex) GetHeight() uint32 {
	if b.ptr == nil {
		return 0
	}
	return uint32(C.kernel_block_index_get_height(b.ptr))
}

// // GetBlockHash returns the hash of the block in Little Endian format
func (b *BlockIndex) GetBlockHash() *BlockHash {
	if b.ptr == nil {
		return nil
	}
	hash := C.kernel_block_index_get_block_hash(b.ptr)
	if hash == nil {
		return nil
	}
	return &BlockHash{ptr: hash}
}

func (b *BlockIndex) Prev() *BlockIndex {
	if b.ptr == nil {
		return nil
	}
	prev := C.kernel_get_previous_block_index(b.ptr)
	if prev == nil {
		return nil
	}
	return &BlockIndex{ptr: prev}
}
