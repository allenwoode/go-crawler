package main

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
)
/*
	html内容精准获取方法
	1. css选择器 e.g. $('#cityList>dd>a') get all a dom
	2. xpath
	3. regex
 */

 // 抽象成解析器，即城市列表解析器
 // 输入：uft-8html文本
 // 输出：request{URL, 对应Parser}, 解析返回

func main() {
	const url = "http://www.zhenai.com/zhenghun"

	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}

	e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}
