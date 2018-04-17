package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
const (
	filename = "abc.txt"
	num = 1
)

const (
	cpp = iota
	_
	python
	golang
	javascript
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	pb
)

func enum()  {
	println(cpp, javascript, python, golang, javascript)
	println(b, kb, mb, gb, pb)
}

func variable6()  {
	println(filename, num)
}

func variable5()  {
	a, b:= 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	println(a, b, c)
}

var aa, bb = 7, 8
var s = "Hello Wrold"
func variable4()  {
	println(aa, bb, s)
}

func variable3()  {
	a, b, s := 5, 6, "def"
	println(a, b, s)
}

func variable2()  {
	var a, b int = 3, 4
	var s string = "abc"
	println(a, b, s)
	//fmt.Println(a, b, s)
}

func variable()  {
	var a int
	var b int
	var s string
	fmt.Printf("%d %d %q\n", a, b, s)
}

func euler()  {
	fmt.Printf("euler: %.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func main() {

	variable()
	variable2()
	variable3()
	variable4()
	variable5()
	variable6()

	enum()
	euler()
}
