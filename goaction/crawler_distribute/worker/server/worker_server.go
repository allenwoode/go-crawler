package main

import (
	"log"
	"feilin.com/gocourse/goaction/crawler_distribute/rpcsupport"
	"feilin.com/gocourse/goaction/crawler_distribute/config"
	"fmt"
	"feilin.com/gocourse/goaction/crawler_distribute/worker"
)

func main()  {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}