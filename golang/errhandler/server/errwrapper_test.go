package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"github.com/pkg/errors"
	"fmt"
)

type userTestErr string

func (e userTestErr) Error() string {
	return e.Message()
}

func (e userTestErr) Message() string {
	return string(e)
}

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}

func errUnKnown(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func errUser(w http.ResponseWriter, r *http.Request) error {
	return userTestErr("user error")
}

func noError(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUser, 400, "user error"},
	{errUnKnown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {

	for _, test := range tests {
		f := errWrapper(test.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://imooc.com",
			nil)
		f(response, request)

		verify(response.Result(), test.code, test.message, t)
	}
}

func TestErrWrapperServer(t *testing.T) {
	for _, test := range tests {
		f := errWrapper(test.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verify(response, test.code, test.message, t)
	}
}

func verify(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T)  {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("exprect (%d %s) got (%d %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}