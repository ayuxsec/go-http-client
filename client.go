package client

import (
	"io"
	"net/http"
)

type Request struct {
	RawURL     string
	Headers    http.Header
	HTTPClient *http.Client
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

// Note: caller MUST close Body
func (r Request) SendPost(body io.Reader) (*Response, error) {
	if r.HTTPClient == nil {
		return &Response{}, ErrNilPtrClient
	}

	req, err := http.NewRequest(http.MethodPost, r.RawURL, body)
	if err != nil {
		return nil, err
	}

	req.Header = r.Headers.Clone()

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Header:     resp.Header.Clone(),
		Body:       resp.Body,
	}, nil
}
