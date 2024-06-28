package chain_test

import (
	"testing"

	"github.com/shufflingpixels/antelope-go/chain"
	"github.com/shufflingpixels/antelope-go/internal/assert"
)

func TestBytes(t *testing.T) {
	bytes := chain.Bytes([]byte{0xbe, 0xef, 0xfa, 0xce})

	assert.ABICoding(t, bytes, []byte{0x04, 0xbe, 0xef, 0xfa, 0xce})
	assert.JSONCoding(t, bytes, `"beefface"`)
}
