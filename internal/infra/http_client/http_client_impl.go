package httpclient

import (
	"context"
	"io"
	"net/http"
	"time"
)

type HTTPClientImpl struct {
	client *http.Client
}

func NewHTTPClientImpl(timeout time.Duration) *HTTPClientImpl {
	return &HTTPClientImpl{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPClientImpl) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
