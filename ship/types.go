package ship

import (
	"github.com/pnx/antelope-go/chain"
)

// Miscellaneous structs that don't fit anywhere else.

type AccountDelta struct {
	Account chain.Name `json:"account"`
	Delta   int64      `json:"delta"`
}

type BlockPosition struct {
	BlockNum uint32
	BlockID  chain.Checksum256
}

type AccountAuthSequence struct {
	Account  chain.Name
	Sequence uint64
}

type Extension struct {
	Type uint16
	Data []byte
}
