package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"feilin.com/gocourse/goaction/crawler/scheduler"
	"log"
	itemsaver "feilin.com/gocourse/goaction/crawler_distribute/itemsaver"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	worker "feilin.com/gocourse/goaction/crawler_distribute/worker/client"
	"flag"
	"strings"
	"net/rpc"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
)

var (
	saverHost = flag.String("saver_host", "", "saver host")
	workerHosts = flag.String("worker_host", "", "worker hosts (comma sep)")
)

func main()  {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*saverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor, err := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		ReqProcessor: processor,
	}

	//e := engine.SimpleEngine{}

	const url = "http://www.zhenai.com/zhenghun"
	log.Printf("distributed crawler start at %s", url)
	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connect to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
