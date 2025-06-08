package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"
import (
	"errors"
	"path/filepath"
	"unsafe"
)

type ChainstateManager struct {
	ptr *C.kernel_ChainstateManager
	ctx *Context
}

// NewChainstateManager creates a new chainstate manager for a given datadir
func NewChainstateManager(ctx *Context, datadir string) (*ChainstateManager, error) {
	blocksDir := filepath.Join(datadir, "blocks")
	cDatadir := C.CString(datadir)
	defer C.free(unsafe.Pointer(cDatadir))
	cBlocksDir := C.CString(blocksDir)
	defer C.free(unsafe.Pointer(cBlocksDir))

	opts := C.kernel_chainstate_manager_options_create(
		ctx.ptr,
		cDatadir,
		C.size_t(len(datadir)),
		cBlocksDir,
		C.size_t(len(blocksDir)),
	)
	if opts == nil {
		return nil, errors.New("failed to create chainstate manager options")
	}
	defer C.kernel_chainstate_manager_options_destroy(opts)

	manager := C.kernel_chainstate_manager_create(ctx.ptr, opts)
	if manager == nil {
		return nil, errors.New("failed to create chainstate manager")
	}
	return &ChainstateManager{ptr: manager, ctx: ctx}, nil
}

// func (c *ChainstateManager) Close() {
// 	if c.ptr != nil {
// 		C.kernel_chainstate_manager_destroy(c.ptr, c.ctx.ptr)
// 		c.ptr = nil
// 	}
// }

// GetBlockIndexFromTip returns the tip block index
func (m *ChainstateManager) GetBlockIndexFromTip() (*BlockIndex, error) {
	idx := C.kernel_get_block_index_from_tip(m.ctx.ptr, m.ptr)
	if idx == nil {
		return nil, errors.New("failed to get block index from tip")
	}
	return &BlockIndex{ptr: idx}, nil
}

// GetBlockIndexFromHeight returns the block index at the specified height in the currently active chain
func (m *ChainstateManager) GetBlockIndexFromHeight(height int) (*BlockIndex, error) {
	if m.ptr == nil || m.ctx == nil {
		return nil, errors.New("chainstate manager or context is nil")
	}
	idx := C.kernel_get_block_index_from_height(m.ctx.ptr, m.ptr, C.int(height))
	if idx == nil {
		return nil, errors.New("failed to get block index from height")
	}
	return &BlockIndex{ptr: idx}, nil
}

func (m *ChainstateManager) Close() {
	if m.ptr != nil {
		C.kernel_chainstate_manager_destroy(m.ptr, m.ctx.ptr)
		m.ptr = nil
	}
}

func (m *ChainstateManager) GetNextBlockIndex(blockIndex *BlockIndex) *BlockIndex {
	if blockIndex == nil {
		return nil
	}
	next := C.kernel_get_next_block_index(m.ctx.ptr, m.ptr, blockIndex.ptr)
	if next == nil {
		return nil
	}
	return &BlockIndex{ptr: next}
}
