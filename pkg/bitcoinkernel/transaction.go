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

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ptr *C.kernel_Transaction
}

// Close frees the memory associated with the transaction
func (t *Transaction) Close() {
	if t.ptr != nil {
		C.kernel_transaction_destroy(t.ptr)
		t.ptr = nil
	}
}

// CreateTransaction creates a new transaction from serialized data
func CreateTransaction(rawTx []byte) (*Transaction, error) {
	if len(rawTx) == 0 {
		return nil, errors.New("raw transaction data is empty")
	}
	tx := C.kernel_transaction_create((*C.uchar)(unsafe.Pointer(&rawTx[0])), C.size_t(len(rawTx)))
	if tx == nil {
		return nil, errors.New("failed to create transaction")
	}
	return &Transaction{ptr: tx}, nil
}

// VerifyScript verifies a script with the given parameters
func VerifyScript(scriptPubkey *ScriptPubkey, amount int64, tx *Transaction, spentOutputs []*TransactionOutput, inputIndex uint, flags uint) (bool, error) {
	if scriptPubkey == nil || tx == nil {
		return false, errors.New("script pubkey or transaction is nil")
	}

	var spentOutputsPtr **C.kernel_TransactionOutput
	if len(spentOutputs) > 0 {
		spentOutputsPtr = (**C.kernel_TransactionOutput)(unsafe.Pointer(&spentOutputs[0].ptr))
	}

	var status C.kernel_ScriptVerifyStatus
	result := C.kernel_verify_script(
		scriptPubkey.ptr,
		C.int64_t(amount),
		tx.ptr,
		spentOutputsPtr,
		C.size_t(len(spentOutputs)),
		C.uint(inputIndex),
		C.uint(flags),
		&status,
	)

	if status != C.kernel_SCRIPT_VERIFY_OK {
		return false, errors.New("script verification failed")
	}

	return bool(result), nil
}
