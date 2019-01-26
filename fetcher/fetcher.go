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
	request, err := http.NewRequest(
		http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(request)
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
