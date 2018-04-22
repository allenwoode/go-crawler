package persist

import "log"

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
