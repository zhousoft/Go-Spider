package scheduler

import (
	"Go-Spider/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	wokerChan   chan chan engine.Request
}

func (s *QueuedScheduler) ConfigMasterWokerChan(c chan engine.Request) {
	s.requestChan = c
}

func (s *QueuedScheduler) WokerReady(w chan engine.Request) {
	s.wokerChan <- w
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.wokerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var wokerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWoker chan engine.Request
			if len(requestQ) > 0 && len(wokerQ) > 0 {
				activeRequest = requestQ[0]
				activeWoker = wokerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.wokerChan:
				wokerQ = append(wokerQ, w)
			case activeWoker <- activeRequest:
				requestQ = requestQ[1:]
				wokerQ = wokerQ[1:]
			}
		}
	}()
}
