package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"
import (
	"unsafe"
)

// BlockHash represents a 32-byte block hash
type BlockHash struct {
	ptr *C.kernel_BlockHash
}

// Close frees the memory associated with the block hash
func (h *BlockHash) Close() {
	if h.ptr != nil {
		C.kernel_block_hash_destroy(h.ptr)
		h.ptr = nil
	}
}

// GetBytes returns the 32-byte hash as a byte slice
func (h *BlockHash) GetBytes() *[32]byte {
	if h.ptr == nil {
		return nil
	}
	// Convert the C array to a Go slice
	return (*[32]byte)(unsafe.Pointer(&h.ptr.hash[0]))
}
