package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// Fetch url
func Fetch(url string) ([]byte, error) {
	log.Printf("Fetcher: fetching url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("Resp status fail, code is : %d", resp.StatusCode)
	}
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	return response, nil
}
