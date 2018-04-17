package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	//panic("implement me")
	return fmt.Sprintf("{Contents: %s}\n", r.Contents)
}

func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	//panic("implement me")
	return r.Contents
}

