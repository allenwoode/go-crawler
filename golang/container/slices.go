package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[:]
	s2 := arr[:3]
	s3 := arr[5:]
	s4 := s2[2]

	fmt.Println(s1, s2, s3, s4)
}
