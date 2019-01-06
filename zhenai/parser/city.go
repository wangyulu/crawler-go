package parser

import (
	"sky.com/case/crawler-go/engine"
	"regexp"
)

// url: http://album.zhenai.com/u/80216934
var cityRgx = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)

func ParseCity(content []byte) engine.ParserRes {
	matchs := cityRgx.FindAllSubmatch(content, -1)
	result := engine.ParserRes{}
	if len(matchs) <= 0 {
		return result
	}

	for _, item := range matchs {
		name := string(item[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParserFunc: func(bytes []byte) engine.ParserRes {
				return ParseProfile(bytes, name)
			}})
		result.Items = append(result.Items, []interface{}{string(item[2])})
	}

	return result
}
