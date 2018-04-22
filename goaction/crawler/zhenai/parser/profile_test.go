package parser

import (
	"testing"
	"io/ioutil"
	"feilin.com/gocourse/goaction/crawler/model"
	"feilin.com/gocourse/goaction/crawler/engine"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/107072849",
		Type: "zhenai",
		Id:   "107072849",
		Payload: model.Profile{
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

	result := ParseProfile(contents, "http://album.zhenai.com/u/107072849","萧然")

	if len(result.Items) < 1 {
		t.Errorf("Result should contain 1 element, but actual was %v", result.Items)
	}

	actual := result.Items[0]
	if actual != expected {
		t.Errorf("expected %v, but %v", expected, actual)
	}
}
