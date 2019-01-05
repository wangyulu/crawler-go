package engine

import (
	"sky.com/case/crawler-go/fetcher"
	"log"
)

func Run(sends ...Request) {
	var requests []Request
	for _, r := range sends {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:err url %s:%v", r.Url, err)
			continue
		}

		parserRes := r.ParserFunc(body)
		requests = append(requests, parserRes.Requests...)

		for _, item := range parserRes.Items {
			log.Printf("item %q", item)
		}
	}

}
