package chain

import "github.com/shufflingpixels/antelope-go/abi"

type TransactionStatus uint8

const (
	TransactionStatusExecuted TransactionStatus = iota ///< succeed, no error handler executed
	TransactionStatusSoftFail                          ///< objectively failed (not executed), error handler executed
	TransactionStatusHardFail                          ///< objectively failed and error handler objectively failed thus no state change
	TransactionStatusDelayed                           ///< transaction delayed
	TransactionStatusExpired                           ///< transaction expired
	TransactionStatusUnknown  = TransactionStatus(255)
)

func (txs TransactionStatus) String() string {
	switch txs {
	case TransactionStatusExecuted:
		return "executed"
	case TransactionStatusSoftFail:
		return "soft_fail"
	case TransactionStatusHardFail:
		return "hard_fail"
	case TransactionStatusDelayed:
		return "delayed"
	case TransactionStatusExpired:
		return "expired"
	default:
		return "unknown"
	}
}

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
