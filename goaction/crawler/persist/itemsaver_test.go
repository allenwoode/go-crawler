package persist

import (
	"testing"
	"feilin.com/gocourse/goaction/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSaver(t *testing.T) {

	expected := model.Profile{
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
		Url:        "",
	}

	id, err := save(expected)
	if err != nil {
		t.Errorf("%v", err)
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search client
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	//t.Logf("%s\n", *resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
