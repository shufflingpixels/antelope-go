package chain

import "github.com/pnx/antelope-go/abi"

type TransactionStatus uint8

const (
	TransactionStatusExecuted TransactionStatus = iota ///< succeed, no error handler executed
	TransactionStatusSoftFail                          ///< objectively failed (not executed), error handler executed
	TransactionStatusHardFail                          ///< objectively failed and error handler objectively failed thus no state change
	TransactionStatusDelayed                           ///< transaction delayed
	TransactionStatusExpired                           ///< transaction expired
	TransactionStatusUnknown  = TransactionStatus(255)
)

// abi.Marshaler conformance

func (txs TransactionStatus) MarshalABI(e *abi.Encoder) error {
	return e.WriteUint8(uint8(txs))
}

// abi.Unmarshaler conformance

func (txs *TransactionStatus) UnmarshalABI(d *abi.Decoder) error {
	v, err := d.ReadUint8()
	if err != nil {
		return err
	}
	*txs = TransactionStatus(v)
	return nil
}
