package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"feilin.com/gocourse/goaction/crawler/scheduler"
	"log"
	itemsaver "feilin.com/gocourse/goaction/crawler_distribute/itemsaver"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	worker "feilin.com/gocourse/goaction/crawler_distribute/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
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
