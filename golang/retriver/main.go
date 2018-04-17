package main

import (
	"fmt"
	"feilin.com/gocourse/golang/retriver/mock"
	"time"
	"feilin.com/gocourse/golang/retriver/real"
)

// http get interface
type Retriever interface {
	Get(url string) string
}

const url  = "http://imooc.com"
// User of Retriever interface
func download(r Retriever) string {
	return r.Get(url)
}

// http post interface
type Poster interface {
	Post(url string, form map[string]string) string
}

// interface composed
type RetrieverPoster interface {
	Retriever
	Poster
}

// User of RetrieverPoster
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	//switch v := r.(type) {
	//case *mock.Retriever:
	//	fmt.Println("Contents:", v.Contents)
	//case *real2.Retriever:
	//	fmt.Println("\nUserAgent:", v.UserAgent)
	//}
}

func main() {
	var r Retriever

	r = &mock.Retriever{"http://fake.imooc.com"}
	//fmt.Println(download(r))
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)

	// Type assertion
	rr := r.(*real2.Retriever)
	fmt.Println(rr.UserAgent)

	fmt.Println("try a session")
	m := &mock.Retriever{"你好，我是假的imooc.com"}
	fmt.Println(session(m))
}