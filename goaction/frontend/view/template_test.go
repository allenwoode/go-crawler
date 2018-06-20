package view

import (
	"testing"
	"feilin.com/gocourse/goaction/frontend/model"
	common "feilin.com/gocourse/goaction/crawler/model"
	"os"
	"feilin.com/gocourse/goaction/crawler/engine"
)

func TestTemplate(t *testing.T)  {
	view := CreateSearchResultView("template.html")

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/107072849",
		Type: "zhenai",
		Id:   "107072849",
		Payload: common.Profile{
			Name:       "萧然",
			Age:        28,
			Height:     160,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "大专",
			Occupation: "--",
			Hokou:      "上海浦东新区",
			Gender:     "女",
			Xinzuo:     "双子座",
			House:      "--",
			Car:        "未购车",
		},
	}

	page := model.SearchResult{
		Hits: 133,
	}

	for i := 0; i < 5; i++ {
		page.Items = append(page.Items, item)
	}

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
