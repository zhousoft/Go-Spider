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
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWokerChan(in)

	for i := 0; i < e.WokerCount; i++ {
		createWoker(in, out)
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

func createWoker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := woker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
