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

// Block represents a Bitcoin block
type Block struct {
	ptr *C.kernel_Block
}

// Close frees the memory associated with the block
func (b *Block) Close() {
	if b.ptr != nil {
		C.kernel_block_destroy(b.ptr)
		b.ptr = nil
	}
}

// GetHash returns the hash of the block
func (b *Block) GetHash() *BlockHash {
	if b.ptr == nil {
		return nil
	}
	hash := C.kernel_block_get_hash(b.ptr)
	if hash == nil {
		return nil
	}
	return &BlockHash{ptr: hash}
}

// GetData returns the serialized block data
func (b *Block) GetData() ([]byte, error) {
	if b.ptr == nil {
		return nil, errors.New("block is nil")
	}
	data := C.kernel_copy_block_data(b.ptr)
	if data == nil {
		return nil, errors.New("failed to copy block data")
	}
	defer C.kernel_byte_array_destroy(data)
	return C.GoBytes(unsafe.Pointer(data.data), C.int(data.size)), nil
}

// BlockUndo represents the undo data for a block
type BlockUndo struct {
	ptr *C.kernel_BlockUndo
}

// Close frees the memory associated with the block undo data
func (u *BlockUndo) Close() {
	if u.ptr != nil {
		C.kernel_block_undo_destroy(u.ptr)
		u.ptr = nil
	}
}

// GetTransactionUndoSize returns the number of previous transaction outputs in the transaction
func (u *BlockUndo) GetTransactionUndoSize(txIndex uint64) uint64 {
	if u.ptr == nil {
		return 0
	}
	return uint64(C.kernel_get_transaction_undo_size(u.ptr, C.uint64_t(txIndex)))
}

// GetPrevoutByIndex returns the previous output at the specified index
func (u *BlockUndo) GetPrevoutByIndex(txIndex, outputIndex uint64) (*TransactionOutput, error) {
	if u.ptr == nil {
		return nil, errors.New("block undo is nil")
	}
	output := C.kernel_get_undo_output_by_index(u.ptr, C.uint64_t(txIndex), C.uint64_t(outputIndex))
	if output == nil {
		return nil, errors.New("failed to get undo output")
	}
	return &TransactionOutput{ptr: output}, nil
}

// TransactionOutput represents a transaction output
type TransactionOutput struct {
	ptr *C.kernel_TransactionOutput
}

// Close frees the memory associated with the transaction output
func (o *TransactionOutput) Close() {
	if o.ptr != nil {
		C.kernel_transaction_output_destroy(o.ptr)
		o.ptr = nil
	}
}

// GetScriptPubkey returns the script pubkey of the output
func (o *TransactionOutput) GetScriptPubkey() (*ScriptPubkey, error) {
	if o.ptr == nil {
		return nil, errors.New("transaction output is nil")
	}
	script := C.kernel_copy_script_pubkey_from_output(o.ptr)
	if script == nil {
		return nil, errors.New("failed to get script pubkey")
	}
	return &ScriptPubkey{ptr: script}, nil
}

// GetAmount returns the amount of the output
func (o *TransactionOutput) GetAmount() uint64 {
	if o.ptr == nil {
		return 0
	}
	return uint64(C.kernel_get_transaction_output_amount(o.ptr))
}

// ScriptPubkey represents a script pubkey
type ScriptPubkey struct {
	ptr *C.kernel_ScriptPubkey
}

// Close frees the memory associated with the script pubkey
func (s *ScriptPubkey) Close() {
	if s.ptr != nil {
		C.kernel_script_pubkey_destroy(s.ptr)
		s.ptr = nil
	}
}

// GetData returns the serialized script pubkey data
func (s *ScriptPubkey) GetData() ([]byte, error) {
	if s.ptr == nil {
		return nil, errors.New("script pubkey is nil")
	}
	data := C.kernel_copy_script_pubkey_data(s.ptr)
	if data == nil {
		return nil, errors.New("failed to copy script pubkey data")
	}
	defer C.kernel_byte_array_destroy(data)
	return C.GoBytes(unsafe.Pointer(data.data), C.int(data.size)), nil
}
