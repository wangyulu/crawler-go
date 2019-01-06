package engine

import (
	"sky.com/case/crawler-go/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(sends ...Request) {
	var requests []Request
	for _, r := range sends {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserRes, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserRes.Requests...)

		for _, item := range parserRes.Items {
			log.Printf("item %v", item)
		}
	}
}

func worker(r Request) (ParserRes, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher-err url: %s; err: %v", r.Url, err)
		return ParserRes{}, err
	}

	return r.ParserFunc(body), nil
}
