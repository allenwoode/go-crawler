package real2

import (
	"time"
	"net/http"
	"net/http/httputil"
	"fmt"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) String() string {
	//panic("implement me")
	return fmt.Sprintf("{UserAgent:%s, TimeOut:%s}", r.UserAgent, r.TimeOut)
}

func (r *Retriever) Get(url string) string {
	//panic("implement me")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}

