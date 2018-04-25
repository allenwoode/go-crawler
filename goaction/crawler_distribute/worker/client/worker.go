package client

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	"feilin.com/gocourse/goaction/crawler_distribute/worker"
	"net/rpc"
)

func CreateProcessor(client chan *rpc.Client) (engine.Processor, error) {

	return func(req engine.Request) (engine.ParseResult, error) {
		serializeReq := worker.SerializeRequest(req)

		var serializeRes worker.ParseResult
		c := <-client
		err := c.Call(config.CrawlServiceRpc, serializeReq, &serializeRes)
		if err != nil {
			return engine.ParseResult{}, nil
		}

		return worker.DeserializeResult(serializeRes), nil
	}, nil
}
