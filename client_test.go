package client

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	client, err := DefaultClientConfig().CreateNewClient()
	if err != nil {
		if errors.Is(err, ErrInvalidProxyURL) {
			t.Fatalf("proxy configuration error: %v", err)
		}
		t.Fatal(err)
	}
	r := Request{
		"https://example.com",
		http.Header{
			"Accept": {"application/json", "text/plain"},
		},
		client,
	}
	resp, err := r.SendPost(strings.NewReader(`{data: "hello server"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	t.Log(resp)
}
