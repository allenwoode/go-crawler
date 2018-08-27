package persist

import (
	"testing"
	"feilin.com/gocourse/goaction/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"feilin.com/gocourse/goaction/crawler/engine"
)

func TestSaver(t *testing.T) {
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
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, "dating_test", expected)
	if err != nil {
		t.Errorf("%v", err)
		panic(err)
	}

	// TODO: Try to start up elasticsearch use a client
	resp, err := client.Get().
		Index("dating_test").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	profile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = profile

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
