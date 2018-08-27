package main

import (
	"log"
	"crawler_distribute/rpcsupport"
	"fmt"
	"crawler_distribute/worker"
	"flag"
)

var port = flag.Int("port", 9000, "the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}