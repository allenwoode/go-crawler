package main

import (
	"testing"
	"crawler_distribute/rpcsupport"
	"time"
	"crawler_distribute/worker"
	"crawler_distribute/config"
	"fmt"
)

func TestCrawlService(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.NewClient(host)
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/107072849",
		Parser: worker.SerializeParser{
			Name: config.ParseProfile,
			Args: "萧然",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
