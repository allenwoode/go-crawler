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

	choosed := matches[:10]
	for _, m := range choosed {
		results.Items = append(results.Items, string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return results
}