package ship_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/pnx/antelope-go/abi"
	"github.com/pnx/antelope-go/internal/assert"
	"github.com/pnx/antelope-go/ship"
)

func TestTableDeltaEncode(t *testing.T) {
	table := ship.TableDelta{
		V0: &ship.TableDeltaV0{
			Name: "some_name",
			Rows: []ship.Row{
				{
					Present: true,
					Data:    []byte{0x01, 0x02, 0x03},
				},
				{
					Present: false,
					Data:    []byte{0x04, 0x05, 0x06},
				},
			},
		},
	}

	expected := []byte{
		0x00, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6e,
		0x61, 0x6d, 0x65, 0x02, 0x01, 0x03, 0x01, 0x02,
		0x03, 0x00, 0x03, 0x04, 0x05, 0x06,
	}

	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(table)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, data, expected)
}

func TestTableDeltaDecode(t *testing.T) {
	data := []byte{
		0x00, 0x0f, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6f,
		0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
		0x65, 0x02, 0x01, 0x03, 0x21, 0x22, 0x23, 0x00,
		0x03, 0x44, 0x45, 0x46,
	}

	actual := ship.TableDelta{}
	err := abi.NewDecoder(bytes.NewBuffer(data), abi.DefaultDecoderFunc).Decode(&actual)
	assert.NoError(t, err)

	expected := ship.TableDelta{
		V0: &ship.TableDeltaV0{
			Name: "some_other_name",
			Rows: []ship.Row{
				{
					Present: true,
					Data:    []byte{0x21, 0x22, 0x23},
				},
				{
					Present: false,
					Data:    []byte{0x44, 0x45, 0x46},
				},
			},
		},
	}

	assert.Equal(t, actual, expected)
}

func TestTableDeltaArrayUnpack(t *testing.T) {
	delta := ship.TableDelta{
		V0: &ship.TableDeltaV0{
			Name: "unpack_me",
			Rows: []ship.Row{
				{
					Present: true,
					Data:    []byte{0x01, 0x02, 0x03},
				},
				{
					Present: false,
					Data:    []byte{0x04, 0x05, 0x06},
				},
			},
		},
	}

	actual := []ship.TableDelta{}
	expected := []ship.TableDelta{delta}
	arr := ship.MustMakeTableDeltaArray(expected)

	err := arr.Unpack(&actual)
	assert.NoError(t, err)
	assert.Equal(t, actual, expected)
}
