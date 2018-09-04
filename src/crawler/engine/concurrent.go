package engine

import (
	"github.com/go-redis/redis"
	"crawler/util"
	"strings"
	"github.com/uber/tchannel-go/crossdock/log"
)

// 并发爬虫引擎
type ConcurrentEngine struct {
	Scheduler    Scheduler
	WorkerCount  int
	ItemChan     chan Item
	ReqProcessor Processor
	RedisClient	*redis.Client
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier //接口组合

	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate url: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		// URL去重 内存哈希表方式
		for _, r := range result.Requests {
			if e.dup(r.Url) {
				log.Printf("Duplicate url: %s", r.Url)
				continue
			}
			log.Printf("submit url: %s", r.Url)
			e.Scheduler.Submit(r)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)

			request := <-in
			result, err := e.ReqProcessor(request) //worker
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	return false
}

func (e *ConcurrentEngine) dup(url string) bool {
	key := util.GetMD5Hash(url)
	val, err := e.RedisClient.Get(key).Result()
	if err != nil {
		log.Printf("set %s to redis", url)
		err = e.RedisClient.Set(key, url, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	return strings.EqualFold(val, url)
}