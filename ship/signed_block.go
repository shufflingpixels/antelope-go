package ship

import (
	"bytes"

	"github.com/pnx/antelope-go/abi"
	"github.com/pnx/antelope-go/chain"
)

// NOTE: SignedBlock is abit special. The data is encoded twice.
// First the data in SignedBlock is encoded as usual.
// Then that data is encoded again as a byte array.
//
// and decoding is the reverse and because of that
// we create a new type SignedBlockBytes to take advantage of this.
// The data is first decoded to a SignedBlockBytes (just a plain byte array)
// and if client code wants to inspect the data, it can decode it to a SignedBlockBytes
// using the Unpack method.

type SignedBlock struct {
	SignedBlockHeader
	Transactions    []TransactionReceipt
	BlockExtensions []Extension
}

type SignedBlockHeader struct {
	chain.BlockHeader
	ProducerSignature chain.Signature // no pointer!!
}

type SignedBlockBytes []byte

// Unpack decodes the SignedBlockBytes into a SignedBlock
func (sbb *SignedBlockBytes) Unpack(sb *SignedBlock) error {
	return abi.NewDecoder(bytes.NewReader(*sbb), abi.DefaultDecoderFunc).Decode(sb)
}

func MakeSignedBlockBytes(sb *SignedBlock, sbb *SignedBlockBytes) error {
	b := new(bytes.Buffer)
	if err := abi.NewEncoder(b, abi.DefaultEncoderFunc).Encode(sb); err != nil {
		return err
	}
	*sbb = SignedBlockBytes(b.Bytes())
	return nil
}

func MustMakeSignedBlockBytes(sb *SignedBlock) *SignedBlockBytes {
	sbb := SignedBlockBytes{}
	if err := MakeSignedBlockBytes(sb, &sbb); err != nil {
		panic(err)
	}
	return &sbb
}
