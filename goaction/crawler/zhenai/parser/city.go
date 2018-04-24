package parser

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(cityRegex)
	results := engine.ParseResult{}

	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//results.Items = append(results.Items, name)
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//name := string(m[2])
		url := string(m[1])
		//results.Items = append(results.Items, name)
		results.Requests = append(results.Requests, engine.Request{
			Url: url,
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		//log.Printf("other city url:%s", url)
	}

	return results
}