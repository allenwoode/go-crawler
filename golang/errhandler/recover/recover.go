package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("error not defined: %v", r))
		}
	}()
	//a, b := 5, 0
	//fmt.Println(a / b)
	//panic(errors.New("this is an error"))
	panic(123) // 123 is not error
}

func main() {
	tryRecover()
}
