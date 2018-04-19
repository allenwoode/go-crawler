package main

import (
	"fmt"
	"time"
	"math/rand"
)
// Use select to schedule
func worker(id int, c chan int) {
	for ch := range c {
		time.Sleep(time.Second)
		fmt.Printf("index %d, received %d\n", id, ch)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func gen() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = gen(), gen()
	w := createWorker(0)
	n := 0
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("values len:", len(values))
		case <-tm:
			fmt.Println("bye!")
			return
		}
	}
}
