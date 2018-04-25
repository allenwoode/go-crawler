package main

import (
	"gopkg.in/olivere/elastic.v5"
	"feilin.com/gocourse/goaction/crawler_distribute/persist"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"log"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	"flag"
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

var port = flag.Int("port", 0, "the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// saver rpc service configure host and index
	log.Fatal(server(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}
