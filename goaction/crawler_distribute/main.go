package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"feilin.com/gocourse/goaction/crawler/scheduler"
	"log"
	"feilin.com/gocourse/goaction/crawler_distribute/itemsaver"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	//e := engine.SimpleEngine{}

	const url = "http://www.zhenai.com/zhenghun"
	log.Printf("distribute crawler start at %s", url)
	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
