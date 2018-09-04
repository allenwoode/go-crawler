package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"log"
	"github.com/go-redis/redis"
	"crawler/config"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}



	e := engine.ConcurrentEngine {
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:		&scheduler.SimpleScheduler{},
		WorkerCount:  100,
		ItemChan:     itemChan,
		ReqProcessor: engine.Worker,
		RedisClient: redis.NewClient(&redis.Options{
			Addr: config.RedisHost,
			Password: config.RedisPassword,
			DB: config.RedisDB,
		}),
	}

	//e, simple := engine.SimpleEngine{}, true
	const url = "http://www.zhenai.com/zhenghun"
	log.Printf("concurrent engine crawler start at %s", url)

	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
