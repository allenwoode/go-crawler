package parser

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"regexp"
)

const cityRegex  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRegex)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		results.Items = append(results.Items, name)
		results.Requests = append(results.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, url)
			},
		})
	}

	return results
}