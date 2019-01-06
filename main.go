package main

import (
	"sky.com/case/crawler-go/engine"
	"sky.com/case/crawler-go/zhenai/parser"
)

func main() {
	// 城市列表
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 城市用户列表
	/*engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/anhui",
		ParserFunc: parser.ParseCity,
	})*/

	// 用户详情
	/*engine.Run(engine.Request{
		Url:        "http://album.zhenai.com/u/76206456",
		ParserFunc: parser.ParseProfile,
	})*/
}
