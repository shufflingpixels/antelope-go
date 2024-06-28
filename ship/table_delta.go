package ship

import (
	"bytes"
	"fmt"

	"github.com/shufflingpixels/antelope-go/abi"
)

type Row struct {
	Present bool
	Data    []byte
}

type TableDeltaV0 struct {
	Name string
	Rows []Row
}

type TableDelta struct {
	V0 *TableDeltaV0
}

type TableDeltaArray []byte

func (a *TableDeltaArray) Unpack(deltas *[]TableDelta) error {
	return abi.NewDecoder(bytes.NewReader(*a), abi.DefaultDecoderFunc).Decode(deltas)
}

func MakeTableDeltaArray(data []TableDelta, arr *TableDeltaArray) error {
	b := new(bytes.Buffer)
	if err := abi.NewEncoder(b, abi.DefaultEncoderFunc).Encode(data); err != nil {
		return err
	}
	*arr = TableDeltaArray(b.Bytes())
	return nil
}

func MustMakeTableDeltaArray(data []TableDelta) *TableDeltaArray {
	arr := TableDeltaArray{}
	if err := MakeTableDeltaArray(data, &arr); err != nil {
		panic(err)
	}
	return &arr
}

// abi.Marshaler conformance

func (t TableDelta) MarshalABI(e *abi.Encoder) error {
	if t.V0 != nil {
		if err := e.WriteByte(0x00); err != nil {
			return err
		}
		return e.Encode(*t.V0)
	}
	return nil
}

// abi.Unmarshaler conformance

func (a *TableDelta) UnmarshalABI(d *abi.Decoder) error {
	typ, err := d.ReadByte()
	if err != nil {
		return err
	}

	switch typ {
	case 0x0:
		a.V0 = &TableDeltaV0{}
		return d.Decode(a.V0)
	default:
		return fmt.Errorf("invalid variant type %d", typ)
	}
}
