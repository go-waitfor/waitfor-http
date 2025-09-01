package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-waitfor/waitfor"
)

const Scheme = "http"

type HTTP struct {
	url *url.URL
}

func Use() waitfor.ResourceConfig {
	return waitfor.ResourceConfig{
		Scheme:  []string{Scheme, Scheme + "s"},
		Factory: New,
	}
}

func New(u *url.URL) (waitfor.Resource, error) {
	if u == nil {
		return nil, fmt.Errorf("%q: %w", "url", waitfor.ErrInvalidArgument)
	}

	return &HTTP{u}, nil
}

func (h *HTTP) Test(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, h.url.String(), nil)

	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req.WithContext(ctx))

	if err != nil {
		return err
	}

	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return errors.New(resp.Status)
	}

	return nil
}
