package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"feilin.com/gocourse/goaction/crawler/scheduler"
	"feilin.com/gocourse/goaction/crawler/persist"
)

func main() {
	const url = "http://www.zhenai.com/zhenghun"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
