package scheduler

import "feilin.com/gocourse/goaction/crawler/engine"

// 简易调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//s.workerChan <- r
	// 避免循环等待
	go func() {
		s.workerChan <- r
	}()
}
