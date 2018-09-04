package engine

// 并发爬虫引擎
type ConcurrentEngine struct {
	Scheduler    Scheduler
	WorkerCount  int
	ItemChan     chan Item
	ReqProcessor Processor
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
			if isDuplicate(r.Url) {
				//log.Printf("Duplicate url: %s", r.Url)
				continue
			}
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
