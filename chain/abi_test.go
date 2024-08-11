package chain_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/shufflingpixels/antelope-go/abi"
	"github.com/shufflingpixels/antelope-go/chain"
	"github.com/shufflingpixels/antelope-go/internal/assert"
)

func loadAbi(v string) *chain.Abi {
	var rv chain.Abi
	err := json.Unmarshal([]byte(v), &rv)
	if err != nil {
		panic(err)
	}
	return &rv
}

var tokenAbi = loadAbi(`
{
    "version": "eosio::abi/1.1",
    "types": [],
    "structs": [
        {
            "name": "account",
            "base": "",
            "fields": [
                {
                    "name": "balance",
                    "type": "asset"
                }
            ]
        },
        {
            "name": "banana",
            "base": "",
            "fields": [
                {
                    "name": "moo",
                    "type": "name"
                }
            ]
        },
        {
            "name": "create",
            "base": "",
            "fields": [
                {
                    "name": "issuer",
                    "type": "name"
                },
                {
                    "name": "maximum_supply",
                    "type": "asset"
                }
            ]
        },
        {
            "name": "currency_stats",
            "base": "",
            "fields": [
                {
                    "name": "supply",
                    "type": "asset"
                },
                {
                    "name": "max_supply",
                    "type": "asset"
                },
                {
                    "name": "issuer",
                    "type": "name"
                }
            ]
        },
        {
            "name": "issue",
            "base": "",
            "fields": [
                {
                    "name": "to",
                    "type": "name"
                },
                {
                    "name": "quantity",
                    "type": "asset"
                },
                {
                    "name": "memo",
                    "type": "string"
                }
            ]
        },
        {
            "name": "open",
            "base": "",
            "fields": [
                {
                    "name": "owner",
                    "type": "name"
                },
                {
                    "name": "symbol",
                    "type": "symbol"
                },
                {
                    "name": "ram_payer",
                    "type": "name"
                }
            ]
        },
        {
            "name": "megatransfer",
            "base": "transfer",
            "fields": [
                {
                    "name": "extra",
                    "type": "mega"
                },
				        {
                    "name": "extra2",
                    "type": "banana[]"
                }
            ]
        },
        {
            "name": "transfer",
            "base": "",
            "fields": [
                {
                    "name": "from",
                    "type": "name"
                },
                {
                    "name": "to",
                    "type": "name"
                },
                {
                    "name": "quantity",
                    "type": "asset"
                },
                {
                    "name": "memo",
                    "type": "string"
                }
            ]
        },
		{
            "name": "noop",
            "base": "",
            "fields": []
        }
    ],
    "actions": [
        {
            "name": "close",
            "type": "close",
            "ricardian_contract": ""
        },
        {
            "name": "create",
            "type": "create",
            "ricardian_contract": ""
        },
        {
            "name": "issue",
            "type": "issue",
            "ricardian_contract": ""
        },
        {
            "name": "open",
            "type": "open",
            "ricardian_contract": ""
        },
        {
            "name": "retire",
            "type": "retire",
            "ricardian_contract": ""
        },
        {
            "name": "transfer",
            "type": "transfer",
            "ricardian_contract": ""
        },
        {
            "name": "bigtransfer",
            "type": "megatransfer",
            "ricardian_contract": ""
        },
		{
            "name": "noop",
            "type": "noop",
            "ricardian_contract": ""
        }
    ],
    "tables": [
        {
            "name": "accounts",
            "index_type": "i64",
            "key_names": [],
            "key_types": [],
            "type": "account"
        },
        {
            "name": "stat",
            "index_type": "i64",
            "key_names": [],
            "key_types": [],
            "type": "currency_stats"
        },
        {
            "name": "noop",
            "index_type": "i64",
            "key_names": [],
            "key_types": [],
            "type": "noop"
        }
    ],
    "ricardian_clauses": [],
    "variants": [
        {
            "name": "mega",
            "types": ["uint64", "string"]
        }
    ]
}
`)

var transferData = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x5d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xae, 0x39,
	0x10, 0x27, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x45, 0x4f, 0x53, 0x00, 0x00, 0x00, 0x00,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	// extra variant
	0x01,
	// utf8 string "foo"
	0x03, 0x66, 0x6f, 0x6f,
	// extra2 array
	0x01,                                           // 1 item
	0x00, 0x00, 0x00, 0x00, 0x00, 0xea, 0x30, 0x55, // name eosio
}

var statRow = []byte{
	0x00, 0x24, 0xca, 0x94, 0x27, 0x00, 0x00, 0x00,
	0x04, 0x45, 0x4f, 0x53, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x20, 0x9b, 0xf3, 0x5e, 0x10, 0x00, 0x00,
	0x04, 0x45, 0x4f, 0x53, 0x00, 0x00, 0x00, 0x00,
	0x50, 0xc8, 0x10, 0x81, 0x2d, 0x95, 0xd0, 0x31,
}

