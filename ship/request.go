package ship

import (
	"fmt"

	"github.com/shufflingpixels/antelope-go/abi"
)

// State History Plugin Requests

type GetStatusRequestV0 struct{}

type GetBlocksAckRequestV0 struct {
	NumMessages uint32
}

type GetBlocksRequestV0 struct {
	StartBlockNum       uint32
	EndBlockNum         uint32
	MaxMessagesInFlight uint32
	HavePositions       []*BlockPosition
	IrreversibleOnly    bool
	FetchBlock          bool
	FetchTraces         bool
	FetchDeltas         bool
}

type Request struct {
	StatusRequest    *GetStatusRequestV0
	BlocksRequest    *GetBlocksRequestV0
	BlocksAckRequest *GetBlocksAckRequestV0
}

// abi.Marshaler conformance

func (r Request) MarshalABI(e *abi.Encoder) error {
	if r.StatusRequest != nil {
		return e.WriteByte(0x00)
	}
	if r.BlocksRequest != nil {
		if err := e.WriteByte(0x01); err != nil {
			return err
		}
		return e.Encode(r.BlocksRequest)
	}
	if err := e.WriteByte(0x02); err != nil {
		return err
	}
	return e.Encode(r.BlocksAckRequest)
}

// abi.Unmarshaler conformance

func (r *Request) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		r.StatusRequest = &GetStatusRequestV0{}
		return nil
	case 0x1:
		r.BlocksRequest = &GetBlocksRequestV0{}
		return d.Decode(r.BlocksRequest)
	case 0x2:
		r.BlocksAckRequest = &GetBlocksAckRequestV0{}
		return d.Decode(r.BlocksAckRequest)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}
