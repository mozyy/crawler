package engin

import (
	"fmt"
	"log"
	"time"

	"github.com/mozyy/crawler/fetcher"
)

var (
	count int
	start = time.Now()
)

// Run is start fetch
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v",
				r.URL, err)
			continue
		}
		parseResult := r.Parser(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
			count++
		}
		fmt.Printf("count: %d, start: %v, end: %v", count, start, time.Now())
	}
}
