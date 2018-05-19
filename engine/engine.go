package engine

import (
	"Go-Spider/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//广度优先遍历
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching: %s", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch error, url: %s, error: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests,
			parseResult.Requests...) //新链接进队  ...可以将数组打散进行append

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
