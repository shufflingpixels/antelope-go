package ship_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/pnx/antelope-go/abi"
	"github.com/pnx/antelope-go/internal/assert"
	"github.com/pnx/antelope-go/ship"
)

func TestStatusRequestEncode(t *testing.T) {
	req := ship.Request{
		StatusRequest: &ship.GetStatusRequestV0{},
	}

	expected := []byte{0x00}

	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(req)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestStatusRequestDecode(t *testing.T) {
	data := []byte{0x00}

	dec := abi.NewDecoder(bytes.NewReader(data), abi.DefaultDecoderFunc)
	req := ship.Request{}
	err := dec.Decode(&req)
	assert.NoError(t, err)

	expected := ship.Request{
		StatusRequest: &ship.GetStatusRequestV0{},
	}

	assert.Equal(t, expected, req)
}

func TestBlocksRequestEncode(t *testing.T) {
	req := ship.Request{
		BlocksRequest: &ship.GetBlocksRequestV0{
			StartBlockNum:       1000,
			EndBlockNum:         2000,
			MaxMessagesInFlight: 20,
			HavePositions: []*ship.BlockPosition{
				{
					BlockNum: 3400,
					BlockID: [32]byte{
						0xc1, 0x3b, 0x28, 0x65, 0x80, 0xb4, 0x1c, 0x5d, 0x0c, 0x13, 0xcc, 0xe7, 0xa0, 0xe4, 0x15, 0xb3,
						0x4e, 0xdc, 0x55, 0x8d, 0x23, 0x62, 0x72, 0x1b, 0xb2, 0x3d, 0xfb, 0x44, 0x9b, 0x0f, 0x4e, 0x4f,
					},
				},
			},
			IrreversibleOnly: false,
			FetchBlock:       true,
			FetchTraces:      false,
			FetchDeltas:      true,
		},
	}

	expected := []byte{
		0x01, 0xe8, 0x03, 0x00, 0x00, 0xd0, 0x07, 0x00,
		0x00, 0x14, 0x00, 0x00, 0x00, 0x01, 0x48, 0x0d,
		0x00, 0x00, 0xc1, 0x3b, 0x28, 0x65, 0x80, 0xb4,
		0x1c, 0x5d, 0x0c, 0x13, 0xcc, 0xe7, 0xa0, 0xe4,
		0x15, 0xb3, 0x4e, 0xdc, 0x55, 0x8d, 0x23, 0x62,
		0x72, 0x1b, 0xb2, 0x3d, 0xfb, 0x44, 0x9b, 0x0f,
		0x4e, 0x4f, 0x00, 0x01, 0x00, 0x01,
	}

	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(req)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestBlocksRequestDecode(t *testing.T) {
	data := []byte{
		0x01, 0xe8, 0x03, 0x00, 0x00, 0xd0, 0x07, 0x00,
		0x00, 0x14, 0x00, 0x00, 0x00, 0x01, 0x48, 0x0d,
		0x00, 0x00, 0xc1, 0x3b, 0x28, 0x65, 0x80, 0xb4,
		0x1c, 0x5d, 0x0c, 0x13, 0xcc, 0xe7, 0xa0, 0xe4,
		0x15, 0xb3, 0x4e, 0xdc, 0x55, 0x8d, 0x23, 0x62,
		0x72, 0x1b, 0xb2, 0x3d, 0xfb, 0x44, 0x9b, 0x0f,
		0x4e, 0x4f, 0x00, 0x01, 0x00, 0x01,
	}

	dec := abi.NewDecoder(bytes.NewReader(data), abi.DefaultDecoderFunc)
	req := ship.Request{}
	err := dec.Decode(&req)
	assert.NoError(t, err)

	expected := ship.Request{
		BlocksRequest: &ship.GetBlocksRequestV0{
			StartBlockNum:       1000,
			EndBlockNum:         2000,
			MaxMessagesInFlight: 20,
			HavePositions: []*ship.BlockPosition{
				{
					BlockNum: 3400,
					BlockID: [32]byte{
						0xc1, 0x3b, 0x28, 0x65, 0x80, 0xb4, 0x1c, 0x5d, 0x0c, 0x13, 0xcc, 0xe7, 0xa0, 0xe4, 0x15, 0xb3,
						0x4e, 0xdc, 0x55, 0x8d, 0x23, 0x62, 0x72, 0x1b, 0xb2, 0x3d, 0xfb, 0x44, 0x9b, 0x0f, 0x4e, 0x4f,
					},
				},
			},
			IrreversibleOnly: false,
			FetchBlock:       true,
			FetchTraces:      false,
			FetchDeltas:      true,
		},
	}

	assert.Equal(t, req, expected)
}

func TestBlocksAckRequestEncode(t *testing.T) {
	req := ship.Request{
		BlocksAckRequest: &ship.GetBlocksAckRequestV0{
			NumMessages: 27123681,
		},
	}

	expected := []byte{0x02, 0xe1, 0xdf, 0x9d, 0x01}

	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(req)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, data, expected)
}

func TestBlocksAckRequestDecode(t *testing.T) {
	data := []byte{0x02, 0xe1, 0xdf, 0x9d, 0x01}

	dec := abi.NewDecoder(bytes.NewReader(data), abi.DefaultDecoderFunc)
	req := ship.Request{}
	err := dec.Decode(&req)
	assert.NoError(t, err)

	expected := ship.Request{
		BlocksAckRequest: &ship.GetBlocksAckRequestV0{
			NumMessages: 27123681,
		},
	}

	assert.Equal(t, req, expected)
}
