package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendContextTimeout(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second * 4)
	}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client := New(testServer.URL)

	err := client.send(ctx, "GET", "/", nil, nil)
	assert.Error(t, err)
	assert.True(t, strings.HasSuffix(err.Error(), "deadline exceeded"), "Error was not deadline exceeded")
}

func TestSendContextCancel(t *testing.T) {
	done := make(chan interface{})
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second * 10)
	}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	client := New(testServer.URL)

	go func() {
		defer close(done)
		err := client.send(ctx, "GET", "/", nil, nil)
		assert.Error(t, err)
		assert.True(t, strings.HasSuffix(err.Error(), "context canceled"), "Error was not context canceled")
	}()

	time.Sleep(time.Second)
	cancel()

	<-done
}

func TestSendUrlParseFails(t *testing.T) {
	client := New("api.mylittleponies.org\n")

	err := client.send(context.Background(), "GET", "/v1/ponies/Rainbow Dash", nil, nil)
	assert.EqualError(t, err, "parse \"api.mylittleponies.org\\n\": net/url: invalid control character in URL")
}

func TestSendDefaultHostHeader(t *testing.T) {
	expected := ""

	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, expected, req.Host)
	}))

	u, err := url.Parse(srv.URL)
	require.NoError(t, err)

	expected = u.Host
	client := New(srv.URL)

	err = client.send(context.Background(), "GET", "/", nil, nil)
	require.NoError(t, err)
}

func TestSendCustomHostHeader(t *testing.T) {
	expected := "CustomHost"

	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, expected, req.Host)
	}))

	client := New(srv.URL)
	client.Host = expected

	err := client.send(context.Background(), "GET", "/", nil, nil)
	require.NoError(t, err)
}
