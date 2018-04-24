package main

import (
	"testing"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"feilin.com/gocourse/goaction/crawler/engine"
	"feilin.com/gocourse/goaction/crawler/model"
	"time"
)

func TestSaverService(t *testing.T)  {

	go server(":1234", "dating_test")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}

	item := engine.Item{
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
	result := ""
	err = client.Call("SaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("item saver service failed. error: %v, result: %s", err, result)
	}
}
