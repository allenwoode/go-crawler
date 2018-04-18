package web

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

type userErr string

func (e userErr) Error() string {
	return e.Message()
}

func (e userErr) Message() string {
	return string(e)
}

const prefix  = "/list/"

func Handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userErr("Path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)

	return nil
}
