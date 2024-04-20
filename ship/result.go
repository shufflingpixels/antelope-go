package ship

import (
	"fmt"

	"github.com/pnx/antelope-go/abi"
)

// State History Plugin Results
type GetStatusResultV0 struct {
	Head                 *BlockPosition
	LastIrreversible     *BlockPosition
	TraceBeginBlock      uint32
	TraceEndBlock        uint32
	ChainStateBeginBlock uint32
	ChainStateEndBlock   uint32
}

type GetBlocksResultV0 struct {
	Head             BlockPosition
	LastIrreversible BlockPosition
	ThisBlock        *BlockPosition        `eosio:"optional"`
	PrevBlock        *BlockPosition        `eosio:"optional"`
	Block            *SignedBlockBytes     `eosio:"optional"`
	Traces           *TransactionTraceArray `eosio:"optional"`
	Deltas           *TableDeltaArray       `eosio:"optional"`
}

type Result struct {
	StatusResult *GetStatusResultV0
	BlocksResult *GetBlocksResultV0
}

// abi.Marshaler conformance

func (r Result) MarshalABI(e *abi.Encoder) error {
	if r.StatusResult != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*r.StatusResult)
	}
	if r.BlocksResult != nil {
		if err := e.WriteByte(0x01); err != nil {
			return err
		}
		return e.Encode(*r.BlocksResult)
	}
	return nil
}

// abi.Unmarshaler conformance

func (r *Result) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		r.StatusResult = &GetStatusResultV0{}
		return d.Decode(r.StatusResult)
	case 0x1:
		r.BlocksResult = &GetBlocksResultV0{}
		return d.Decode(r.BlocksResult)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}
