package main

import (
	"gopkg.in/olivere/elastic.v5"
	"crawler_distribute/persist"
	"crawler_distribute/rpcsupport"
	"log"
	"fmt"
	"crawler_distribute/config"
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

var port = flag.Int("port", 5050, "the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// saver rpc service configure host and index
	log.Fatal(server(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}
