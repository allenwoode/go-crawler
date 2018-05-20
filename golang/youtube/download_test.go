package main

import (
	"testing"
	"log"
)

func TestDownload(t *testing.T)  {

	filename, url := "lucy", "https://www.youtube.com/watch?v=VIb0W3ogxak"
	log.Println("url: "+url)
	run(filename, url)
}
