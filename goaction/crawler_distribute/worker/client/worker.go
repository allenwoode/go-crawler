package client

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	"feilin.com/gocourse/goaction/crawler_distribute/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		serializeReq := worker.SerializeRequest(req)

		var serializeRes worker.ParseResult

		err := client.Call(config.CrawlServiceRpc, serializeReq, &serializeRes)
		if err != nil {
			return engine.ParseResult{}, nil
		}

		return worker.DeserializeResult(serializeRes), nil
	}, nil
}
