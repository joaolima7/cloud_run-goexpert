package httpclient

import "context"

type HTTPClient interface {
	Get(ctx context.Context, url string) ([]byte, error)
}
