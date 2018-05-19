package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler  Scheduler
	WokerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWokerChan(chan Request)
	WokerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WokerCount; i++ {
		createWoker(out, e.Scheduler)
	}

	for _, url := range seeds {
		e.Scheduler.Submit(url)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWoker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WokerReady(in)
			request := <-in
			result, err := woker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
