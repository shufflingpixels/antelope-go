package chain

import "github.com/shufflingpixels/antelope-go/abi"

type CompressionType uint8

const (
	CompressionNone = CompressionType(iota)
	CompressionZlib
)

func (c CompressionType) String() string {
	switch c {
	case CompressionNone:
		return "none"
	case CompressionZlib:
		return "zlib"
	default:
		return ""
	}
}

// abi.Marshaler conformance

func (b CompressionType) MarshalABI(e *abi.Encoder) error {
	return e.WriteByte(byte(b))
}

// abi.Unmarshaler conformance

func (b *CompressionType) UnmarshalABI(d *abi.Decoder) error {
	var err error
	if v, err := d.ReadByte(); err == nil {
		*b = CompressionType(v)
	}
	return err
}
