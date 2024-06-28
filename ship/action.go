package ship

import (
	"fmt"
	"unsafe"

	"github.com/shufflingpixels/antelope-go/abi"
	"github.com/shufflingpixels/antelope-go/chain"
)

type ActionTraceV0 struct {
	ActionOrdinal        uint
	CreatorActionOrdinal uint
	Receipt              *ActionReceipt `eosio:"optional"`
	Receiver             chain.Name
	Act                  chain.Action
	ContextFree          bool
	Elapsed              int64
	Console              string
	AccountRamDeltas     []AccountDelta
	Except               string `eosio:"optional"`
	ErrorCode            uint64 `eosio:"optional"`
}

type ActionTraceV1 struct {
	ActionOrdinal        uint
	CreatorActionOrdinal uint
	Receipt              *ActionReceipt `eosio:"optional"`
	Receiver             chain.Name
	Act                  chain.Action
	ContextFree          bool
	Elapsed              int64
	Console              string
	AccountRamDeltas     []AccountDelta
	Except               string `eosio:"optional"`
	ErrorCode            uint64 `eosio:"optional"`
	ReturnValue          []byte
}

type ActionTrace struct {
	V0 *ActionTraceV0
	V1 *ActionTraceV1
}

type ActionReceiptV0 struct {
	Receiver       chain.Name
	ActDigest      chain.Checksum256
	GlobalSequence uint64
	RecvSequence   uint64
	AuthSequence   []AccountAuthSequence
	CodeSequence   uint
	ABISequence    uint
}

type ActionReceipt struct {
	V0 *ActionReceiptV0
}

// abi.Marshaler conformance

func (at ActionTrace) MarshalABI(e *abi.Encoder) error {
	if at.V0 != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*at.V0)
	}
	if at.V1 != nil {
		if err := e.WriteByte(0x01); err != nil {
			return err
		}
		return e.Encode(*at.V1)
	}
	return nil
}

func (at ActionTraceV1) MarshalABI(e *abi.Encoder) error {
	// pointer magic here to reuse V0 encoding.
	if err := e.Encode(*(*ActionTraceV0)(unsafe.Pointer(&at))); err != nil {
		return err
	}

	if err := e.Encode(at.ReturnValue); err != nil {
		return err
	}

	return nil
}

func (ar ActionReceipt) MarshalABI(e *abi.Encoder) error {
	if ar.V0 != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*ar.V0)
	}
	return nil
}

// abi.Unmarshaler conformance

func (a *ActionTrace) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x00:
		a.V0 = &ActionTraceV0{}
		return d.Decode(a.V0)
	case 0x01:
		a.V1 = &ActionTraceV1{}
		return d.Decode(a.V1)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}

func (a *ActionReceipt) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		a.V0 = &ActionReceiptV0{}
		return d.Decode(a.V0)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}
