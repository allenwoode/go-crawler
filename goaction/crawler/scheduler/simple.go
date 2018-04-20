package scheduler

import "feilin.com/gocourse/goaction/crawler/engine"

// 简易调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(req chan engine.Request) {
	//panic("implement me")
	s.workerChan = req
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//panic("implement me")
	//s.workerChan <- r
	// 避免循环等待
	go func() {
		s.workerChan <- r
	}()
}

