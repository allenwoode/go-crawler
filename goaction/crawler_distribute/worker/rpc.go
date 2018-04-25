package worker

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"log"
)

type CrawlService struct {}

func (CrawlService) Process(r Request, res *ParseResult) error {
	req, err := DeserializeRequest(r)
	if err != nil {
		return err
	}

	result, err := engine.Worker(req)
	if err != nil {
		return err
	}

	*res = SerializeResult(result)
	log.Printf("WorkerService request: %s", r.String())
	return nil
}
