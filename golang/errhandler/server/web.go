package main

import (
	"net/http"
	"feilin.com/gocourse/golang/errhandler/server/handler"
	"os"
	"log"
)

type userErr interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(h appHandler) func(w http.ResponseWriter, r *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := h(writer, request)

		if err != nil {
			log.Printf("error occured: %s", err.Error())

			if uErr, ok := err.(userErr); ok {
				http.Error(writer,
					uErr.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(web.Handler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
