package main

import (
	"gopkg.in/olivere/elastic.v5"
	"feilin.com/gocourse/goaction/crawler_distribute/persist"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"log"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
)

func server(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.SaverService{
		Client: client,
		Index:  index,
	})
}

func main() {
	// saver rpc service configure host and index
	log.Fatal(server(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}
