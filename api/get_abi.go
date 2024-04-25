package api

import (
	"context"

	"github.com/pnx/antelope-go/chain"
)

type AbiResp struct {
	AccountName string    `json:"account_name"`
	Abi         chain.Abi `json:"abi"`
}

//	GetAbi - Fetches "/v1/chain/get_abi" from API
//
// ---------------------------------------------------------
func (c *Client) GetAbi(ctx context.Context, account string) (abi AbiResp, err error) {
	body := map[string]string{
		"account": account,
	}

	err = c.send(ctx, "GET", "/v1/chain/get_abi", body, &abi)
	return
}
