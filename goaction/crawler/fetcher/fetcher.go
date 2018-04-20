package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("http request error code:", resp.StatusCode)
		return nil, errors.New("http request error code:" + string(resp.StatusCode))
	}
	e := determineEncoding(resp.Body)
	uft8Body := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(uft8Body)
}

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}