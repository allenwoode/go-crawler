package engine

import (
	"log"
)

type SimpleEngine struct {}

// Simple single task engine
func (e *SimpleEngine) Run(requests ...Request)  {
	var queue []Request
	for _, req := range requests {
		queue = append(queue, req)
	}

	count := 0
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		result, err := Worker(q)
		if err != nil {
			continue
		}
		queue = append(queue, result.Requests...)

		for _, item := range result.Items {
			count++
			log.Printf("Got item #%d: %v", count, item)
		}
	}
}