func TestAbiDecode(t *testing.T) {
	rv, err := tokenAbi.Decode(bytes.NewReader(transferData), "megatransfer")
	assert.NoError(t, err)
	assert.Equal(t, rv, map[string]interface{}{
		"from":     chain.N("foo"),
		"to":       chain.N("bar"),
		"quantity": *chain.A("1.0000 EOS"),
		"memo":     "hello",
		"extra":    []interface{}{"string", "foo"},
		"extra2": []interface{}{
			map[string]interface{}{"moo": chain.N("eosio")},
		},
	})
}

func TestAbiEncode(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	err := tokenAbi.Encode(buf, "megatransfer", map[string]interface{}{
		"from":     chain.N("foo"),
		"to":       chain.N("bar"),
		"quantity": *chain.A("1.0000 EOS"),
		"memo":     "hello",
		"extra":    []interface{}{"string", "foo"},
		"extra2": []interface{}{
			map[string]interface{}{"moo": chain.N("eosio")},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, buf.Bytes(), transferData)
}

func TestAbiEncodeAction(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	err := tokenAbi.EncodeAction(buf, chain.N("bigtransfer"), map[string]interface{}{
		"from":     chain.N("foo"),
		"to":       chain.N("bar"),
		"quantity": *chain.A("1.0000 EOS"),
		"memo":     "hello",
		"extra":    []interface{}{"string", "foo"},
		"extra2": []interface{}{
			map[string]interface{}{"moo": chain.N("eosio")},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, buf.Bytes(), transferData)
}

func TestAbiDecodeAction(t *testing.T) {
	rv, err := tokenAbi.DecodeAction(bytes.NewReader(transferData), chain.N("bigtransfer"))
	assert.NoError(t, err)
	assert.Equal(t, rv, map[string]interface{}{
		"from":     chain.N("foo"),
		"to":       chain.N("bar"),
		"quantity": *chain.A("1.0000 EOS"),
		"memo":     "hello",
		"extra":    []interface{}{"string", "foo"},
		"extra2": []interface{}{
			map[string]interface{}{"moo": chain.N("eosio")},
		},
	})
}

func TestAbiGetActionFound(t *testing.T) {
	act := tokenAbi.GetAction(chain.N("create"))
	assert.True(t, act != nil)
}

func TestAbiGetActionNotFound(t *testing.T) {
	act := tokenAbi.GetAction(chain.N("not_found"))
	assert.True(t, act == nil)
}

func TestAbiDecodeActionEmptyStruct(t *testing.T) {
	_, err := tokenAbi.DecodeAction(bytes.NewBuffer([]byte{}), chain.N("noop"))
	assert.NoError(t, err)
}

func TestAbiEncodeTable(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	err := tokenAbi.EncodeTable(buf, chain.N("stat"), map[string]interface{}{
		"supply":     *chain.A("17000000.0000 EOS"),
		"max_supply": *chain.A("1800000000.0000 EOS"),
		"issuer":     chain.N("abcdefg12345"),
	})
	assert.NoError(t, err)
	assert.Equal(t, buf.Bytes(), statRow)
}

func TestAbiDecodeTable(t *testing.T) {
	rv, err := tokenAbi.DecodeTable(bytes.NewReader(statRow), chain.N("stat"))
	assert.NoError(t, err)
	assert.Equal(t, rv, map[string]interface{}{
		"supply":     *chain.A("17000000.0000 EOS"),
		"max_supply": *chain.A("1800000000.0000 EOS"),
		"issuer":     chain.N("abcdefg12345"),
	})
}

func TestAbiGetTableFound(t *testing.T) {
	table := tokenAbi.GetTable(chain.N("accounts"))
	assert.True(t, table != nil)
}

func TestAbiGetTableNotFound(t *testing.T) {
	table := tokenAbi.GetTable(chain.N("not_found"))
	assert.True(t, table == nil)
}

func TestAbiDecodeTableEmptyStruct(t *testing.T) {
	_, err := tokenAbi.DecodeTable(bytes.NewBuffer([]byte{}), chain.N("noop"))
	assert.NoError(t, err)
}

func TestAbiDecodeBinary(t *testing.T) {
	fd, err := os.Open("../testdata/abi/setabi.tools.mc.bin")
	assert.NoError(t, err)
	defer fd.Close()

	assert.NoError(t, err)

	dec := abi.NewDecoder(hex.NewDecoder(fd), abi.DefaultDecoderFunc)
	actual := struct {
		Account chain.Name
		Data    chain.Bytes
	}{}
	err = dec.Decode(&actual)
	assert.NoError(t, err)

	fmt.Println(actual.Data)

	abi := chain.Abi{}

	abi_dec := chain.NewDecoder(bytes.NewBuffer(actual.Data))

	err = abi_dec.Decode(&abi)
	assert.NoError(t, err)
}
