package main

import (
	"sky.com/case/crawler-go/engine"
	"sky.com/case/crawler-go/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
