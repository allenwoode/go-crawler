package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"log"
)

func main() {
	//itemChan, err := persist.ItemSaver("dating_profile")
	//if err != nil {
	//	panic(err)
	//}
	//
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//	ItemChan:    itemChan,
	//	ReqProcessor: engine.Worker,
	//}

	e := engine.SimpleEngine{}
	const url = "http://www.zhenai.com/zhenghun"
	log.Printf("concurrent crawler start at %s", url)
	e.Run(engine.Request{
		Url:        url,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
