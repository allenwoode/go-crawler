package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("jobs_citylist_test.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	for _, r := range result.Requests {
		fmt.Printf("%s %s\n", r.Url, r.Parser)
	}
}
