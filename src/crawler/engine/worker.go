package engine

import (
	"log"
	"crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)

	//////////////////////////////////
	body, err := fetcher.Fetch(r.Url)
	//////////////////////////////////

	if err != nil {
		log.Printf("Fetching error: %v. Url: %s", err, r.Url)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
