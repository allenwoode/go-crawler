package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"feilin.com/gocourse/goaction/crawler/scheduler"
)

func main() {
	const url = "http://www.zhenai.com/zhenghun"

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}
