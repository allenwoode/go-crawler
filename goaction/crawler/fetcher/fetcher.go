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
	"feilin.com/gocourse/goaction/crawler_distribute/config"
)

var limiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-limiter

	// TODO: 模拟登陆
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//log.Printf("Request error: %s", http.StatusText(resp.StatusCode))
		return nil, errors.Errorf("status code %d", resp.StatusCode)
	}

	body := bufio.NewReader(resp.Body)
	e := determineEncoding(body)
	uft8Body := transform.NewReader(body, e.NewDecoder())

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