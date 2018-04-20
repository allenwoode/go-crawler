package engine

import (
	"feilin.com/gocourse/goaction/crawler/fetcher"
	"log"
)

// 爬虫引擎主方法
func Run(requests ...Request)  {
	var queue []Request
	for _, req := range requests {
		queue = append(queue, req)
	}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		log.Printf("Fetching %s", q.Url)
		body, err := fetcher.Fetch(q.Url)
		if err != nil {
			log.Printf("Error: %v Url: %s", err, q.Url)
			continue
		}

		parserResult := q.ParserFunc(body)
		queue = append(queue, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
