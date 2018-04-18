package main

import (
	"fmt"
	"feilin.com/gocourse/golang/functional/fib"
	"os"
	"bufio"
)

func deferFunc()  {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)
	panic("error!")
	return

}

func fibWriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	deferFunc()

	//fibWriteFile("fib.txt")
}
