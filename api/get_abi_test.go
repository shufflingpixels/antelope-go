package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pnx/antelope-go/chain"
	"github.com/stretchr/testify/require"
)

func TestGetAbi(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		require.Equal(t, req.Method, "POST")
		require.Equal(t, req.URL.String(), "/v1/chain/get_abi")
		body := struct {
			AccountName string `json:"account_name"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&body)
		require.NoError(t, err)

		require.Equal(t, body.AccountName, "eosio.token")

		file, err := os.Open("../testdata/api/chain_get_abi.json")
		require.NoError(t, err)
		defer file.Close()
		_, err = io.Copy(res, file)
		require.NoError(t, err)
	}))

	client := New(testServer.URL)

	abi, err := client.GetAbi(context.Background(), "eosio.token")
	require.NoError(t, err)

	expected := AbiResp{
		AccountName: "eosio.token",
		Abi: chain.Abi{
			Version: "eosio::abi/1.1",
			Types:   []chain.AbiType{},
			Structs: []chain.AbiStruct{
				{
					Name: "account",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "balance",
							Type: "asset",
						},
					},
				},
				{
					Name: "close",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "owner",
							Type: "name",
						},
						{
							Name: "symbol",
							Type: "symbol",
						},
					},
				},
				{
					Name: "create",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "issuer",
							Type: "name",
						},
						{
							Name: "maximum_supply",
							Type: "asset",
						},
					},
				},
				{
					Name: "currency_stats",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "supply",
							Type: "asset",
						},
						{
							Name: "max_supply",
							Type: "asset",
						},
						{
							Name: "issuer",
							Type: "name",
						},
					},
				},
				{
					Name: "issue",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "to",
							Type: "name",
						},
						{
							Name: "quantity",
							Type: "asset",
						},
						{
							Name: "memo",
							Type: "string",
						},
					},
				},
				{
					Name: "open",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "owner",
							Type: "name",
						},
						{
							Name: "symbol",
							Type: "symbol",
						},
						{
							Name: "ram_payer",
							Type: "name",
						},
					},
				},
				{
					Name: "retire",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "quantity",
							Type: "asset",
						},
						{
							Name: "memo",
							Type: "string",
						},
					},
				},
				{
					Name: "transfer",
					Base: "",
					Fields: []chain.AbiField{
						{
							Name: "from",
							Type: "name",
						},
						{
							Name: "to",
							Type: "name",
						},
						{
							Name: "quantity",
							Type: "asset",
						},
						{
							Name: "memo",
							Type: "string",
						},
					},
				},
			},
			Actions: []chain.AbiAction{
				{
					Name:              chain.N("close"),
					Type:              "close",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Close Token Balance\nsummary: 'Close {{nowrap owner}}’s zero quantity balance'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/token.png#207ff68b0406eaa56618b08bda81d6a0954543f36adc328ab3065f31a5c5d654\n---\n\n{{owner}} agrees to close their zero quantity balance for the {{symbol_to_symbol_code symbol}} token.\n\nRAM will be refunded to the RAM payer of the {{symbol_to_symbol_code symbol}} token balance for {{owner}}.",
				},
				{
					Name:              chain.N("create"),
					Type:              "create",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Create New Token\nsummary: 'Create a new token'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/token.png#207ff68b0406eaa56618b08bda81d6a0954543f36adc328ab3065f31a5c5d654\n---\n\n{{$action.account}} agrees to create a new token with symbol {{asset_to_symbol_code maximum_supply}} to be managed by {{issuer}}.\n\nThis action will not result any any tokens being issued into circulation.\n\n{{issuer}} will be allowed to issue tokens into circulation, up to a maximum supply of {{maximum_supply}}.\n\nRAM will deducted from {{$action.account}}’s resources to create the necessary records.",
				},
				{
					Name:              chain.N("issue"),
					Type:              "issue",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Issue Tokens into Circulation\nsummary: 'Issue {{nowrap quantity}} into circulation and transfer into {{nowrap to}}’s account'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/token.png#207ff68b0406eaa56618b08bda81d6a0954543f36adc328ab3065f31a5c5d654\n---\n\nThe token manager agrees to issue {{quantity}} into circulation, and transfer it into {{to}}’s account.\n\n{{#if memo}}There is a memo attached to the transfer stating:\n{{memo}}\n{{/if}}\n\nIf {{to}} does not have a balance for {{asset_to_symbol_code quantity}}, or the token manager does not have a balance for {{asset_to_symbol_code quantity}}, the token manager will be designated as the RAM payer of the {{asset_to_symbol_code quantity}} token balance for {{to}}. As a result, RAM will be deducted from the token manager’s resources to create the necessary records.\n\nThis action does not allow the total quantity to exceed the max allowed supply of the token.",
				},
				{
					Name:              chain.N("open"),
					Type:              "open",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Open Token Balance\nsummary: 'Open a zero quantity balance for {{nowrap owner}}'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/token.png#207ff68b0406eaa56618b08bda81d6a0954543f36adc328ab3065f31a5c5d654\n---\n\n{{ram_payer}} agrees to establish a zero quantity balance for {{owner}} for the {{symbol_to_symbol_code symbol}} token.\n\nIf {{owner}} does not have a balance for {{symbol_to_symbol_code symbol}}, {{ram_payer}} will be designated as the RAM payer of the {{symbol_to_symbol_code symbol}} token balance for {{owner}}. As a result, RAM will be deducted from {{ram_payer}}’s resources to create the necessary records.",
				},
				{
					Name:              chain.N("retire"),
					Type:              "retire",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Remove Tokens from Circulation\nsummary: 'Remove {{nowrap quantity}} from circulation'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/token.png#207ff68b0406eaa56618b08bda81d6a0954543f36adc328ab3065f31a5c5d654\n---\n\nThe token manager agrees to remove {{quantity}} from circulation, taken from their own account.\n\n{{#if memo}} There is a memo attached to the action stating:\n{{memo}}\n{{/if}}",
				},
				{
					Name:              chain.N("transfer"),
					Type:              "transfer",
					RicardianContract: "---\nspec_version: \"0.2.0\"\ntitle: Transfer Tokens\nsummary: 'Send {{nowrap quantity}} from {{nowrap from}} to {{nowrap to}}'\nicon: http://127.0.0.1/ricardian_assets/eosio.contracts/icons/transfer.png#5dfad0df72772ee1ccc155e670c1d124f5c5122f1d5027565df38b418042d1dd\n---\n\n{{from}} agrees to send {{quantity}} to {{to}}.\n\n{{#if memo}}There is a memo attached to the transfer stating:\n{{memo}}\n{{/if}}\n\nIf {{from}} is not already the RAM payer of their {{asset_to_symbol_code quantity}} token balance, {{from}} will be designated as such. As a result, RAM will be deducted from {{from}}’s resources to refund the original RAM payer.\n\nIf {{to}} does not have a balance for {{asset_to_symbol_code quantity}}, {{from}} will be designated as the RAM payer of the {{asset_to_symbol_code quantity}} token balance for {{to}}. As a result, RAM will be deducted from {{from}}’s resources to create the necessary records.",
				},
			},
			Tables: []chain.AbiTable{
				{
					Name:      chain.N("accounts"),
					IndexType: "i64",
					KeyNames:  []string{},
					KeyTypes:  []string{},
					Type:      "account",
				},
				{
					Name:      chain.N("stat"),
					IndexType: "i64",
					KeyNames:  []string{},
					KeyTypes:  []string{},
					Type:      "currency_stats",
				},
			},
			RicardianClauses: []chain.AbiClause{},
			ErrorMessages:    []chain.AbiErrorMessage{},
			Extensions:       []*chain.AbiExtension{},
			Variants:         []chain.AbiVariant{},
		},
	}

	require.Equal(t, expected, abi)
}
