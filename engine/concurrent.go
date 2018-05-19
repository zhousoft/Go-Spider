package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler  Scheduler
	WokerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WokerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WokerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WokerCount; i++ {
		createWoker(e.Scheduler.WokerChan(), out, e.Scheduler)
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

func createWoker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WokerReady(in)
			request := <-in
			result, err := woker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
