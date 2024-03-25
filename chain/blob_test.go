package chain_test

import (
	"encoding/json"
	"testing"

	"github.com/pnx/antelope-go/chain"
	"github.com/pnx/antelope-go/internal/assert"
)

func TestBlob(t *testing.T) {
	blob := chain.Blob([]byte{0xbe, 0xef, 0xfa, 0xce})
	assert.ABICoding(t, blob, []byte{0x04, 0xbe, 0xef, 0xfa, 0xce})
	assert.JSONCoding(t, blob, `"vu/6zg=="`)

	// make sure we can decode the blobs coming from nodeos
	var blob2 chain.Blob
	err := json.Unmarshal([]byte(`"vu/6zg="`), &blob2)
	assert.NoError(t, err)
	assert.Equal(t, blob, blob2)
	var blob3 chain.Blob
	err = json.Unmarshal([]byte(`"vu/6zg"`), &blob3)
	assert.NoError(t, err)
	assert.Equal(t, blob, blob3)
}
