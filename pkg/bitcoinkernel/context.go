package bitcoinkernel

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lbitcoinkernel
// #include <stdlib.h>
// #include "bitcoinkernel.h"
import "C"
import "errors"

var (
	chainType = C.kernel_CHAIN_TYPE_MAINNET
)

/*
kernel_CHAIN_TYPE_MAINNET = 0,
kernel_CHAIN_TYPE_TESTNET,
kernel_CHAIN_TYPE_TESTNET_4,
kernel_CHAIN_TYPE_SIGNET,
kernel_CHAIN_TYPE_REGTEST,
*/

type ChainType int

const (
	ChainTypeMainnet ChainType = iota
	ChainTypeTestnet
	ChainTypeTestnet4
	ChainTypeSignet
	ChainTypeRegtest
)

func getChainType(chainType ChainType) C.kernel_ChainType {
	var cChainType C.kernel_ChainType
	switch chainType {
	case ChainTypeMainnet:
		cChainType = C.kernel_CHAIN_TYPE_MAINNET
	case ChainTypeTestnet:
		cChainType = C.kernel_CHAIN_TYPE_TESTNET
	case ChainTypeTestnet4:
		cChainType = C.kernel_CHAIN_TYPE_TESTNET_4
	case ChainTypeSignet:
		cChainType = C.kernel_CHAIN_TYPE_SIGNET
	case ChainTypeRegtest:
		cChainType = C.kernel_CHAIN_TYPE_REGTEST
	default:
		panic("invalid chain type")
	}
	return cChainType
}

type Context struct {
	ptr *C.kernel_Context
}

// NewContext creates a new kernel context (with default options)
func NewContext(chainType ChainType) (*Context, error) {
	cChainType := getChainType(chainType)
	opts := C.kernel_context_options_create()
	if opts == nil {
		return nil, errors.New("failed to create context options")
	}
	defer C.kernel_context_options_destroy(opts)

	// Create chain parameters for mainnet
	chainParams := C.kernel_chain_parameters_create(cChainType)
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
