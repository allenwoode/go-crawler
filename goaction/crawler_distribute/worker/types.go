package worker

import (
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	"feilin.com/gocourse/goaction/crawler/zhenai/parser"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializeParser
}

func (r *Request) String() string {
	return fmt.Sprintf("url: %s, parser: %s, args: %v", r.Url, r.Parser.Name, r.Parser.Args)
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializeParser{
			Name:name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	
	for _, req := range r.Requests {

		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url: r.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("deserialize request error: %v", err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializeParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unkonw parser name")
	}
}

