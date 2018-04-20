package parser

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"regexp"
)

const cityListRegex  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRegex)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("city: %s url: %s\n", m[2], m[1])
		results.Items = append(results.Items, string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return results
}