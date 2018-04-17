package main

import (
	"feilin.com/gocourse/golang/queue"

	"fmt"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	//q.Push("abc")
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q)
}
