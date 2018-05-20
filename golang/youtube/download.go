package main

import (
	"log"
	"os"
	"path/filepath"

	. "github.com/kkdai/youtube"
	"fmt"
)

func run(filename string, url string)  {

	currentFile, _ := filepath.Abs(fmt.Sprintf("~/%s.mp4", filename))
	log.Println("download to file: ", currentFile)

	// NewYoutube(debug)
	y := NewYoutube(true)
	//"https://www.youtube.com/watch?v=rFejpH_tAHM"
	y.DecodeURL(url)
	y.StartDownload(currentFile)
}

func main() {
	//fmt.Println(os.Args)
	currentFile, _ := filepath.Abs(os.Args[1])
	log.Println("download to file: ", currentFile)

	// NewYoutube(debug)
	y := NewYoutube(true)
	url := os.Args[2] //"https://www.youtube.com/watch?v=rFejpH_tAHM"
	y.DecodeURL(url)
	y.StartDownload(currentFile)
}