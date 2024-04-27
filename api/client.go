package api

import (
	"context"
	"net/url"

	"github.com/imroc/req/v3"
)

type Client struct {
	Url    string
	Host   string
	client *req.Client
}

func New(url string) *Client {
	rc := req.C().
		SetJsonMarshal(json.Marshal).
		SetJsonUnmarshal(customJsonUnmarshal)

	return &Client{
		Url:    url,
		Host:   "",
		client: rc,
	}
}

func (c *Client) send(ctx context.Context, method string, path string, body interface{}, out interface{}) error {
	host := c.Host
	if len(host) < 1 {
		u, err := url.Parse(c.Url)
		if err != nil {
			return err
		}
		host = u.Host
	}

	// Go's net.http (that `req` uses) sends the port in the host header.
	// nodeos api does not like that, so we need to provide our
	// own Host header with just the host.
	r, err := c.client.R().
		SetContext(ctx).
		SetHeader("Host", host).
		SetBody(body).
		Send(method, c.Url+path)
	if err != nil {
		return err
	}

	if !r.IsError() {
		return r.UnmarshalJson(&out)
	}

	return handleError(r)
}

func handleError(r *req.Response) error {
	var api_err APIError
	// Parse error object.
	err := r.UnmarshalJson(&api_err)
	if err != nil || api_err.IsEmpty() {
		// Failed to parse error object. just return an generic HTTP error
		return HTTPError{Code: r.StatusCode}
	}
	return api_err
}
