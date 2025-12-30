package client

import (
	"errors"
)

var ErrNilPtrClient = errors.New("http.Client pointer is nil")
var ErrInvalidProxyURL = errors.New("invalid proxy URL")
