package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/config"
)

/* 解析器抽象
   输入：uft-8html文本
   输出：request{URL, 对应Parser}, 解析返回
*/

/*
	解析HTML内容方法
	1. css选择器 e.g. $('#cityList>dd>a') get all a dom
	2. xpath
	3. regex - 通用性更好
*/

var cityListRe  = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//const cityListRegex  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	results := engine.ParseResult{}

	matches := cityListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//results.Items = append(results.Items, string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return results
}