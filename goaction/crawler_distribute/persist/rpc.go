package persist

import (
	"feilin.com/gocourse/goaction/crawler/persist"
	"feilin.com/gocourse/goaction/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type SaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *SaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		log.Printf("SaverService item: %v", item)
		*result = "ok"
	} else {
		log.Printf("SaverService error: %v", err)
	}

	return err
}
