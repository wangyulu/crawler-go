package parser

import (
	"regexp"
	"sky.com/case/crawler-go/engine"
)

// [^>]*表示：只要不是>括号，如果有就匹配，没有就不匹配
// [^<]+表示：只要不是<括号，且至少匹配一个字符

// url: http://www.zhenai.com/zhenghun/anhui
var cityListRgx = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(content []byte) engine.ParserRes {
	matchs := cityListRgx.FindAllSubmatch(content, -1)

	result := engine.ParserRes{}
	for _, item := range matchs {
		result.Items = append(result.Items, string(item[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(item[1]), ParserFunc: ParseCity})
	}

	return result
}
