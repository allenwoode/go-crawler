package client

import (
	"log"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"feilin.com/gocourse/goaction/crawler/engine"
	"errors"
	"net/rpc"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("ItemSaver: #%d, %v", count, item)
			count++

			err := Save(client, item)
			if err != nil {
				log.Printf("Save error: %v", err)
				continue
			}
		}
	}()
	return out, nil
}

func Save(client *rpc.Client, item engine.Item) error {
	if item.Type == "" {
		return errors.New("elasticsearch need a type")
	}

	result := ""
	err := client.Call(config.SaverServiceRpc, item, &result)
	if err != nil || result != "ok" {
		return errors.New("rpc client error")
	}

	return nil
}