package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	const url = "http://album.zhenai.com/u/108656138"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("code %d: url %s", resp.StatusCode, url)
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
