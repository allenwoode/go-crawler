package main

import (
	"net/http"
	"feilin.com/gocourse/goaction/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("goaction/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("goaction/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
