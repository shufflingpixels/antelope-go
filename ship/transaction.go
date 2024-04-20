package ship

import (
	"bytes"
	"fmt"

	"github.com/pnx/antelope-go/abi"
	"github.com/pnx/antelope-go/chain"
)

type TransactionTraceV0 struct {
	ID              chain.Checksum256       `json:"id"`
	Status          chain.TransactionStatus `json:"status"`
	CPUUsageUS      uint32                  `json:"cpu_usage_us"`
	NetUsageWords   uint                    `json:"net_usage_words"`
	Elapsed         int64                   `json:"elapsed"`
	NetUsage        uint64                  `json:"net_usage"`
	Scheduled       bool                    `json:"scheduled"`
	ActionTraces    []*ActionTrace          `json:"action_traces"`
	AccountDelta    *AccountDelta           `json:"account_delta" eosio:"optional"`
	Except          string                  `json:"except" eosio:"optional"`
	ErrorCode       uint64                  `json:"error_code" eosio:"optional"`
	FailedDtrxTrace *TransactionTrace       `json:"failed_dtrx_trace" eosio:"optional"`
	Partial         *PartialTransaction     `json:"partial" eosio:"optional"`
}

type TransactionTrace struct {
	V0 *TransactionTraceV0
}

type TransactionTraceArray []byte

type Transaction struct {
	TxId   *chain.Checksum256
	Packed *chain.PackedTransaction
}

type TransactionReceipt struct {
	chain.TransactionReceiptHeader
	Trx Transaction
}

type PartialTransactionV0 struct {
	Expiration            uint32
	RefBlockNum           uint16
	RefBlockPrefix        uint32
	MaxNetUsageWords      uint
	MaxCpuUsageMs         uint8
	DelaySec              uint
	TransactionExtensions []Extension
	Signatures            []chain.Signature
	ContextFreeData       []byte
}

type PartialTransaction struct {
	V0 *PartialTransactionV0
}


func MakeTransactionTraceArray(data []TransactionTrace, arr *TransactionTraceArray) error {
	b := new(bytes.Buffer)
	if err := abi.NewEncoder(b, abi.DefaultEncoderFunc).Encode(data); err != nil {
		return err
	}
	*arr = TransactionTraceArray(b.Bytes())
	return nil
}

func MustMakeTransactionTraceArray(data []TransactionTrace) *TransactionTraceArray {
	arr := TransactionTraceArray{}
	if err := MakeTransactionTraceArray(data, &arr); err != nil {
		panic(err)
	}
	return &arr
}

func (a *TransactionTraceArray) Unpack(traces []TransactionTrace) error {
	return abi.NewDecoder(bytes.NewReader(*a), abi.DefaultDecoderFunc).Decode(traces)
}

// abi.Marshaler conformance

func (t Transaction) MarshalABI(e *abi.Encoder) error {
	if t.TxId != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*t.TxId)
	}
	if t.Packed != nil {
		if err := e.WriteByte(0x01); err != nil {
			return err
		}
		return e.Encode(*t.Packed)
	}
	return nil
}

func (t TransactionTrace) MarshalABI(e *abi.Encoder) error {
	if t.V0 != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*t.V0)
	}
	return nil
}

func (pt PartialTransaction) MarshalABI(e *abi.Encoder) error {
	if pt.V0 != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*pt.V0)
	}
	return nil
}

// abi.Unmarshaler conformance

func (tx *Transaction) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		tx.TxId = &chain.Checksum256{}
		return d.Decode(tx.TxId)
	case 0x1:
		tx.Packed = &chain.PackedTransaction{}
		return d.Decode(tx.Packed)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}

func (a *TransactionTrace) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		a.V0 = &TransactionTraceV0{}
		return d.Decode(a.V0)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}

func (a *PartialTransaction) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		a.V0 = &PartialTransactionV0{}
		return d.Decode(a.V0)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}
