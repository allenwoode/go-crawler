package main

import (
	"net/http"
	"feilin.com/gocourse/goaction/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("goaction/crawler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("goaction/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
