package engine

import (
	"log"
	"feilin.com/gocourse/goaction/crawler/fetcher"
)

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

	return r.ParserFunc(body, r.Url), nil
}
