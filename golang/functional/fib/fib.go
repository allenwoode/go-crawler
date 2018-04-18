package fib

import (
	"io"
	"fmt"
	"strings"
)

type G func() int

func Fibonacci() G {
	a, b := 0, 1
	// 闭包
	return func() int {
		a, b = b, a + b
		return a
	}
}

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
