package scheduler

import "sky.com/case/crawler-go/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	s.WorkerChan = r
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}
