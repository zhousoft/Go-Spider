package main

import (
	"Go-Spider/engine"
	"Go-Spider/scheduler"
	"Go-Spider/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.SimpleScheduler{},
		WokerCount: 2,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
