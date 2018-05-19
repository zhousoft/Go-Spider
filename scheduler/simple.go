package scheduler

import (
	"Go-Spider/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWokerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) WokerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
func (s *SimpleScheduler) WokerReady(chan engine.Request) {

}
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
