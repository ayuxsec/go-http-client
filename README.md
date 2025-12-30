```go
import client "github.com/ayuxsec/go-http-client" // go get github.com/ayuxsec/go-http-client@latest

import (
	"errors"
	"net/http"
	"strings"
	"testing"
    "log"
)

func main() {
    cfg := client.DefaultClientConfig()
    cfg.ProxyURL = "http://127.0.0.1:8080"
	c, err := client.DefaultClientConfig().CreateNewClient()
	if err != nil {
		if errors.Is(err, client.ErrInvalidProxyURL) {
            log.Print("[warn] proxy will be ignored. proxy configuration error: ", err)
		}
	}

	r := client.Request{
		RawURL:  "https://api.github.com/markdown",
		Headers: nil,
		Client:  c,
	}

	resp, err := r.SendPost(strings.NewReader(`{data: "hello server"}`))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // must
	log.Print(resp)
}
```