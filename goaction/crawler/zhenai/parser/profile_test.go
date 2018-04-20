package parser

import (
	"testing"
	"io/ioutil"
	"feilin.com/gocourse/goaction/crawler/model"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(contents))
	result := ParseProfile(contents, "萧然")

	if len(result.Items) < 1 {
		t.Errorf("Result should contain 1 element, but actual was %v", result.Items)
	}

	expected := model.Profile{
		Name:      "萧然",
		Age:       28,
		Height:    160,
		Weight:    0,
		Income:    "3000元以下",
		Marriage:  "离异",
		Education: "大专",
		Occupation: "--",
		Hokou:     "上海浦东新区",
		Gender:    "女",
		Xinzuo:    "双子座",
		House:     "--",
		Car:       "未购车",
	}

	profile := result.Items[0].(model.Profile)
	if profile != expected {
		t.Errorf("expected %v, but %v", expected, profile)
	}
}
