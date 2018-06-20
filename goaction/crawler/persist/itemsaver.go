package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"feilin.com/gocourse/goaction/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false))
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

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Save error: %v", err)
				continue
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("elasticsearch need a type")
	}

	service := client.Index().
		Index(index).
		Type(item.Type).
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