package main

import (
	"fmt"
)

type worker struct {
	in chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool)  {
	for ch := range c {
			fmt.Printf("index %d, received %c\n", id, ch)
			go func() {
				done <- true
			}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo()  {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, w := range workers {
		w.in <- 'a' + i
	}

	for _, w := range workers {
		<-w.done
		//<-w.done
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}

	//time.Sleep(time.Millisecond)
	// wait for all channel
	for _, w := range workers {
		//<-w.done
		<-w.done
	}
}

func main() {
	chanDemo()
}
