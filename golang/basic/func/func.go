package main

import (
	"fmt"
	"math"
)

func apply(op func(a, b int) int, a, b int) int {
	return op(a, b)
}

func swap(a, b int) (int, int)  {
	return b, a
}

func main() {
	a, b := 2, 10
	c := apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, a, b)

	fmt.Printf("pow(%d, %d) = %d\n", a, b, c)

	d := apply(func(a, b int) int {
		return a * b
	}, a, b)

	fmt.Printf("mul(%d, %d) = %d\n", a, b, d)

	fmt.Println(swap(a, b))
}
