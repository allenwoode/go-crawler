package engine

import (
	"feilin.com/gocourse/goaction/crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

// Simple single task engine
func (e *SimpleEngine) Run(requests ...Request)  {
	var queue []Request
	for _, req := range requests {
		queue = append(queue, req)
	}

	count := 0
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		result, err := Worker(q)
		if err != nil {
			continue
		}
		queue = append(queue, result.Requests...)

		for _, item := range result.Items {
			count++
			log.Printf("Got item #%d: %v", count, item)
		}
	}
}

// Engine fetcher
func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	//////////////////////////////////
	body, err := fetcher.Fetch(r.Url)
	//////////////////////////////////
	if err != nil {
		log.Printf("Error: %v Url: %s", err, r.Url)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
