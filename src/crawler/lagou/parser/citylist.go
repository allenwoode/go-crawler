package parser

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
	"strings"
)

func ParseCityList(contents []byte) engine.ParseResult {
	results := engine.ParseResult{}

	sel, err := css.Compile(".city_list>li>input")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}

	for _, ele := range sel.Select(node) {
		results.Requests = append(results.Requests, engine.Request{
			Url:    ele.Attr[1].Val,
			Parser: engine.NilParser{},
		})
	}

	return results
}
