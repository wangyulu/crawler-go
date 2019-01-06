package main

import (
	"sky.com/case/crawler-go/engine"
	"sky.com/case/crawler-go/scheduler"
	"sky.com/case/crawler-go/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList},
	)
}

func t1() {
	// 城市列表
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	// 城市用户列表
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/anhui",
		ParserFunc: parser.ParseCity,
	})*/

	// 用户详情
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://album.zhenai.com/u/76206456",
		ParserFunc: parser.ParseProfile,
	})*/
}
