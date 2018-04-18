package main

import (
	"fmt"
	"feilin.com/gocourse/golang/container/nonrepeating"
)

func main() {

	fmt.Println(algo.LenOfNoRepeatingString("abcabcbb"))
	fmt.Println(algo.LenOfNoRepeatingString("aaaaa"))
	fmt.Println(algo.LenOfNoRepeatingString(""))
	fmt.Println(algo.LenOfNoRepeatingString("b"))
	fmt.Println(algo.LenOfNoRepeatingString("一二一三三二一"))
	fmt.Println(algo.LenOfNoRepeatingString("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
