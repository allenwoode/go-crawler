package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
	"github.com/pkg/errors"
)

var limiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-limiter

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//log.Printf("Request error: %s", http.StatusText(resp.StatusCode))
		return nil, errors.Errorf("response status code: %d", resp.StatusCode)
	}

	b := bufio.NewReader(resp.Body)
	e := determineEncoding(b)
	uft8Body := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(uft8Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding  {
	bytes, err := r.Peek(1024)
	//fmt.Println(string(bytes))
	if err != nil {
		log.Printf("Peek error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}