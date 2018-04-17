package main

import "fmt"

func lenOfNoRepeatingString(str string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(str) {
		if last, ok := lastOccurred[ch]; ok && last >= start {
			start = last + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
		//fmt.Printf("%d %c\n", maxLength, ch)
	}
	return maxLength
}

func main() {

	fmt.Println(lenOfNoRepeatingString("abcabcbb"))
	fmt.Println(lenOfNoRepeatingString("aaaaa"))
	fmt.Println(lenOfNoRepeatingString(""))
	fmt.Println(lenOfNoRepeatingString("b"))
	fmt.Println(lenOfNoRepeatingString("一二一三三二一"))
}
