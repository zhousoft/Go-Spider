package engine

import (
	"Go-Spider/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//广度优先遍历
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := woker(r)
		if err != nil {
			continue
		}

		requests = append(requests,
			parseResult.Requests...) //新链接进队  ...可以将数组打散进行append

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}

//woker 获取并解析url，返回解析结果
func woker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch error, url: %s, error: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
