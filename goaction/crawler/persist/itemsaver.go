package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"fmt"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("ItemSaver: #%d, %v", count, item)
			count++
		}
	}()
	return out
}

func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		//panic(err)
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		//panic(err)
		return "", err
	}

	fmt.Printf("%+v", resp)

	return resp.Id, nil
}