package main

import (
	"time"
	"fmt"
)

func worker(id int, c chan int)  {
	for ch := range c {
			fmt.Printf("index %d, received %c\n", id, ch)
	}
	//for {
	//	if ch, ok := <-c; ok {
	//		fmt.Printf("index %d, received %c\n", id, ch)
	//	}
	//}
}
// create channel of int and (立刻) return the channel, then (不断地) receive channel by a go routine
// chan<- just receive
// <-chan just send
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo()  {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

// Buffer channel
func bufferedChannel()  {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	c <- 'f'
	c <- 'g'
	c <- 'h'
	c <- 'i'
	c <- 'j'
	c <- 'k'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// channel as first citizen
	chanDemo()
	// buffed channel
	bufferedChannel()
	// sender close channel only by self
	channelClose()
}
