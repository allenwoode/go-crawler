package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"log"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e, simple := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:		&scheduler.SimpleScheduler{},
		WorkerCount:  100,
		ItemChan:     itemChan,
		ReqProcessor: engine.Worker,
	}, false

	//e, simple := engine.SimpleEngine{}, true
	const url = "http://www.zhenai.com/zhenghun"
	if simple {
		log.Printf("simple engine crawler start at %s", url)
	} else {
		log.Printf("concurrent engine crawler start at %s", url)
	}

	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
