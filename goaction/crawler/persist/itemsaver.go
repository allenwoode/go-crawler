package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"feilin.com/gocourse/goaction/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
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

			err := save(client, index, item)
			if err != nil {
				log.Printf("Save error: %v", err)
				continue
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	//client, err := elastic.NewClient(elastic.SetSniff(false))

	if item.Type == "" {
		return errors.New("")
	}

	service := client.Index().
		Index(index).
		Type("zhenai").
		BodyJson(item)
	if item.Id != "" {
		service.Id(item.Id)
	}

	_, err := service.Do(context.Background())
	if err != nil {
		//panic(err)
		return err
	}

	return nil
}