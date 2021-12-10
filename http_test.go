package http_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-http"
)

func TestUse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	uStr := "http://api.test.com/"

	httpmock.RegisterResponder("GET", uStr,
		httpmock.NewStringResponder(200, `OK`))

	w := waitfor.New(http.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := w.Test(ctx, []string{uStr})

	assert.NoError(t, err)
}

func TestHTTP(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	uStr := "http://api.test.com/"

	httpmock.RegisterResponder("GET", uStr,
		httpmock.NewStringResponder(200, `OK`))

	u, err := url.Parse(uStr)

	assert.NoError(t, err)

	r, err := http.New(u)

	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	assert.NoError(t, err)
}

func TestHTTP_Failed(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	uStr := "http://foobar"

	httpmock.RegisterResponder("GET", uStr,
		httpmock.NewStringResponder(404, ``))

	u, err := url.Parse(uStr)

	assert.NoError(t, err)

	r, err := http.New(u)

	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	assert.Error(t, err)
}

func TestHTTPS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	uStr := "https://api.test.com/"

	httpmock.RegisterResponder("GET", uStr,
		httpmock.NewStringResponder(200, `OK`))

	u, err := url.Parse(uStr)

	assert.NoError(t, err)

	r, err := http.New(u)

	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	assert.NoError(t, err)
}
