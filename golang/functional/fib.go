package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

func fibonacci() G {
	a, b := 0, 1
	// 闭包
	return func() int {
		a, b = b, a + b
		return a
	}
}

type G func() int

// 函数接口实现
func (g G) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d ", next)

	// TODO: incorrect if p is not enough
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	/*
	fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13
	fmt.Println(f()) // 21
	fmt.Println(f()) // ...
	fmt.Println(f()) // ...
	fmt.Println(f()) // ...
	fmt.Println(f()) // ...
	*/
	printFileContents(f)
}
