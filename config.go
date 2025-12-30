package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ClientConfig struct {
	DialTimeout           time.Duration `yaml:"dial_timeout"`
	TLSHandShakeTimeout   time.Duration `yaml:"tls_handshake_timeout"`
	ResponseHeaderTimeout time.Duration `yaml:"response_header_timeout"`
	SkipTLSVerify         bool          `yaml:"skip_tls_verify"`
	ProxyURL              string        `yaml:"proxy_url"`
}

func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		DialTimeout:           5 * time.Second,
		TLSHandShakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ProxyURL:              "",
		SkipTLSVerify:         true,
	}
}

func (c ClientConfig) CreateNewClient() (*http.Client, error) {
	customTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: c.DialTimeout,
		}).DialContext,
		TLSHandshakeTimeout:   c.TLSHandShakeTimeout,
		ResponseHeaderTimeout: c.ResponseHeaderTimeout,
	}
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: c.SkipTLSVerify}
	if c.ProxyURL != "" {
		if url, err := url.Parse(c.ProxyURL); err != nil {
			customTransport.Proxy = http.ProxyURL(url)
		} else {
			return &http.Client{Transport: customTransport}, fmt.Errorf("%w: %v", ErrInvalidProxyURL, err)
		}
	}
	return &http.Client{Transport: customTransport}, nil
}
