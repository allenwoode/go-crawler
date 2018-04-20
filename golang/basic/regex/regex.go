package main

import (
	"regexp"
	"fmt"
)

const text = `
my email is gfw.crack.out@gmail.com@qq.com
abc@golang.org
    47591130x@qq.com
		ddd.fe@google.co.jp
`
func main() {
	re := regexp.MustCompile(`([.a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)

	fmt.Println(match)
}
