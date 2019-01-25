package engin

import (
	"fmt"
	"log"
	"time"

	"github.com/mozyy/crawler/fetcher"
)

// RunSync is start fetch
func RunSync(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	loop(requests)
}

func loop(requests []Request) {
	chs := make(chan Result, len(requests))
	for _, r := range requests {
		go syncFetcher(r, chs)
	}
	for {
		select {
		case result := <-chs:
			for _, item := range result.Items {
				log.Printf("Got item %s", item)
				count++
			}
			fmt.Printf("count: %d, start: %v, end: %v \n", count, start, time.Now())
			go loop(result.Requests)
		}
	}
}

func syncFetcher(r Request, ch chan Result) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
	} else {
		ch <- r.Parser(body)
	}
}
