package main

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
)

// convert int to binary string, 5 to 101, or 13 to 1101
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
		println(result)
	}
	return result
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//convertToBin(13)
	//fmt.Println(
	//	convertToBin(5),
	//	convertToBin(13),
	//)
	printFile("golang/abc.txt")
}
