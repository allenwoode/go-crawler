package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//fmt.Printf("index of gorutine: %d\n", i)
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
