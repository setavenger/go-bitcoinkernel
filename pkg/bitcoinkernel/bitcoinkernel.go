package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Context struct {
	ptr *C.kernel_Context
}

type ChainstateManager struct {
	ptr *C.kernel_ChainstateManager
	ctx *Context
}

type BlockIndex struct {
	ptr *C.kernel_BlockIndex
}

// NewContext creates a new kernel context (with default options)
func NewContext() (*Context, error) {
	opts := C.kernel_context_options_create()
	if opts == nil {
		return nil, errors.New("failed to create context options")
	}
	defer C.kernel_context_options_destroy(opts)

	// Create chain parameters for mainnet
	chainParams := C.kernel_chain_parameters_create(C.kernel_CHAIN_TYPE_MAINNET)
	if chainParams == nil {
		return nil, errors.New("failed to create chain parameters")
	}
	defer C.kernel_chain_parameters_destroy(chainParams)

	// Set chain parameters in options
	C.kernel_context_options_set_chainparams(opts, chainParams)

	ctx := C.kernel_context_create(opts)
	if ctx == nil {
		return nil, errors.New("failed to create kernel context")
	}
	return &Context{ptr: ctx}, nil
}

func (c *Context) Close() {
	if c.ptr != nil {
		C.kernel_context_destroy(c.ptr)
		c.ptr = nil
	}
}

// NewChainstateManager creates a new chainstate manager for a given datadir
func (c *Context) NewChainstateManager(datadir string) (*ChainstateManager, error) {
	cDatadir := C.CString(datadir)
	defer C.free(unsafe.Pointer(cDatadir))

	opts := C.kernel_chainstate_manager_options_create(
		c.ptr,
		cDatadir,
		C.size_t(len(datadir)),
		nil, 0,
	)
	if opts == nil {
		return nil, errors.New("failed to create chainstate manager options")
	}
	defer C.kernel_chainstate_manager_options_destroy(opts)

	manager := C.kernel_chainstate_manager_create(c.ptr, opts)
	if manager == nil {
		return nil, errors.New("failed to create chainstate manager")
	}
	return &ChainstateManager{ptr: manager, ctx: c}, nil
}

func (m *ChainstateManager) Close() {
	if m.ptr != nil {
		C.kernel_chainstate_manager_destroy(m.ptr, m.ctx.ptr)
		m.ptr = nil
	}
}

// GetBlockIndexFromTip returns the tip block index
func (m *ChainstateManager) GetBlockIndexFromTip() (*BlockIndex, error) {
	idx := C.kernel_get_block_index_from_tip(m.ctx.ptr, m.ptr)
	if idx == nil {
		return nil, errors.New("failed to get block index from tip")
	}
	return &BlockIndex{ptr: idx}, nil
}

func (b *BlockIndex) Close() {
	if b.ptr != nil {
		C.kernel_block_index_destroy(b.ptr)
		b.ptr = nil
	}
}
