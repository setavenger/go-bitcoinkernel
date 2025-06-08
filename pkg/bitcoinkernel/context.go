package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"
import "errors"

type Context struct {
	ptr *C.kernel_Context
}

// NewContext creates a new kernel context (with default options)
func NewContext() (*Context, error) {
	opts := C.kernel_context_options_create()
	if opts == nil {
		return nil, errors.New("failed to create context options")
	}
	defer C.kernel_context_options_destroy(opts)

	// Create chain parameters for mainnet
	chainParams := C.kernel_chain_parameters_create(C.kernel_CHAIN_TYPE_SIGNET)
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
